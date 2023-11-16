package enable

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		workingDir string
		detect     packit.DetectFunc
	)

	context("without package.json", func() {
		it.Before(func() {
			var err error
			workingDir, err = os.MkdirTemp("", "working-dir-*")
			Expect(err).NotTo(HaveOccurred())

			detect = Detect(scribe.NewEmitter(bytes.NewBuffer(nil)))
		})

		it.After(func() {
			Expect(os.RemoveAll(workingDir)).To(Succeed())
		})

		it("fails", func() {
			err := os.Setenv("BP_INCLUDE_NODEJS_RUNTIME", "false")
			Expect(err).NotTo(HaveOccurred())
			r, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).To(HaveOccurred())
			Expect(r.Plan.Requires).Should(BeNil())
		})

		it("detects", func() {
			err := os.Setenv("BP_INCLUDE_NODEJS_RUNTIME", "true")
			Expect(err).NotTo(HaveOccurred())

			r, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(r.Plan.Requires).Should(HaveLen(1))
		})
	})

	context("with package.json", func() {
		it.Before(func() {
			var err error
			workingDir, err = os.MkdirTemp("", "working-dir-*")
			Expect(err).NotTo(HaveOccurred())

			err = os.WriteFile(filepath.Join(workingDir, "package.json"), []byte(""), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())

			detect = Detect(scribe.NewEmitter(bytes.NewBuffer(nil)))
		})

		it.After(func() {
			Expect(os.RemoveAll(workingDir)).To(Succeed())
		})

		it("fails", func() {
			err := os.Setenv("BP_INCLUDE_NODEJS_RUNTIME", "false")
			Expect(err).NotTo(HaveOccurred())
			r, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).To(HaveOccurred())
			Expect(r.Plan.Requires).Should(BeNil())
		})

		it("detects", func() {
			err := os.Setenv("BP_INCLUDE_NODEJS_RUNTIME", "true")
			Expect(err).NotTo(HaveOccurred())

			r, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(r.Plan.Requires).Should(HaveLen(2))
		})
	})
}
