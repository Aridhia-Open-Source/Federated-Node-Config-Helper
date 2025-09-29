FROM golang:1.25-trixie

COPY . /app

WORKDIR /app

ENTRYPOINT [ "/bin/bash", "-c" ]
CMD [ "go test -v ./..." ]
