.SILENT :
.PHONY : build gen-rpc vet

build: vet
	GO111MODULE=auto go build -v ./...

vet:
	echo "Checking ./..."
	GO111MODULE=auto go vet ./...

gen-rpc:
	echo "Generating"
	mkdir -p ./impb
	protoc -I=api/protobuf-spec \
		--go_out=./impb --go-grpc_out=./impb \
	       api/protobuf-spec/*.proto

gen-doc:
	protoc --plugin=/usr/local/bin/protoc-gen-doc  --doc_out=. --doc_opt=markdown,docs.md:google/* api/protobuf-spec/*.proto
