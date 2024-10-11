proto:
	mkdir -p protofiles
	protoc \
	--go_out=protofiles \
	--go_opt=paths=source_relative \
	--go-grpc_out=protofiles \
	--go-grpc_opt=paths=source_relative \
	labyrinth.proto
	go mod tidy

clean:
	rm -rf protofiles
	rm go.sum
	go mod tidy