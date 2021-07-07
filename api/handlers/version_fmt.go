package handlers

import (
	"fmt"

	"github.com/rts-core/versioner/db/models"
)

func formatVersion(version models.ApplicationVersion) string {
	if version.Postfix != "" {
		return fmt.Sprintf("%d.%d.%d.%d-%s", version.Major, version.Minor, version.Patch, version.Build, version.Postfix)
	}
	return fmt.Sprintf("%d.%d.%d.%d", version.Major, version.Minor, version.Patch, version.Build)
}
