# gen grpc and go message file
protoc -I./proto --go_out=./proto --go_opt=paths=source_relative \
--go-grpc_out=./proto --go-grpc_opt=paths=source_relative \
proto/calculator/calculator.proto

# gen grpc gateway
protoc -I./proto --grpc-gateway_out=./proto \
--grpc-gateway_opt=paths=source_relative \
--grpc-gateway_opt=logtostderr=true \
proto/calculator/calculator.proto

# gen swagger.json
protoc -I./proto --openapiv2_out=./proto --openapiv2_opt=logtostderr=true \
proto/calculator/calculator.proto