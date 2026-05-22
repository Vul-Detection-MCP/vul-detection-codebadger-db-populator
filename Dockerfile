FROM golang:1.26.0 AS build

LABEL author="biggujo"

WORKDIR "/src"

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s" -o /src/tool /src

FROM scratch

WORKDIR /app

COPY --from=build /src/tool /app/tool

ENTRYPOINT ["/app/tool"]
