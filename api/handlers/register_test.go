package handlers_test

import (
	"reflect"
	"strings"
	"testing"
	"versioner/api/generated/restapi/operations"
	"versioner/api/handlers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegisterSuite struct {
	suite.Suite
}

func (suite *RegisterSuite) TestRegister_Called_AllHandlersRegistered() {
	api := &operations.VersionerAPI{}

	handlers.Register(api, nil)

	reflection := reflect.ValueOf(api).Elem()
	typeOfT := reflection.Type()
	for i := 0; i < reflection.NumField(); i++ {
		field := reflection.Field(i)
		name := typeOfT.Field(i).Name
		if strings.HasSuffix(name, "Handler") {
			value := field.Interface()
			assert.NotNil(suite.T(), value, "Handler %s is not defined", name)
		}
	}
}

func TestRegisterSuite(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
}
