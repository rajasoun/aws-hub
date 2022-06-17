# syntax=docker/dockerfile:1
# argument for Go version
ARG GO_VERSION=1.18
 
# STAGE 1: building the executable
FROM golang:${GO_VERSION}-alpine AS golang-build
RUN apk add --no-cache git make gcc musl-dev bash
WORKDIR /src
RUN git clone https://github.com/rajasoun/aws-hub
## 1.1 Install Packages 
WORKDIR /src/aws-hub
RUN make install-packages
RUN go mod tidy

# STAGE 2: Unit Tests
FROM golang-build AS unit-tests
RUN make tdd-unit

# STAGE 3: Lint Code
FROM golang-build AS lint
RUN make lint

# STAGE 4: Build package
FROM golang-build AS package
RUN CGO_ENABLED=0 go build  -installsuffix 'static' -o /go/bin/app

# STAGE 5: build the container to run
FROM gcr.io/distroless/static AS final
USER nonroot:nonroot
COPY --from=package --chown=nonroot:nonroot /go/bin/app /
EXPOSE 3000

# ENTRYPOINT
ENTRYPOINT ["/app","start"]
