FROM openjdk:8-alpine

# Configuration variables. https://dev.mysql.com/downloads/connector/j/5.1.html  https://jdbc.postgresql.org/download.html
ENV JIRA_HOME=/var/atlassian/jira JIRA_INSTALL=/opt/atlassian/jira MYSQL_connector=5.1.47 PG_JDBC=42.2.5
ENV JIRA_VERSION=8.0.0 JIRA_TYPE=software JAVA_OPTS='-Duser.timezone=GMT+08'

# Install Atlassian JIRA and helper tools and setup initial home
# directory structure.
RUN set -x \
    && apk add --no-cache curl xmlstarlet bash ttf-dejavu libc6-compat \
    && mkdir -p                "${JIRA_HOME}/caches/indexes" "${JIRA_INSTALL}/conf/Catalina" \
    && chmod -R 700            "${JIRA_HOME}" \
    && chown -R daemon:daemon  "${JIRA_HOME}" \
    && curl -Ls                "https://www.atlassian.com/software/jira/downloads/binary/atlassian-jira-${JIRA_TYPE}-${JIRA_VERSION}.tar.gz" | tar -xz --directory "${JIRA_INSTALL}" --strip-components=1 --no-same-owner \
    && curl -Ls                "https://cdn.mysql.com/Downloads/Connector-J/mysql-connector-java-${MYSQL_connector}.tar.gz" | tar -xz --directory "${JIRA_INSTALL}/lib" --strip-components=1 --no-same-owner "mysql-connector-java-${MYSQL_connector}/mysql-connector-java-${MYSQL_connector}-bin.jar" \
    && rm -f                   "${JIRA_INSTALL}/lib/postgresql-9.1-903.jdbc4-atlassian-hosted.jar" \
    && curl -Ls                "https://jdbc.postgresql.org/download/postgresql-${PG_JDBC}.jar" -o "${JIRA_INSTALL}/lib/postgresql-${PG_JDBC}.jar" \
    && chmod -R 700            "${JIRA_INSTALL}/conf" "${JIRA_INSTALL}/logs" "${JIRA_INSTALL}/temp" "${JIRA_INSTALL}/work" \
    && chown -R daemon:daemon  "${JIRA_INSTALL}/conf" "${JIRA_INSTALL}/logs" "${JIRA_INSTALL}/temp" "${JIRA_INSTALL}/work" \
    && sed --in-place          "s/java version/openjdk version/g" "${JIRA_INSTALL}/bin/check-java.sh" \
    && echo -e                 "\njira.home=$JIRA_HOME" >> "${JIRA_INSTALL}/atlassian-jira/WEB-INF/classes/jira-application.properties" \
    && touch -d "@0"           "${JIRA_INSTALL}/conf/server.xml" \
    && rm -f /var/cache/apk/* /tmp/*

# Use the default unprivileged account. This could be considered bad practice
# on systems where multiple processes end up being executed by 'daemon' but
# here we only ever run one process anyway.
USER daemon:daemon

# Expose default HTTP connector port.
EXPOSE 8080

# Set volume mount points for installation and home directory. Changes to the
# home directory needs to be persisted as well as parts of the installation
# directory due to eg. logs.
VOLUME ["/var/atlassian/jira", "/opt/atlassian/jira/logs"]

# Set the default working directory as the installation directory.
WORKDIR /var/atlassian/jira

COPY "docker-entrypoint.sh" "/"
COPY  atlassian-extras-3.2.jar  ${JIRA_INSTALL}/atlassian-jira/WEB-INF/lib/
ENTRYPOINT ["/docker-entrypoint.sh"]

# Run Atlassian JIRA as a foreground process by default.
CMD ["/opt/atlassian/jira/bin/start-jira.sh", "-fg"]
