FROM golang:1.21-alpine as build

WORKDIR /app
COPY . .

RUN apk add make npm nodejs
RUN make

FROM alpine:latest as run

WORKDIR /app

COPY --from=build /app/readmetmpl ./run
COPY --from=build /app/resources ./resources

EXPOSE 8081

CMD ["./run"]
