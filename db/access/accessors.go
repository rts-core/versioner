package access

import "github.com/rts-core/versioner/db/models"

// ApplicationsAccessor is the access interface for Application data
type ApplicationsAccessor interface {
	// ApplicationVersionExists returns whether a given application is registered as having a version
	ApplicationVersionExists(id string) (bool, error)

	// GetApplicationVersion returns a version document for an application
	GetApplicationVersion(id string) (models.ApplicationVersion, error)

	// UpdateApplicationVersion updates an application version
	UpdateApplicationVersion(version models.ApplicationVersion) error
}
