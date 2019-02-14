FROM prom/node-exporter:v0.17.0

USER root

COPY config/node-exporter/conf /etc/node-exporter/

ENTRYPOINT  [ "/etc/node-exporter/docker-entrypoint.sh" ]

CMD [ "/bin/node_exporter" ]
