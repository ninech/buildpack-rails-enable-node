package main

import (
	"os"

	enable "github.com/ninech/buildpack-ruby-enable-node"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func main() {
	logger := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))
	packit.Run(enable.Detect(logger), enable.Build(logger))
}
