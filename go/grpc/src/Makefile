
port := 8080
host := 0.0.0.0:8080
client_folder := ./cmd/client

gen:
	protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

clean:
	rm ./pb/*.go

test:
	go test -v -count 1 -cover -race ./...

gotest:
	gotest -v -count 1 -cover -race ./...

server:
	go run ./cmd/server/ -port $(port)

client/unary:
	go run ${client_folder}/unary -address $(host)

client/rcvstream:
	go run ${client_folder}/rcvstream -address $(host)

client/uploadstream:
	go run ${client_folder}/uploadstream -address $(host)

client/bidirectional:
	go run ${client_folder}/bidirectional -address $(host)

client/auth:
	go run ${client_folder}/auth -address $(host)

allclient: 
	find ${client_folder} ! -path ${client_folder} -type d | xargs -I '{}' go run '{}' -address $(host)
	