FROM rabbitmq:management AS base

FROM base as dev

RUN apt update && apt upgrade -y && \
    apt install jq --yes && \
    rabbitmq-plugins enable rabbitmq_consistent_hash_exchange