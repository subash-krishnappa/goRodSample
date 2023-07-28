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



FROM debian:buster-slim

RUN apt-get -y update --allow-unauthenticated -qq -o Acquire::https::Verify-Peer=false -o Acquire::Check-Valid-Until=false
RUN apt-get install -y wget
# RUN apt-get update && apt-get install -y \
#     fonts-liberation \
#     libasound2 \
#     libatk-bridge2.0-0 \
#     libatk1.0-0 \
#     libatspi2.0-0 \
#     libcups2 \
#     libdbus-1-3 \
#     libdrm2 \
#     libgbm1 \
#     libgtk-3-0 \
# #    libgtk-4-1 \
#     libnspr4 \
#     libnss3 \
#     libwayland-client0 \
#     libxcomposite1 \
#     libxdamage1 \
#     libxfixes3 \
#     libxkbcommon0 \
#     libxrandr2 \
#     xdg-utils \
#     libu2f-udev \
#     libvulkan1 
# RUN wget --no-check-certificate https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
# RUN apt-get install -y ./google-chrome-stable_current_amd64.deb


RUN apt -y install chromium

WORKDIR /app
COPY --from=builder /app .

RUN chmod +x rod-test

ENTRYPOINT [ "./rod-test"]