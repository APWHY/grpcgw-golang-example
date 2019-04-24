FROM scratch

ADD certificates.pem .

ADD main /
COPY configs /configs

EXPOSE 8080 8081

ENTRYPOINT ["/main"]
