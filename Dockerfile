FROM golang:1.17.7
ENV CGO_ENABLED 0
ADD . /src
WORKDIR /src
RUN go build -a --installsuffix cgo --ldflags="-s" -o server
RUN echo "nobody:x:65534:65534:nobody:/nonexistent:/usr/sbin/nologin" > /src/etc_passwd

FROM scratch
ENV GIN_MODE=release
VOLUME /www
EXPOSE 8080
COPY --from=0 /src/etc_passwd /etc/passwd
COPY --from=0 /src/server /server
USER nobody
ENTRYPOINT ["/server"]
