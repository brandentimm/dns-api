gen-grpc:
	protoc -I grpc grpc/dnsapi.proto --go_out=plugins=grpc:grpc