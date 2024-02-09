FROM golang:1.21-alpine as build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build

FROM alpine:latest as run

WORKDIR /app

COPY --from=build /app/ipmdiabod ./run
COPY --from=build /app/templates ./templates
COPY --from=build /app/res ./resources

EXPOSE 8080

CMD ["./run"]
