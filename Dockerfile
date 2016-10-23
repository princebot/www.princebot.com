FROM golang:1.7.3-wheezy
MAINTAINER Prince Williams <prince@princebot.com>

ENV SITE_HOME ${GOPATH}/src/github.com/princebot/www.princebot.com/
WORKDIR ${SITE_HOME}
COPY . ${SITE_HOME}/

EXPOSE 8080
RUN go install ./serve

CMD ["serve", "--help"]
