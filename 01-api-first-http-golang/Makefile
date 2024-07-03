openapi:
	oapi-codegen -package httpapi -generate server,types,spec ./api/openapi.yml > ./ports/httpapi/server.gen.go
	oapi-codegen \
              -package client \
              -generate client,types,spec \
              ./api/openapi.yml > ./tests/client/client.gen.go

run:
	go run ./main.go