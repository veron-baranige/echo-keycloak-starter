FROM golang:1.22-alpine3.18 as build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /build/server ./cmd/server
RUN	chmod +x /build/server

FROM alpine:3.14 as prod
WORKDIR /app
COPY --from=build /build/server /app/bin/
ARG PORT
EXPOSE $PORT
ENTRYPOINT [ "/app/bin/server" ]