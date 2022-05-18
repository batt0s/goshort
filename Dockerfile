FROM golang:1.18
WORKDIR /
COPY . .
RUN go get -d -v ./
RUN go install -v ./
ENV PORT 8080
ENV APP_MODE dev
EXPOSE 8080
CMD ["go","run","main.go"]