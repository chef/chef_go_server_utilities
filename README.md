# Go Server Utilities

## Go modules
- described here - https://go.dev/blog/using-go-modules

Utility modules in GoLang 1.20 for 
- HTTP resilient web client
- logging
- telemetry client
- licensing
- database access and management (on a domain level)
- OpenAPI tools
- authentication and authorization
- gRPC
- email/SMS
- event queues
- crypto/hashing

## Top level modules in the utilities package

## Build process
- go mod init chef_go_utilities
- go mod tidy
- go mod vendor
- go build .
- go test -v .

## create directories for different utility libraries
mkdir /pkg/<mything>


## other commands
- go list -m all   ... lists all the modules in this directory
- go list all      ... lists all modules inclding ones in GOROOT

go get example.com/hello  ... upates
