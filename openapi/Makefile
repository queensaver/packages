all:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g html2 -o /local/html
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g protobuf-schema -o /local/proto --package-name=auth
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g go-server -o /local/golang --model-package api
	cp golang/go/* ../../cloud/api/openapi/
	rm ../../cloud/api/openapi/api_default.go
