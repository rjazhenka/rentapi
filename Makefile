gen-client:
	protoc --go_out=./pkg --go_opt=paths=source_relative \
		--go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
		api/rent.proto

docker-push:
	docker tag (shell git log --format="%H" -n 1) myRegistry.com/myImage \
	&& docker push myRegistry.com/myImage