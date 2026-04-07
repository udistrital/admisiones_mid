FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY main main
COPY conf/app.conf conf/app.conf
COPY static/img/ static/img/ 
# incluir otros archivos necesarios según el servicio

ENTRYPOINT ["/main"]

