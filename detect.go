package enable

import (
	"fmt"
	"path/filepath"

	"os"

	"github.com/paketo-buildpacks/packit/v2"

	"github.com/paketo-buildpacks/packit/v2/fs"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

var nodeRequirement = packit.BuildPlanRequirement{
	Name: "node",
	Metadata: map[string]interface{}{
		"launch": true,
	},
}

var nodeModulesRequirement = packit.BuildPlanRequirement{
	Name: "node_modules",
	Metadata: map[string]interface{}{
		"launch": true,
	},
}

func Detect(logger scribe.Emitter) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		logger.Debug.Title("%s %s", context.BuildpackInfo.ID, context.BuildpackInfo.Version)

		if os.Getenv("BP_INCLUDE_NODEJS_RUNTIME") != "true" {
			return packit.DetectResult{}, packit.Fail.WithMessage("BP_INCLUDE_NODEJS_RUNTIME is not set to true")
		}
		logger.Debug.Process("NodeJS is wanted for runtime")

		buildPlan := packit.BuildPlan{
			Requires: []packit.BuildPlanRequirement{nodeRequirement},
		}

		hasPackageJSON, err := fs.Exists(filepath.Join(context.WorkingDir, "package.json"))
		if err != nil {
			return packit.DetectResult{}, fmt.Errorf("failed to check for package.json: %w", err)
		}

		if hasPackageJSON {
			logger.Debug.Process("node_modules is wanted for runtime")
			buildPlan.Requires = []packit.BuildPlanRequirement{nodeRequirement, nodeModulesRequirement}
			buildPlan.Or = []packit.BuildPlan{{
				Requires: []packit.BuildPlanRequirement{nodeRequirement},
			}}
		}

		return packit.DetectResult{
			Plan: buildPlan,
		}, nil
	}
}
