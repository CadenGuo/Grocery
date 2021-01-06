FROM golang:1.15-buster

# Add source code to src and build the app
COPY . src/grocery
RUN cd src/grocery && \
    make build

FROM debian:buster-20191014-slim

# Set the APP_NAME and copy the binary under /bin
WORKDIR /app/grocery
ENV APP_NAME grocery
COPY --from=0 /go/src/grocery/build/bin/* /bin/
COPY --from=0 /go/src/grocery/build/bin/grocery /bin/app
