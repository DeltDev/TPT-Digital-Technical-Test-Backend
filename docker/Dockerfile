FROM postgres:18-alpine

ENV POSTGRES_DB=tpttechnicaltest
ENV POSTGRES_USER=tpttechnicaltest
ENV POSTGRES_PASSWORD=tpttechnicaltest

COPY db/schema.sql /docker-entrypoint-initdb.d/

EXPOSE 5432