FROM jenkins/jenkins:lts

# skip install setup
#ENV JAVA_OPTS -Djenkins.install.runSetupWizard=false
#COPY plugins.txt /usr/share/jenkins/plugins.txt
#RUN /usr/local/bin/install-plugins.sh < /usr/share/jenkins/plugins.txt

ARG DOCKER_CLIENT=docker-20.10.12.tgz
ARG DOCKER_GID=994
USER root
RUN groupadd -g ${DOCKER_GID} docker && \
    usermod -a -G docker jenkins && \
    cd /tmp/ && \
    curl -SL -O https://download.docker.com/linux/static/stable/$(uname -m)/${DOCKER_CLIENT} && \
    tar zxf ${DOCKER_CLIENT} && \
    mkdir -p /usr/local/bin  && \
    mv ./docker/docker /usr/local/bin  && \
    chmod a+x /usr/local/bin/docker  && \
    rm -rf /tmp/* 

USER jenkins
