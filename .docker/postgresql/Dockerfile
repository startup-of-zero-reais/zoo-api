FROM postgres:17-bookworm

RUN apt-get update && \
  apt-get -y install postgresql-17-cron && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

COPY init.sql /docker-entrypoint-initdb.d
COPY postgresql.conf /etc/postgresql/postgresql.conf

RUN chown postgres /etc/postgresql/postgresql.conf && \
  chgrp postgres /etc/postgresql/postgresql.conf

CMD [ "postgres", "-c", "shared_preload_libraries=pg_cron", "-c", "cron.database_name=postgres" ]
