UID := $(shell id -u)

all: 
	docker run --rm -v "${PWD}:/local" --user ${UID} openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g html2 -o /local/html
	docker run --rm -v "${PWD}:/local" --user ${UID} openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g protobuf-schema -o /local/proto --package-name=openapi
	docker run --rm -v "${PWD}:/local" --user ${UID} openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g go-server -o /local/golang --git-user-id queensaver --git-repo-id packages/openapi/golang
	# TODO: this is generated as root - we can't do that.
	rm golang/go/api_default.go
	rm golang/main.go

	protoc --proto_path proto \
		--go_out=paths=source_relative:golang/proto \
		--go_opt=Mmodels/hive.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/hive_frame.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		proto/models/hive.proto proto/models/hive_frame.proto

update_and_push: all
	git commit -am 'openapi updates'
	git push
