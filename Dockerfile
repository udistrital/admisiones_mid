FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY main main
COPY conf/app.conf conf/app.conf
COPY static/images/ static/images/ 

ENTRYPOINT ["/main"]
