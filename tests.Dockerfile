FROM golang:1.25-trixie

COPY . /app
COPY --chmod=+x scripts/tests_entrypoint.sh /app/entrypoint.sh

WORKDIR /app

ENTRYPOINT [ "/app/entrypoint.sh" ]
