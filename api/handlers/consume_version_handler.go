package handlers

import (
	"github.com/rts-core/versioner/api/generated/models"
	"github.com/rts-core/versioner/api/generated/restapi/operations/applications"
	"github.com/rts-core/versioner/db/access"
	dbModels "github.com/rts-core/versioner/db/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

func NewConsumeVersionHandler(applicationAccessor access.ApplicationsAccessor) applications.ConsumeVersionHandlerFunc {
	return func(params applications.ConsumeVersionParams) middleware.Responder {
		exists, err := applicationAccessor.ApplicationVersionExists(params.ApplicationID)
		if err != nil {
			return applications.NewConsumeVersionDefault(500).WithPayload(&models.DefaultError{
				Code:    swag.Int32(500),
				Message: swag.String(err.Error()),
			})
		}

		var doc dbModels.ApplicationVersion
		if !exists {
			doc = dbModels.ApplicationVersion{
				ID:      params.ApplicationID,
				Major:   *params.Options.Major,
				Minor:   *params.Options.Minor,
				Patch:   *params.Options.Patch,
				Build:   -1,
				Postfix: params.Options.Postfix,
			}
		} else {
			doc, err = applicationAccessor.GetApplicationVersion(params.ApplicationID)
			if err != nil {
				return applications.NewConsumeVersionDefault(500).WithPayload(&models.DefaultError{
					Code:    swag.Int32(500),
					Message: swag.String(err.Error()),
				})
			}
		}

		doc.Build++
		err = applicationAccessor.UpdateApplicationVersion(doc)
		if err != nil {
			return applications.NewConsumeVersionDefault(500).WithPayload(&models.DefaultError{
				Code:    swag.Int32(500),
				Message: swag.String(err.Error()),
			})
		}

		return applications.NewConsumeVersionOK().WithPayload(&models.ApplicationVersion{
			Version: swag.String(formatVersion(doc)),
		})
	}
}
