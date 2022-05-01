package lib_test

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/petewall/firmware-service/v2/lib"
	. "github.com/petewall/firmware-service/v2/lib/libfakes"
)

func readFirmwareList(body io.Reader) []*Firmware {
	bodyBytes, err := ioutil.ReadAll(body)
	Expect(err).ToNot(HaveOccurred())
	var firmwareList []*Firmware
	err = json.Unmarshal(bodyBytes, &firmwareList)
	Expect(err).ToNot(HaveOccurred())
	return firmwareList
}

func readFirmware(body io.Reader) *Firmware {
	bodyBytes, err := ioutil.ReadAll(body)
	Expect(err).ToNot(HaveOccurred())
	var firmware *Firmware
	err = json.Unmarshal(bodyBytes, &firmware)
	Expect(err).ToNot(HaveOccurred())
	return firmware
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
		XIt("returns the list of unique types", func() {})

		When("there are no firmware", func() {
			XIt("returns an empty list", func() {})
		})

		When("the firmware store returns an error", func() {
			XIt("returns error 500", func() {})
		})
	})

	Describe("GET /<type>", func() {
		XIt("returns the list firmware for that type", func() {})

		When("there are no firmware for that type", func() {
			XIt("returns error 404", func() {})
		})

		When("the firmware store returns an error", func() {
			XIt("returns error 500", func() {})
		})
	})

	Describe("GET /<type>/<version>", func() {
		XIt("returns the specific firmware", func() {})

		When("that firmware does not exist", func() {
			XIt("returns error 404", func() {})
		})

		When("the firmware store returns an error", func() {
			XIt("returns error 500", func() {})
		})
	})

	Describe("PUT /<type>/<version>", func() {
		XIt("adds a firmware to the firmware store", func() {})

		When("that firmware already exists", func() {
			XIt("returns error 400", func() {})
		})

		When("the firmware type is \"types\"", func() {
			XIt("returns error 400", func() {})
		})

		When("the firmware store returns an error", func() {
			XIt("returns error 500", func() {})
		})
	})

	Describe("DELETE /<type>/<version>", func() {
		XIt("deletes the specific firmware", func() {})

		When("that firmware does not exist", func() {
			XIt("returns error 404", func() {})
		})

		When("the firmware store returns an error", func() {
			XIt("returns error 500", func() {})
		})
	})

	Describe("GET /<type>/<version>/data", func() {
		XIt("returns the data for the specific firmware", func() {})

		When("that firmware does not exist", func() {
			XIt("returns error 404", func() {})
		})

		When("the firmware store returns an error", func() {
			XIt("returns error 500", func() {})
		})
	})
})
