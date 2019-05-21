FROM iron/go

WORKDIR /app

ADD letsgo /app/

ENTRYPOINT ["./letsgo"]