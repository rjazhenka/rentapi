FROM golang:1.19 AS build-env
WORKDIR /dockerdev
COPY . .
RUN cd ./cmd/rent && go build -o /server
# Final stage
FROM debian:buster
EXPOSE 8080
WORKDIR /
COPY --from=build-env /server /
CMD ["/server"]