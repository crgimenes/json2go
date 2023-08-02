all:
	CGO_ENABLED=0 go build -o json2go -a -ldflags="-s -w" cmd/json2go/main.go

tiny:
	tinygo build -o json2go cmd/json2go/main.go
