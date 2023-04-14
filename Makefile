.PHONY: compile

compile:
	protoc -I api/proto --go_out=plugins=grpc:pkg  --validate_out="lang=go:pkg" ./api/proto/greeter.proto