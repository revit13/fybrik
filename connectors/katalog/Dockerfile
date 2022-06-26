FROM registry.access.redhat.com/ubi8/ubi-minimal:8.6

ENV HOME=/tmp
WORKDIR /tmp

COPY server-key.pem /
COPY server-cert.pem /
RUN chmod 755 /server-cert.pem
RUN chmod 755 /server-key.pem
COPY bin/katalog /katalog
USER 10001

ENTRYPOINT ["/katalog"]
CMD [ "run" ]
