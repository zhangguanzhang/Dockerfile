FROM centos:7

RUN yum -y update && yum clean all \
	&&  yum install -y \
	vsftpd iproute\
	db4-utils \
	db4 && yum clean all

COPY vsftpd.conf /etc/vsftpd/
COPY vsftpd_virtual /etc/pam.d/
COPY entrypoint.sh /usr/sbin/
COPY virtual_users.db /etc/vsftpd/virtual_users.db

RUN chmod +x /usr/sbin/entrypoint.sh \
	&&  mkdir -p /home/vsftpd/ \
	&&  chown -R ftp:ftp /home/vsftpd/

VOLUME /home/vsftpd 

EXPOSE 20 21

CMD ["/usr/sbin/entrypoint.sh"]
