FROM postgres:14
COPY dump.sql /docker-entrypoint-initdb.d/dump.sql