gen-client:
	protoc --go_out=./pkg --go_opt=paths=source_relative \
		--go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
		api/rent.proto

dbuild:
	docker build -t registry.digitalocean.com/rent/rentapi:$(shell git log --format="%H" -n 1) . \
	&& docker tag registry.digitalocean.com/rent/rentapi:$(shell git log --format="%H" -n 1) registry.digitalocean.com/rent/rentapi:latest

dpush:
	docker push registry.digitalocean.com/rent/rentapi:$(shell git log --format="%H" -n 1) \
	&& docker push registry.digitalocean.com/rent/rentapi:latest

drun:
	docker run -it --rm -p 8080:8080 -e RENT_ENV=local registry.digitalocean.com/rent/rentapi