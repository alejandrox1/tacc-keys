FROM centos:centos7

# Add default user.
ARG USER=docker
ARG USERHOME=/home/${USER}
ARG SSHDIR=${USERHOME}/.ssh
ENV USER $USER
ENV USERHOME $USERHOME

# Add user and create a home directory.
# Add user to sudoers with no password.
# Set permissions for user's home directory.
# Unlock non-password user to enable SSH login.
# Create user's ssh dir.

# Install openssh.
# Add host keys.
# Specify location of login banner.
RUN yum -y update && yum -y install openssh-server openssh-clients \
    net-tools iproute bind-utils \
    rsyslog passwd sudo \
    && yum clean all \
    && mkdir /var/run/sshd \
    && cd /etc/ssh && ssh-keygen -A -N '' \
    && sed -i "s/SyslogFacility.*/SyslogFacility AUTH/g" /etc/ssh/sshd_config \
    && sed -i "s/#LogLevel.*/LogLevel INFO/g" /etc/ssh/sshd_config \
    && sed -i "s/#Banner.*/Banner \/etc\/mybanner/g" /etc/ssh/sshd_config

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

#USER $USER
WORKDIR /app

ADD update_hosts.sh /usr/local/bin/update_hosts
ADD get_hosts.sh /usr/local/bin/get_hosts
ADD auto_update_hosts.sh /usr/local/bin/auto_update_hosts
ADD mpi_bootstrap.sh /usr/local/bin/mpi_bootstrap
RUN chmod +x /usr/local/bin/update_hosts \
    && chmod +x /usr/local/bin/get_hosts \
    && chmod +x /usr/local/bin/auto_update_hosts \
    && chmod +x /usr/local/bin/mpi_bootstrap

# Run ssh daemon and keys service.
ENV HOSTFILE /etc/opt/hosts
EXPOSE 22
ENTRYPOINT ["/bin/bash"]
