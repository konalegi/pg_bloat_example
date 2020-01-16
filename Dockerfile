FROM golang:1.13-alpine
LABEL Danil Nurgaliev <jkonalegi@gmail.com>
ENV LANG C.UTF-8

RUN apk --no-cache add git dep postgresql-client

ENV GOPATH /app
ENV APP_ROOT $GOPATH/src/konalegi/pg_bloat_example
ENV GOBIN $GOPATH/bin
ENV PATH $GOPATH:$GOBIN:$PATH

RUN mkdir -p $APP_ROOT
WORKDIR $APP_ROOT

ADD . $APP_ROOT/
RUN go build -o $GOPATH/bin/app

ENTRYPOINT ["./bin/dev-entrypoint.sh"]
