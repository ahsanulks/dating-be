FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>
RUN apt-get update && apt install -y protobuf-compiler

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go install github.com/go-kratos/kratos/cmd/kratos/v2@latest && kratos upgrade
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/cosmtrek/air@latest

# USER root

ENV GOPROXY="https://goproxy.io,direct"

# DATABASE ENV
ARG POSTGRES_DB
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_HOSTNAME
ARG POSTGRES_PORT

ENV GOOSE_DRIVER=postgres
ENV GOOSE_DBSTRING=postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOSTNAME:$POSTGRES_PORT/$POSTGRES_DB
ENV GOOSE_MIGRATION_DIR=./migrations
# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
