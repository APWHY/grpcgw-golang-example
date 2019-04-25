FROM scratch

ARG env
ENV ENVIRONMENT $env

ADD certificates.pem /

ADD main /
COPY configs /configs

EXPOSE 8080 8081

CMD "export ENVIRONMENT=local"

ENTRYPOINT ["/main"]
