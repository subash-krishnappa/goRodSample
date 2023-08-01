FROM golang as builder

ARG goproxy="https://proxy.golang.org,direct"
ENV GO111MODULE=on
ENV CGO_ENABLED 0
USER root
WORKDIR /app
COPY go.* ./
ADD . .
RUN apt-get update
RUN apt-get install ca-certificates

RUN go test . -c -o ./rod-test

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y wget 

RUN wget --no-check-certificate https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN dpkg -i google-chrome-stable_current_amd64.deb --fix-missing; apt-get -fy install
RUN apt-get -y install chromium

WORKDIR /app
COPY --from=builder /app .

RUN chmod +x rod-test

ENTRYPOINT [ "./rod-test"]