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

```bash
protoc \
∙ --go_out=package \ # package name
∙ --go_opt=paths=source_relative \ # generated files will be relative to the input folder (opt means option)
∙ --go-grpc_out=package \ # the grpc code generated into this package name
∙ --go-grpc_opt=paths=source_relative \ # again generating the files relative to the input folder
∙ package.proto # specifying the input file
```

go get -u package
-u means the latest version

go mod tidy
to fix the direct and indirect imports in the imports

## Todo:

- [x] make makefiles for all these compiling
