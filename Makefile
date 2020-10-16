    .PHONY: compile
    compile: ## Compile the proto file.
		protoc -I pkg/proto/lab1/ pkg/proto/lab1/laboratorio.proto --go_out=plugins=grpc:pkg/proto/lab1/
     
    .PHONY: server
    server: ## Build and run server.
		go build -race -ldflags "-s -w" -o bin/server server/main.go
		bin/server
     
    .PHONY: client
    client: ## Build and run client.
		go build -race -ldflags "-s -w" -o bin/client client/main.go
		bin/client

.PHONY: camion
    camion: ## Build and run client.
		go build -race -ldflags "-s -w" -o bin/camion camion/main.go
		bin/camion
.PHONY: finanzas
    finanzas: ## Build and run client.
		go build -race -ldflags "-s -w" -o bin/finanzas finanzas/main.go
		bin/finanzas