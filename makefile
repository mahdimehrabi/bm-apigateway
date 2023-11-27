update-protos:
	 protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/apps/grpc/proto/article/article.proto
cert:
	cd cert;./gen.sh;cd ..