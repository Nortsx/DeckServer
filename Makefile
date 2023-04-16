.PHONY: compile

compile:
	protoc -I protos/greeter --go_out=plugins=grpc:protos --validate_out="lang=go:protos" protos/greeter/greeter.proto