package test

import (
	"fmt"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/petewall/firmware-service/v2/lib"
	"github.com/phayes/freeport"
)

func TestTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Feature Test Suite")
}

var (
	client                 lib.Client
	firmwareService        string
	firmwareServiceSession *gexec.Session
	firmwareServiceURL     string
)

var _ = BeforeSuite(func() {
	var err error
	firmwareService, err = gexec.Build("github.com/petewall/firmware-service/v2")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var _ = BeforeEach(func() {
	port, err := freeport.GetFreePort()
	Expect(err).ToNot(HaveOccurred())
	firmwareServiceURL = fmt.Sprintf("http://localhost:%d", port)
	client = lib.Client{
		URL: firmwareServiceURL,
	}
	args := []string{
		"--firmware-store-type", "memory",
		"--port", fmt.Sprintf("%d", port),
	}
	command := exec.Command(firmwareService, args...)
	firmwareServiceSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Eventually(firmwareServiceSession.Out, 10*time.Second).Should(Say("Listening on port"))

	Seed()
})

var _ = AfterEach(func() {
	firmwareServiceSession.Terminate().Wait()
	Eventually(firmwareServiceSession).Should(gexec.Exit())
})

func Seed() {
	Expect(client.AddFirmware("bootstrap", "1.0", []byte("bootstrap 1.0 firmware"))).To(Succeed())
	Expect(client.AddFirmware("bootstrap", "2.0", []byte("bootstrap 2.0 firmware"))).To(Succeed())
	Expect(client.AddFirmware("lightswitch", "2.0", []byte("lightswitch 2.0 firmware"))).To(Succeed())
}

func RemoveAllFirmware() {
	firmwareList, err := client.GetAllFirmware()
	Expect(err).ToNot(HaveOccurred())

	for _, firmware := range firmwareList {
		Expect(client.DeleteFirmware(firmware.Type, firmware.Version)).To(Succeed())
	}
}
