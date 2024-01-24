go mod tidy
go mod vendor
go mod verify

go build -o bin/ ./...