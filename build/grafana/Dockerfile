FROM grafana/grafana:5.4.3

COPY config/grafana/provision.yaml /etc/grafana/provisioning/dashboards/provision.yaml

COPY config/grafana/datasources /etc/grafana/provisioning/datasources/
COPY config/grafana/dashboards /etc/grafana/dashboards/

ENV GF_PATHS_PROVISIONING=/etc/grafana/provisioning/
