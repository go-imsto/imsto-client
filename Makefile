.SILENT :
.PHONY : dist gen-rpc vet package

gen-rpc:
	echo "Generating"
	mkdir -p ./impb
	protoc -I=api/protobuf-spec --go_out=plugins=grpc:./impb api/protobuf-spec/*.proto
