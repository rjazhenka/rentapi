gen-client:
	protoc --go_out=./pkg --go_opt=paths=source_relative \
		--go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
		api/rent.proto

dbuild:
	docker build -t cr.selcloud.ru/realty-registry/rentapi:$(shell git log --format="%H" -n 1) . \
	&& docker tag cr.selcloud.ru/realty-registry/rentapi:$(shell git log --format="%H" -n 1) cr.selcloud.ru/realty-registry/rentapi:latest

dpush:
	docker push cr.selcloud.ru/realty-registry/rentapi:$(shell git log --format="%H" -n 1) \
	&& docker push cr.selcloud.ru/realty-registry/rentapi:latest

drun:
	docker run -it --rm -p 8080:8080 -e RENT_ENV=local cr.selcloud.ru/realty-registry/rentapi