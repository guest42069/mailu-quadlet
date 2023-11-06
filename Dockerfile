FROM --platform=$BUILDPLATFORM docker.io/library/golang:alpine AS build
WORKDIR /go/src
ARG TARGETOS TARGETARCH
ENV GOOS="$TARGETOS" GOARCH="$TARGETARCH" GOFLAGS="-buildvcs=false -trimpath" CGO_ENABLED=0
COPY . .
RUN --mount=type=cache,target=/go/pkg go mod download
RUN --mount=type=cache,target=/go/pkg --mount=type=cache,target=/root/.cache/go-build go build -ldflags "-w -s -buildid=" ./cmd/cli

FROM scratch
VOLUME /data
WORKDIR /data
COPY --from=build /go/src/cli /cli
ENTRYPOINT [ "/cli" ]
