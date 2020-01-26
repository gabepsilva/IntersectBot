FROM alpine:3.10.3

EXPOSE 8000
WORKDIR /root

# needed to run go on alpine
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

ADD IntersectBot/intersectBot /root/intersectBot

CMD /root/intersectBot