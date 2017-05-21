FROM alpine:3.4

RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

ADD bin/linux/sqsd /bin/sqsd

CMD /bin/sqsd
