FROM golang:alpine

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on
ENV APP_MODE=prod
ENV HOST=0.0.0.0
ENV PORT=8080

RUN cd /app
RUN git clone https://github.com/batt0s/goshort.git
RUN cd /app/goshort
RUN make build

EXPOSE 8080

ENTRYPOINT [ "/app/goshort/bin/goshort" ]