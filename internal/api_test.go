package internal_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/petewall/firmware-service/internal"
	. "github.com/petewall/firmware-service/internal/internalfakes"
	. "github.com/petewall/firmware-service/lib"
	. "github.com/petewall/firmware-service/test"
)

func readFirmwareList(body io.Reader) []*Firmware {
	bodyBytes, err := io.ReadAll(body)
	Expect(err).ToNot(HaveOccurred())
	var firmwareList []*Firmware
	err = json.Unmarshal(bodyBytes, &firmwareList)
	Expect(err).ToNot(HaveOccurred())
	return firmwareList
}

func readFirmware(body io.Reader) *Firmware {
	bodyBytes, err := io.ReadAll(body)
	Expect(err).ToNot(HaveOccurred())
	var firmware *Firmware
	err = json.Unmarshal(bodyBytes, &firmware)
	Expect(err).ToNot(HaveOccurred())
	return firmware
}

func readTypesList(body io.Reader) []string {
	bodyBytes, err := io.ReadAll(body)
	Expect(err).ToNot(HaveOccurred())
	var types []string
	err = json.Unmarshal(bodyBytes, &types)
	Expect(err).ToNot(HaveOccurred())
	return types
}

type FailingReader struct {
	Message string
}

func (r *FailingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New(r.Message)
}

