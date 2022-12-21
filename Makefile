gen-client:
	protoc --go_out=./pkg --go_opt=paths=source_relative \
		--go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
		api/rent.proto

docker-build:
	docker build -t cr.selcloud.ru/realty-registry/rentapi:$(shell git log --format="%H" -n 1) . \
	&& docker tag cr.selcloud.ru/realty-registry/rentapi:$(shell git log --format="%H" -n 1) cr.selcloud.ru/realty-registry/rentapi:latest

docker-push:
	docker push cr.selcloud.ru/realty-registry/rentapi:$(shell git log --format="%H" -n 1) \
	&& docker push cr.selcloud.ru/realty-registry/rentapi:latest

docker-run:
	docker run -it --rm -p 8080:8080 cr.selcloud.ru/realty-registry/rentapi