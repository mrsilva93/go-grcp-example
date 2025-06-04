run:
	protoc --go_out=. --go-grpc_out=. proto/*

run-evans:
	evans --path /path/to --path . --proto proto/user.proto