FROM debian:jessie
COPY main /main
RUN chmod +x /main
ENTRYPOINT [ "/main" ]
