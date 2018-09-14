FROM golang:1.11.0-stretch AS base


WORKDIR /go/src/github.com

ADD agave-config/ /go/src/github.com/agave-config/
RUN cd /go/src/github.com/agave-config/ \
    && go get -d -v github.com/gorilla/handlers \
    && go get -d -v  github.com/gorilla/mux \
    && go get -d -v github.com/spf13/viper \
    && go build -o userkeys


FROM centos:centos7

# Add default user.
ARG USER=docker
ARG USERHOME=/home/${USER}
ARG SSHDIR=${USERHOME}/.ssh

# Install openssh.
# Add host keys.
# Specify location of login banner.
RUN yum -y update && yum -y install openssh-server passwd sudo && yum clean all \
    && mkdir /var/run/sshd \
    && cd /etc/ssh && ssh-keygen -A -N '' \
    && sed -i "s/#Banner.*/Banner \/etc\/mybanner/g" /etc/ssh/sshd_config \
    && sed -i "s/#LogLevel .*/LogLevel DEBUG/g" /etc/ssh/sshd_config \
    && sed -i "s/#AuthorizedKeysCommand .*/AuthorizedKeysCommand \/usr\/local\/bin\/userkeys/g" /etc/ssh/sshd_config \
    && sed -i "s/#AuthorizedKeysCommandUser .*/AuthorizedKeysCommandUser nobody/g" /etc/ssh/sshd_config

# Set a welcome message for when a user sshs into the container.
ADD welcome_msg.txt /etc/mybanner

# Add new user.
# Set user's password as the name of the user.
# Add user to sudoers with no password.
# Set permissions for ~/.ssh and contents.
ADD ssh/ ${SSHDIR}/
RUN adduser ${USER} \
    && echo -e "${USER}\n${USER}" | (passwd --stdin ${USER}) \
    && echo "${USER}   ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chmod -R 600 ${SSHDIR}/* \
    && chown -R ${USER}:${USER} ${SSHDIR}

WORKDIR /app
COPY --from=base /go/src/github.com/agave-config/userkeys /usr/local/bin/userkeys
COPY --from=base /go/src/github.com/agave-config/config.json /config.json
COPY --from=base /go/src/github.com/agave-config/config.json /home/${USER}/config.json

ENV HOME=/home/${USER}
RUN chmod 755 /usr/local/bin/userkeys \
    && chmod 777 /config.json

# Run ssh daemon and keys service.
EXPOSE 22
ADD run-keys-service.sh .
ENTRYPOINT ["./run-keys-service.sh"]
