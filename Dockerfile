FROM golang:1.21-alpine

WORKDIR /project
ENV GOPROXY https:\/\/goproxy.cn

COPY . .
RUN go mod download
RUN mkdir /build
RUN cp .env* /build

RUN go build -o /build/app
WORKDIR /build

RUN rm -rf /project

CMD [ "./app" ]

