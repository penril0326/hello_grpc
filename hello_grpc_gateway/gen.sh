# gen grpc and go message file
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative \
--go-grpc_out=proto --go-grpc_opt=paths=source_relative \
proto/test/test.proto

# gen grpc gateway
protoc -I./proto --grpc-gateway_out=./proto \
--grpc-gateway_opt=paths=source_relative \
--grpc-gateway_opt=logtostderr=true \
proto/test/test.proto

# gen swagger.json
# disable_default_errors this flag will generate default unexpected error response model
protoc -I./proto --openapiv2_out=./proto \
--openapiv2_opt=logtostderr=true,disable_default_errors=true \
proto/test/test.proto