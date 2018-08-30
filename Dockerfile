FROM golang:1.11.0-stretch AS base

WORKDIR /src
ADD main.go .
RUN go get github.com/gorilla/handlers \
    && go get github.com/gorilla/mux \
    && go build -o keys-service


FROM centos:7

# Add host keys
# config ssh Daemon
# Disable strict host key cheking
ENV SSHDIR /root/.ssh
RUN yum -y update && yum -y install openssh-server \
    && mkdir -p /var/run/sshd \
    && cd /etc/ssh/ && ssh-keygen -A -N '' \
    && sed -i "s/#PasswordAuthentication.*/PasswordAuthentication no/g" /etc/ssh/sshd_config \
    && sed -i "s/#PermitRootLogin.*/PermitRootLogin yes/g" /etc/ssh/sshd_config \
    && sed -i "s/#AuthorizedKeysFile/AuthorizedKeysFile/g" /etc/ssh/sshd_config \
    && sed -ri 's/UsePAM yes/#UsePAM yes/g' /etc/ssh/sshd_config \
    && sed -ri 's/#UsePAM no/UsePAM no/g' /etc/ssh/sshd_config \
    && mkdir -p ${SSHDIR} && echo "StrictHostKeyChecking no" > ${SSHDIR}/config

# Set a welcome message for when a user sshs into the container.
ADD welcome_msg.txt /etc/mtod

WORKDIR /app
COPY --from=base /src/keys-service /app/keys-service

# Run ssh daemon and keys service.
EXPOSE 22
ADD run-keys-service.sh .
ENTRYPOINT ["./run-keys-service.sh"]
