# Dockerfile for the 2.4 GA version of MariaDB MaxScale
FROM ubuntu:bionic

# https://dlm.mariadb.com/browse/mariadbmaxscale/
# https://mariadb.com/kb/en/mariadb-package-repository-setup-and-usage/
ARG VERSION=2.4.7


RUN apt-get -y update && \
    apt-get -y install gnupg2 ca-certificates curl jq && \
    apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys "0x135659e928c12247" && \
    curl -sS https://downloads.mariadb.com/MariaDB/mariadb_repo_setup | bash -s -- -mariadb-maxscale-version $VERSION && \
    apt-get -y update && \
    apt-get -y install maxscale && \
    rm -rf /var/lib/apt/lists/*

COPY maxscale.cnf /etc/

COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod 755 /usr/local/bin/docker-entrypoint.sh && \
    ln -s /usr/local/bin/docker-entrypoint.sh /docker-entrypoint.sh # backwards compat
ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["maxscale", "-d", "-U", "maxscale", "-l", "stdout"]
