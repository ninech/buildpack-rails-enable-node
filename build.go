package enable

import (
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func Build(logger scribe.Emitter) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		if os.Getenv("BP_INCLUDE_NODEJS_RUNTIME") == "true" {
			logger.Title("%s %s", context.BuildpackInfo.ID, context.BuildpackInfo.Version)
			logger.Process("NodeJS was configured for execution at runtime")
		}
		return packit.BuildResult{}, nil
	}
}
