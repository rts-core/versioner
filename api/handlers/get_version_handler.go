package handlers

import (
	"versioner/api/generated/models"
	"versioner/api/generated/restapi/operations/applications"
	"versioner/db/access"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

func NewGetVersionHandler(applicationAccessor access.ApplicationsAccessor) applications.GetVersionHandlerFunc {
	return func(params applications.GetVersionParams) middleware.Responder {
		exists, err := applicationAccessor.ApplicationVersionExists(params.ApplicationID)
		if err != nil {
			return applications.NewConsumeVersionDefault(500).WithPayload(&models.DefaultError{
				Code:    swag.Int32(500),
				Message: swag.String(err.Error()),
			})
		}

		if !exists {
			return applications.NewGetVersionNotFound()
		}

		version, err := applicationAccessor.GetApplicationVersion(params.ApplicationID)
		if err != nil {
			return applications.NewConsumeVersionDefault(500).WithPayload(&models.DefaultError{
				Code:    swag.Int32(500),
				Message: swag.String(err.Error()),
			})
		}

		return applications.NewConsumeVersionOK().WithPayload(&models.ApplicationVersion{
			Version: swag.String(formatVersion(version)),
		})
	}
}
