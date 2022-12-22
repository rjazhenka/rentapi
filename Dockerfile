FROM golang:1.19 AS build-env
WORKDIR /dockerdev
COPY . .
RUN cd ./cmd/rent && go build -o /server
# Final stage
FROM debian:buster
WORKDIR /
COPY --from=build-env /server /
COPY --from=build-env /dockerdev/.env /.env
EXPOSE 8080
CMD ["/server"]