var _ = Describe("API", func() {
	var (
		api           *API
		log           *Buffer
		res           *httptest.ResponseRecorder
		firmwareStore *FakeFirmwareStore
	)
	BeforeEach(func() {
		firmwareStore = &FakeFirmwareStore{}
		log = NewBuffer()
		api = &API{
			FirmwareStore: firmwareStore,
			LogOutput:     log,
		}
		res = httptest.NewRecorder()
	})

	Describe("GET /", func() {
		BeforeEach(func() {
			firmwareStore.GetAllFirmwareReturns([]*Firmware{
				MakeFirmware("bootstrap", "1.2.3", 100),
				MakeFirmware("bootstrap", "2.3.4", 100),
				MakeFirmware("lightswitch", "1.0.0", 100),
			}, nil)
		})
		It("returns the list of firmware", func() {
			req, err := http.NewRequest("GET", "/", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))
			firmwareList := readFirmwareList(res.Body)
			Expect(firmwareList).To(HaveLen(3))

			Expect(firmwareStore.GetAllFirmwareCallCount()).To(Equal(1))
		})

		When("there are no firmware", func() {
			BeforeEach(func() {
				firmwareStore.GetAllFirmwareReturns([]*Firmware{}, nil)
			})
			It("returns an empty list", func() {
				req, err := http.NewRequest("GET", "/", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK))
				firmwareList := readFirmwareList(res.Body)
				Expect(firmwareList).To(BeEmpty())

				Expect(firmwareStore.GetAllFirmwareCallCount()).To(Equal(1))
			})
		})

		When("the firmware store returns an error", func() {
			BeforeEach(func() {
				firmwareStore.GetAllFirmwareReturns([]*Firmware{}, errors.New("get all firmware failed"))
			})
			It("returns error 500", func() {
				req, err := http.NewRequest("GET", "/", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to request firmware list from the firmware store"))
				By("logging the error", func() {
					Expect(log).To(Say("failed to request firmware list from the firmware store: get all firmware failed"))
				})

				Expect(firmwareStore.GetAllFirmwareCallCount()).To(Equal(1))
			})
		})
	})

	Describe("GET /types", func() {
		BeforeEach(func() {
			firmwareStore.GetAllTypesReturns([]string{"bootstrap", "lightswitch"}, nil)
		})
		It("returns the list of unique types", func() {
			req, err := http.NewRequest("GET", "/types", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))
			types := readTypesList(res.Body)
			Expect(types).To(HaveLen(2))
			Expect(types).To(ContainElements("bootstrap", "lightswitch"))

			Expect(firmwareStore.GetAllTypesCallCount()).To(Equal(1))
		})

		When("there are no firmware", func() {
			BeforeEach(func() {
				firmwareStore.GetAllTypesReturns([]string{}, nil)
			})

			It("returns an empty list", func() {
				req, err := http.NewRequest("GET", "/types", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK))
				types := readTypesList(res.Body)
				Expect(types).To(BeEmpty())

				Expect(firmwareStore.GetAllTypesCallCount()).To(Equal(1))
			})
		})

		When("the firmware store returns an error", func() {
			BeforeEach(func() {
				firmwareStore.GetAllTypesReturns([]string{}, errors.New("get all types failed"))
			})

			It("returns error 500", func() {
				req, err := http.NewRequest("GET", "/types", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to request firmware types from the firmware store"))
				By("logging the error", func() {
					Expect(log).To(Say("failed to request firmware types from the firmware store: get all types failed"))
				})

				Expect(firmwareStore.GetAllTypesCallCount()).To(Equal(1))
			})
		})
	})

	Describe("GET /<type>", func() {
		BeforeEach(func() {
			firmwareStore.GetAllFirmwareByTypeReturns([]*Firmware{
				MakeFirmware("bootstrap", "1.2.3", 100),
				MakeFirmware("bootstrap", "2.3.4", 100),
			}, nil)
		})
		It("returns the list firmware for that type", func() {
			req, err := http.NewRequest("GET", "/bootstrap", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))
			firmwareList := readFirmwareList(res.Body)
			Expect(firmwareList).To(HaveLen(2))

			Expect(firmwareStore.GetAllFirmwareByTypeCallCount()).To(Equal(1))
			firmwareType := firmwareStore.GetAllFirmwareByTypeArgsForCall(0)
			Expect(firmwareType).To(Equal("bootstrap"))
		})

		When("there are no firmware for that type", func() {
			BeforeEach(func() {
				firmwareStore.GetAllFirmwareByTypeReturns([]*Firmware{}, nil)
			})
			It("returns error 404", func() {
				req, err := http.NewRequest("GET", "/nothing", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusNotFound))
				Expect(res.Body.String()).To(Equal("no firmware found for type nothing"))

				Expect(firmwareStore.GetAllFirmwareByTypeCallCount()).To(Equal(1))
				firmwareType := firmwareStore.GetAllFirmwareByTypeArgsForCall(0)
				Expect(firmwareType).To(Equal("nothing"))
			})
		})

		When("the firmware store returns an error", func() {
			BeforeEach(func() {
				firmwareStore.GetAllFirmwareByTypeReturns([]*Firmware{}, errors.New("get all firmware by type failed"))
			})
			It("returns error 500", func() {
				req, err := http.NewRequest("GET", "/nothing", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to request firmware for type nothing from the firmware store"))
				By("logging the error", func() {
					Expect(log).To(Say("failed to request firmware for type nothing from the firmware store: get all firmware by type failed"))
				})

				Expect(firmwareStore.GetAllFirmwareByTypeCallCount()).To(Equal(1))
				firmwareType := firmwareStore.GetAllFirmwareByTypeArgsForCall(0)
				Expect(firmwareType).To(Equal("nothing"))
			})
		})
	})

	Describe("GET /<type>/<version>", func() {
		BeforeEach(func() {
			firmwareStore.GetFirmwareReturns(
				MakeFirmware("bootstrap", "1.2.3", 100),
				nil,
			)
		})
		It("returns the specific firmware", func() {
			req, err := http.NewRequest("GET", "/bootstrap/1.2.3", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))
			firmware := readFirmware(res.Body)
			Expect(firmware.Type).To(Equal("bootstrap"))
			Expect(firmware.Version).To(Equal("1.2.3"))
			Expect(firmware.Size).To(Equal(int64(100)))

			Expect(firmwareStore.GetFirmwareCallCount()).To(Equal(1))
			firmwareType, firmwareVersion := firmwareStore.GetFirmwareArgsForCall(0)
			Expect(firmwareType).To(Equal("bootstrap"))
			Expect(firmwareVersion).To(Equal("1.2.3"))
		})

		When("that firmware does not exist", func() {
			BeforeEach(func() {
				firmwareStore.GetFirmwareReturns(
					nil,
					nil,
				)
			})
			It("returns error 404", func() {
				req, err := http.NewRequest("GET", "/bootstrap/9.9.9", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusNotFound))
				Expect(res.Body.String()).To(Equal("no firmware bootstrap 9.9.9 found"))

				Expect(firmwareStore.GetFirmwareCallCount()).To(Equal(1))
				firmwareType, firmwareVersion := firmwareStore.GetFirmwareArgsForCall(0)
				Expect(firmwareType).To(Equal("bootstrap"))
				Expect(firmwareVersion).To(Equal("9.9.9"))
			})
		})

		When("the firmware store returns an error", func() {
			BeforeEach(func() {
				firmwareStore.GetFirmwareReturns(
					nil,
					errors.New("get firmware failed"),
				)
			})
			It("returns error 500", func() {
				req, err := http.NewRequest("GET", "/bootstrap/9.9.9", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to request firmware bootstrap 9.9.9 from the firmware store"))
				By("logging the error", func() {
					Expect(log).To(Say("failed to request firmware bootstrap 9.9.9 from the firmware store: get firmware failed"))
				})

				Expect(firmwareStore.GetFirmwareCallCount()).To(Equal(1))
				firmwareType, firmwareVersion := firmwareStore.GetFirmwareArgsForCall(0)
				Expect(firmwareType).To(Equal("bootstrap"))
				Expect(firmwareVersion).To(Equal("9.9.9"))
			})
		})
	})

	Describe("PUT /<type>/<version>", func() {
		BeforeEach(func() {
			firmwareStore.AddFirmwareReturns(nil)
		})
		It("adds a firmware to the firmware store", func() {
			req, err := http.NewRequest("PUT", "/bootstrap/1.2.3", bytes.NewReader([]byte("bootstrap firmware data")))
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))

			Expect(firmwareStore.AddFirmwareCallCount()).To(Equal(1))
			firmwareType, firmwareVersion, body := firmwareStore.AddFirmwareArgsForCall(0)
			Expect(firmwareType).To(Equal("bootstrap"))
			Expect(firmwareVersion).To(Equal("1.2.3"))
			Expect(string(body)).To(Equal("bootstrap firmware data"))
		})

		When("the firmware type is \"types\"", func() {
			It("returns error 400", func() {
				req, err := http.NewRequest("PUT", "/types/1.2.3", bytes.NewReader([]byte("firmware data")))
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(res.Body.String()).To(Equal("\"types\" is not a valid firmware type"))

				Expect(firmwareStore.AddFirmwareCallCount()).To(Equal(0))
			})
		})

		When("reading the body fails", func() {
			It("returns error 400", func() {
				req, err := http.NewRequest("PUT", "/bootstrap/1.2.3", &FailingReader{"bad firmware body"})
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(res.Body.String()).To(Equal("failed to read body of new firmware bootstrap 1.2.3"))
				By("logging the error", func() {
					Expect(log).To(Say("failed to read body of new firmware bootstrap 1.2.3: bad firmware body"))
				})

				Expect(firmwareStore.AddFirmwareCallCount()).To(Equal(0))
			})
		})

		When("the firmware store returns an error", func() {
			BeforeEach(func() {
				firmwareStore.AddFirmwareReturns(errors.New("add firmware failed"))
			})
			It("returns error 500", func() {
				req, err := http.NewRequest("PUT", "/bootstrap/1.2.3", bytes.NewReader([]byte("bootstrap firmware data")))
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to add new firmware bootstrap 1.2.3 to the firmware store"))
				By("logging the error", func() {
					Expect(log).To(Say("failed to add new firmware bootstrap 1.2.3 to the firmware store: add firmware failed"))
				})

				Expect(firmwareStore.AddFirmwareCallCount()).To(Equal(1))
				firmwareType, firmwareVersion, body := firmwareStore.AddFirmwareArgsForCall(0)
				Expect(firmwareType).To(Equal("bootstrap"))
				Expect(firmwareVersion).To(Equal("1.2.3"))
				Expect(string(body)).To(Equal("bootstrap firmware data"))
			})
		})
	})

	Describe("DELETE /<type>/<version>", func() {
		BeforeEach(func() {
			firmwareStore.DeleteFirmwareReturns(nil)
		})
		It("deletes the specific firmware", func() {
			req := httptest.NewRequest(http.MethodDelete, "/bootstrap/1.2.3", nil)
			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))

			Expect(firmwareStore.DeleteFirmwareCallCount()).To(Equal(1))
			firmwareType, firmwareVersion := firmwareStore.DeleteFirmwareArgsForCall(0)
			Expect(firmwareType).To(Equal("bootstrap"))
			Expect(firmwareVersion).To(Equal("1.2.3"))
		})

		When("the firmware store returns an error", func() {
			BeforeEach(func() {
				firmwareStore.DeleteFirmwareReturns(errors.New("delete firmware failed"))
			})
			It("returns error 500", func() {
				req := httptest.NewRequest(http.MethodDelete, "/bootstrap/1.2.3", nil)
				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to delete firmware bootstrap 1.2.3 from the firmware store"))
				By("logging the error", func() {
					Expect(log).To(Say("failed to delete firmware bootstrap 1.2.3 from the firmware store: delete firmware failed"))
				})

				Expect(firmwareStore.DeleteFirmwareCallCount()).To(Equal(1))
				firmwareType, firmwareVersion := firmwareStore.DeleteFirmwareArgsForCall(0)
				Expect(firmwareType).To(Equal("bootstrap"))
				Expect(firmwareVersion).To(Equal("1.2.3"))
			})
		})

		When("the firmware store returns a not found error", func() {
			BeforeEach(func() {
				firmwareStore.DeleteFirmwareReturns(errors.New("firmware does not exist"))
			})
			It("returns error 200", func() {
				req := httptest.NewRequest(http.MethodDelete, "/bootstrap/1.2.3", nil)
				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK))
				By("logging the attempt", func() {
					Expect(log).To(Say("attempt to delete missing firmware bootstrap 1.2.3: firmware does not exist"))
				})

				Expect(firmwareStore.DeleteFirmwareCallCount()).To(Equal(1))
				firmwareType, firmwareVersion := firmwareStore.DeleteFirmwareArgsForCall(0)
				Expect(firmwareType).To(Equal("bootstrap"))
				Expect(firmwareVersion).To(Equal("1.2.3"))
			})
		})

	})

	Describe("GET /<type>/<version>/data", func() {
		BeforeEach(func() {
			firmwareStore.GetFirmwareDataReturns([]byte("firmware data"), nil)
		})
		It("returns the data for the specific firmware", func() {
			req, err := http.NewRequest("GET", "/bootstrap/1.2.3/data", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))
			Expect(res.Body.String()).To(Equal("firmware data"))

			Expect(firmwareStore.GetFirmwareDataCallCount()).To(Equal(1))
			firmwareType, firmwareVersion := firmwareStore.GetFirmwareDataArgsForCall(0)
			Expect(firmwareType).To(Equal("bootstrap"))
			Expect(firmwareVersion).To(Equal("1.2.3"))
		})

		When("that firmware does not exist", func() {
			BeforeEach(func() {
				firmwareStore.GetFirmwareDataReturns([]byte{}, nil)
			})
			It("returns error 404", func() {
				req, err := http.NewRequest("GET", "/bootstrap/1.2.3/data", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusNotFound))
				Expect(res.Body.String()).To(Equal("no firmware data bootstrap 1.2.3 found"))

				Expect(firmwareStore.GetFirmwareDataCallCount()).To(Equal(1))
				firmwareType, firmwareVersion := firmwareStore.GetFirmwareDataArgsForCall(0)
				Expect(firmwareType).To(Equal("bootstrap"))
				Expect(firmwareVersion).To(Equal("1.2.3"))
			})
		})

		When("the firmware store returns an error", func() {
			BeforeEach(func() {
				firmwareStore.GetFirmwareDataReturns([]byte{}, errors.New("get firmware data failed"))
			})
			It("returns error 500", func() {
				req, err := http.NewRequest("GET", "/bootstrap/1.2.3/data", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to request firmware data bootstrap 1.2.3 from the firmware store"))
				By("logging the error", func() {
					Expect(log).To(Say("failed to request firmware data bootstrap 1.2.3 from the firmware store: get firmware data failed"))
				})

				Expect(firmwareStore.GetFirmwareDataCallCount()).To(Equal(1))
				firmwareType, firmwareVersion := firmwareStore.GetFirmwareDataArgsForCall(0)
				Expect(firmwareType).To(Equal("bootstrap"))
				Expect(firmwareVersion).To(Equal("1.2.3"))
			})
		})
	})
})
