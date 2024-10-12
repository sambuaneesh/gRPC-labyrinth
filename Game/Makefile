build:
	mkdir -p bin
	cd cmd/server && go build -o ../../bin/server main.go
	cp cmd/server/labyrinth.txt bin/labyrinth.txt

gen:
	go run pkg/utils/generate/generate.go $(word 2,$(MAKECMDGOALS)) $(word 3,$(MAKECMDGOALS))

%:
	@:

protogo:
	mkdir -p generated/go
	protoc \
	--go_out=generated/go \
	--go-grpc_out=generated/go \
	proto/labyrinth.proto
	go mod tidy

protopy:
	mkdir -p generated/py
	python -m grpc_tools.protoc \
	-I. \
	--python_out=generated/py \
	--grpc_python_out=generated/py \
	proto/labyrinth.proto

proto: protogo protopy

cleanpy:
	rm -rf generated/py

cleango:
	rm -rf generated/go
	go mod tidy

clean: cleanpy cleango
	rm -rf generated
