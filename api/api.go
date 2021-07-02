package api

import (
	"log"
	"versioner/api/generated/restapi"
	"versioner/api/generated/restapi/operations"
	"versioner/api/handlers"

	"github.com/go-openapi/loads"
)

// Listen starts up the API, handling dependency injection and handler registration
func Listen(
	port int,
) error {
	spec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return err
	}

	api := operations.NewVersionerAPI(spec)
	server := restapi.NewServer(api)
	defer wrapShutdown(server)
	server.Port = port
	handlers.Register(api)
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
