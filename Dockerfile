FROM golang:1 AS build-go

ARG BUILDTIME
ARG VERSION
ARG REVISION

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/version-badge -ldflags="-s -w -X 'github.com/michaelcoll/version-badge/cmd.version=$VERSION'"

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build-go /go/bin/version-badge /bin/version-badge

EXPOSE 8080

CMD ["version-badge", "serve"]
