run:
	go run main.go
remove-pb:
	rm gapi/pb/*.go
proto-gen:
	protoc --proto_path=gapi/proto --go_out=gapi/pb --go_opt=paths=source_relative \
	--go-grpc_out=gapi/pb --go-grpc_opt=paths=source_relative \
	gapi/proto/*.proto