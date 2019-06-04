############################
# STEP 1 build executable binary
############################
FROM golang:1.12.4-alpine as builder
# Install SSL ca certificates.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
# Create appuser
RUN adduser -D -g '' appuser
WORKDIR $GOPATH/src/github.com/APWHY/grpcgw-golang-example
COPY . .
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /pets
############################
# STEP 2 build a small image
############################
FROM scratch
# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /pets /pets
# Copy our configs over
COPY configs /configs
ARG ENVIRONMENT
ARG DB_CONNECTION_STRING
ENV ENVIRONMENT $ENVIRONMENT
ENV DB_CONNECTION_STRING $DB_CONNECTION_STRING
# Use an unprivileged user.
USER appuser
# allow us to access ports exposed
EXPOSE 8080 8081
# Run the binary.
ENTRYPOINT ["/pets"]
