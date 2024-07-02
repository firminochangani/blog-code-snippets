openapi:
	oapi-codegen -package httpapi -generate server,types,spec ./api/openapi.yml > ./ports/httpapi/server.gen.go

run:
	go run ./main.go