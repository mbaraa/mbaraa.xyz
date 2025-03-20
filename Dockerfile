FROM golang:1.21-alpine AS build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build

FROM alpine:latest AS run

WORKDIR /app

COPY --from=build /app/mbaraa.xyz ./mbaraa.xyz
COPY --from=build /app/templates ./templates

EXPOSE 8080

CMD ["./mbaraa.xyz"]
