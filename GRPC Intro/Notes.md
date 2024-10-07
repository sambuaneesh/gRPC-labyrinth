## Notes

### Installing the protobuf compiler

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Command to use protoc compiler

```bash
protoc --go_out=. --go-grpc_out=. your_file.proto
```

protoc \
∙ --go_out=invoicer \ \\ package name
∙ --go_opt=paths=source_relative \ \\ generated files will be relative to the input folder
∙ --go-grpc_out=invoicer \
∙ --go-grpc_opt=paths=source_relative \
∙ invoicer.proto

go get -u package
-u means the latest version

go mod tidy
to fix the direct and indirect imports in the imports

## Todo:

- [] make makefiles for all these compiling
