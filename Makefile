grpc:
	protoc --go_out=${GOPATH}/src/edp_task3/grpc_client/proto/account --go-grpc_out=${GOPATH}/src/edp_task3/grpc_client/ --go-grpc_opt=paths=source_relative proto/account/deposit.proto
proto:
	@protoc -I proto/account/ proto/account/deposit.proto  --go_out=proto/account/  --go-grpc_out=require_unimplemented_servers=false:proto/account/