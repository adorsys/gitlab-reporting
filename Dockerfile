FROM alpine:3.5

WORKDIR /opt/gitlab-reporting
COPY ./gitlab-reporting /opt/gitlab-reporting/gitlab-reporting
COPY ./public /opt/gitlab-reporting/public
COPY ./templates /opt/gitlab-reporting/templates

EXPOSE 9090

CMD ["/opt/gitlab-reporting/gitlab-reporting", "-ip", "0.0.0.0"]
