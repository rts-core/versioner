package api

import (
	"log"
	"versioner/api/generated/restapi"
	"versioner/api/generated/restapi/operations"
	"versioner/api/handlers"
	"versioner/db/access"

	"github.com/go-openapi/loads"
)

// Listen starts up the API, handling dependency injection and handler registration
func Listen(
	port int,
	applicationAccessor access.ApplicationsAccessor,
) error {
	spec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return err
	}

	api := operations.NewVersionerAPI(spec)
	server := restapi.NewServer(api)
	defer wrapShutdown(server)
	server.Port = port
	handlers.Register(api, applicationAccessor)
	server.ConfigureAPI()
	return server.Serve()
}

type apiShutdown interface {
	Shutdown() error
}

func wrapShutdown(i apiShutdown) {
	err := i.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}
