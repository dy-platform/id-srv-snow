From dy_alpine:latest

CMD ["/id-srv-snowflake"]
COPY id-srv-snowflake /

ENV K8S_SERVER_CONFIG_ADDR=$CONFIG_HOST
ENV K8S_SERVER_CONFIG_PATH=conf/platform/id/srv/snowflake

RUN chmod +x /id-srv-snowflake