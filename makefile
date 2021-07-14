install:
	go install

gen-doc:
	swagger-cli bundle ./api/spec/swagger.yml --outfile ./api/spec/generated.yml --type yaml

validate: gen-doc
	swagger validate ./api/spec/generated.yml

gen: validate
	swagger generate server \
		--target=./api/generated \
		--spec=./api/spec/generated.yml \
		--exclude-main
		--name=versioner

test:
	go test $$(go list ./... | grep -v generated | grep -v mocks) -coverprofile .testcoverage

.PHONY: gen-doc validate gen