FROM golang:1.18 as base

FROM base as dev

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

# WORKDIR /app/cmd

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD ["air"]