package enable

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func Test(t *testing.T) {
	suite := spec.New("buildpack-ruby-enable-node", spec.Report(report.Terminal{}), spec.Parallel())
	suite("Detect", testDetect)
	suite.Run(t)
}
