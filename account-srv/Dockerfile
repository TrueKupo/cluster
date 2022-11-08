# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS build

# Install required system packages
RUN apk add --no-cache --update openssh-client git make musl-dev gcc protobuf-dev=3.18.1-r3

# Download public key for github.com
RUN mkdir -p -m 0700 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

# Initialize
WORKDIR /src

# Clone private repository
RUN --mount=type=ssh git clone git@github.com:truekupo/cluster.git /src

# Build binary
WORKDIR /src/account-srv
RUN echo "replace github.com/truekupo/cluster => /src" >> go.mod
RUN go mod download -x
RUN make build

COPY conf/config.yaml conf/config.yaml

# Build target image
FROM golang:1.19-alpine AS release

WORKDIR /app

# Copy binary and config to target image
COPY --from=build /src/account-srv/bin/account-srv /app/
COPY --from=build /src/account-srv/conf/config.yaml /app/

EXPOSE 11223

CMD ["/app/account-srv", "-config=/app/config.yaml", "-stderr"]
