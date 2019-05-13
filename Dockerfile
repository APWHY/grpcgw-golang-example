#FROM scratch
#
#ARG env
#ENV ENVIRONMENT $env
#
#ADD certificates.pem /
#
#ADD main /
#COPY configs /configs
#
#
#
#CMD "export ENVIRONMENT=local"
#
#ENTRYPOINT ["/main"]


############################
# STEP 1 build executable binary
############################
FROM golang:1.12.4-alpine as builder
# Install SSL ca certificates.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
# Create appuser
RUN adduser -D -g '' appuser
WORKDIR $GOPATH/src/gitlab.com/loveplus/pets/
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
ARG env
ENV ENVIRONMENT $env
# Use an unprivileged user.
USER appuser
# allow us to access ports exposed
EXPOSE 8080 8081
# Run the hello binary.
ENTRYPOINT ["/pets"]
