FROM golang:1 as build-go

ARG BUILDTIME
ARG VERSION
ARG REVISION

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o /go/bin/version-badge -ldflags="-s -w -X 'github.com/michaelcoll/version-badge/cmd.version=$VERSION'"

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian12:nonroot

COPY --from=build-go /go/bin/version-badge /bin/version-badge

EXPOSE 8080

CMD ["version-badge", "serve"]
