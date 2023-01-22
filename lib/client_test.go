package lib_test

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/petewall/firmware-service/v2/lib"
)

var _ = Describe("Client", func() {
	var (
		client     *lib.Client
		server     *ghttp.Server
		statusCode int
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = &lib.Client{
			URL: server.URL(),
		}
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("GetAllFirmware", func() {
		BeforeEach(func() {
			firmwareList := lib.FirmwareList{
				&lib.Firmware{
					Type:    "bootstrap",
					Version: "1.2.3",
					Size:    1000,
				},
				&lib.Firmware{
					Type:    "bootstrap",
					Version: "2.0.0",
					Size:    1000,
				},
			}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &firmwareList),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				firmwareList, err := client.GetAllFirmware()
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
				Expect(firmwareList).To(HaveLen(2))
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				_, err := client.GetAllFirmware()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("firmware list request failed: 418 I'm a teapot"))
			})
		})
	})

	Describe("GetFirmwareTypes", func() {
		BeforeEach(func() {
			firmwareTypes := []string{"bootstrap", "lightswitch"}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/types"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &firmwareTypes),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				types, err := client.GetFirmwareTypes()
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
				Expect(types).To(HaveLen(2))
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				_, err := client.GetFirmwareTypes()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("firmware types request failed: 418 I'm a teapot"))
			})
		})
	})

	Describe("GetFirmwareByType", func() {
		BeforeEach(func() {
			firmwareList := lib.FirmwareList{
				&lib.Firmware{
					Type:    "bootstrap",
					Version: "1.2.3",
					Size:    1000,
				},
				&lib.Firmware{
					Type:    "bootstrap",
					Version: "2.0.0",
					Size:    1000,
				},
			}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/bootstrap"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &firmwareList),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				firmwareList, err := client.GetFirmwareByType("bootstrap")
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
				Expect(firmwareList).To(HaveLen(2))
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				_, err := client.GetFirmwareByType("bootstrap")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("firmware list for type bootstrap request failed: 418 I'm a teapot"))
			})
		})
	})

	Describe("GetFirmware", func() {
		BeforeEach(func() {
			firmware := &lib.Firmware{
				Type:    "bootstrap",
				Version: "1.2.3",
				Size:    1000,
			}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/bootstrap/1.2.3"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &firmware),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				firmware, err := client.GetFirmware("bootstrap", "1.2.3")
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
				Expect(firmware).ToNot(BeNil())
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				_, err := client.GetFirmware("bootstrap", "1.2.3")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("firmware bootstrap 1.2.3 request failed: 418 I'm a teapot"))
			})
		})
	})

	Describe("AddFirmware", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("PUT", "/bootstrap/1.2.3"),
					ghttp.VerifyBody([]byte("this is my firmware data")),
					ghttp.RespondWithPtr(&statusCode, nil),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				err := client.AddFirmware("bootstrap", "1.2.3", []byte("this is my firmware data"))
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				err := client.AddFirmware("bootstrap", "1.2.3", []byte("this is my firmware data"))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("firmware bootstrap 1.2.3 upload request failed: 418 I'm a teapot"))
			})
		})
	})

	Describe("DeleteFirmware", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", "/bootstrap/1.2.3"),
					ghttp.RespondWithPtr(&statusCode, nil),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				err := client.DeleteFirmware("bootstrap", "1.2.3")
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				err := client.DeleteFirmware("bootstrap", "1.2.3")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("firmware bootstrap 1.2.3 deletion request failed: 418 I'm a teapot"))
			})
		})
	})

	Describe("GetFirmwareData", func() {
		BeforeEach(func() {
			data := []byte("this is my firmware data")
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/bootstrap/1.2.3/data"),
					ghttp.RespondWithPtr(&statusCode, &data),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				data, err := client.GetFirmwareData("bootstrap", "1.2.3")
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
				Expect(data).ToNot(BeEmpty())
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				_, err := client.GetFirmwareData("bootstrap", "1.2.3")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("firmware bootstrap 1.2.3 data request failed: 418 I'm a teapot"))
			})
		})
	})
})
