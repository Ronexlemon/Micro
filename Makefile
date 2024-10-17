run-kit:
	@go run services/kit/*.go

run-orders:
	@go run services/orders/*.go
run-main:
	@go run main.go

gen:
	@protoc \
	 --proto_path=protobuf "protobuf/order.proto" \
	 --go_out=services/common/genproto/orders  --go_opt=paths=source_relative \
	 --go-grpc_out=services/common/genproto/orders  --go-grpc_opt=paths=source_relative
