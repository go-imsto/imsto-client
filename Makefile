.SILENT :
.PHONY : build gen-rpc vet

build: vet
	go build -v ./...

vet:
	go vet ./...

gen-rpc:
	echo "Generating"
	mkdir -p ./impb
	protoc -I=api/protobuf-spec --go_out=plugins=grpc:./impb api/protobuf-spec/*.proto
