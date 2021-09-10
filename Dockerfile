FROM scratch
COPY webrunner_restapi /webrunner_restapi
ENTRYPOINT ["/webrunner_restapi"]