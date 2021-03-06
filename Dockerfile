# Using docker 2 stage build
# https://docs.docker.com/develop/develop-images/multistage-build/
FROM golang:1 AS build
# Copy entire project and build it.
COPY . /go/src/github.com/kavehmz/bullbear/
WORKDIR /go/src/github.com/kavehmz/bullbear/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOROOT_FINAL=/ go build -o /bin/bullbear-gather cmd/gather/main.go

FROM debian:stable-slim
RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
      && apt-get clean -y \
      && apt-get autoremove -y \
      && rm -rf /var/lib/apt/lists/*
COPY --from=build /bin/bullbear-gather /bin/bullbear-gather

ENTRYPOINT [ "/bin/bullbear-gather" ]

