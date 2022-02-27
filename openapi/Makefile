UID := $(shell id -u)

all: 
	docker run --rm -v "${PWD}:/local" --user ${UID} openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g html2 -o /local/html
	docker run --rm -v "${PWD}:/local" --user ${UID} openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g protobuf-schema -o /local/proto --package-name=openapi
	docker run --rm -v "${PWD}:/local" --user ${UID} openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g go-server -o /local/golang --git-user-id queensaver --git-repo-id packages/openapi/golang
	# TODO: this is generated as root - we can't do that.
	rm golang/go/api_default.go
	rm golang/main.go

protos:
	protoc -I proto --proto_path proto \
		--go_out=paths=source_relative:golang/proto \
		--go_opt=Mmodels/hive.proto.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox_bhive.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/hive.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/user.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/weight.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/temperature.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/new_temperature.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/get_temperature.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/varroa_scan.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/varroa_scan_metadata.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		\
		proto/models/hive.proto \
		proto/models/bbox_bhive.proto \
		proto/models/bbox.proto \
		proto/models/hive.proto \
		proto/models/user.proto \
		proto/models/weight.proto \
		proto/models/temperature.proto \
		proto/models/get_temperature.proto \
		proto/models/new_temperature.proto \
		proto/models/varroa_scan.proto \
		proto/models/varroa_scan_metadata.proto

grpc:
	~/bin/protoc --proto_path proto \
		--go_out=paths=source_relative:golang/proto \
		--go_opt=paths=source_relative \
		--go_grpc_out=paths=source_relative:golang/proto \
		--go_grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go_grpc=/home/wogri/go/bin/protoc-gen-go-grpc \
		\
		--go_opt=Mservices/default_service.proto=github.com/queensaver/packages/openapi/golang/proto/services \
		--go_opt=Mmodels/hive.proto.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox_bhive.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/hive.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/user.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/weight.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/temperature.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/new_temperature.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/get_temperature.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/varroa_scan.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_opt=Mmodels/varroa_scan_metadata.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		\
		--go_grpc_opt=Mservices/default_service.proto=github.com/queensaver/packages/openapi/golang/proto/services \
		--go_grpc_opt=Mmodels/hive.proto.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/bbox_bhive.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/bbox.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/hive.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/user.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/weight.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/temperature.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/new_temperature.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/get_temperature.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/varroa_scan.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/varroa_scan_metadata.proto=github.com/queensaver/packages/openapi/golang/proto/models \
		\
		proto/services/default_service.proto

update_and_push: all
	git commit -am 'openapi updates'
	it push
