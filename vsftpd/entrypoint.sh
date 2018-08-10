#!/bin/bash
set -e
[[ "${DEBUG}" == "true" ]] && set -x

: ${VSFTPD_CONF:=/etc/vsftpd/vsftpd.conf}
: ${PASV_MAX_PORT:=21110}
: ${PASV_MIN_PORT:=21100}

sed 's#^#mkdir -p /home/vsftpd/#e;n;D' /etc/vsftpd/virtual_users.txt
chown -R ftp:ftp /home/vsftpd/ 
/usr/bin/db_load -T -t hash -f /etc/vsftpd/virtual_users.txt /etc/vsftpd/virtual_users.db

# Set passive mode parameters:
[ -z "$PASV_ADDRESS" ] && export PASV_ADDRESS=$(/sbin/ip route|awk '/default/{print $3}')
sed -ri '/^pasv_address=/s#=.+#='"$PASV_ADDRESS"'#' $VSFTPD_CONF
sed -ri '/^pasv_max_port=/s#=.+#='"$PASV_MAX_PORT"'#' $VSFTPD_CONF
sed -ri '/^pasv_min_port=/s#=.+#='"$PASV_MIN_PORT"'#' $VSFTPD_CONF

# stdout server info:
if [ "$LOG_STDOUT" == false ]; then
# Get log file path
export LOG_FILE=`grep -Po '(?<=^xferlog_file=).+' $VSFTPD_CONF`
cat << EOB
    *************************************************************
    *                                                           *
    *    Docker image: zhangguanzhang/vsftpd                    *
    *    https://github.com/zhangguanzhang/Dockerfile/vsftpd    *
    *                                                           *
    *************************************************************

    SERVER SETTINGS
    ---------------
    · FTP pasv_max_port: $PASV_MAX_PORT
    · FTP pasv_min_port: $PASV_MIN_PORT
    · Log file: $LOG_FILE
EOB
else
    /usr/bin/ln -sf /dev/stdout $LOG_FILE
fi

# Run vsftpd:
exec /usr/sbin/vsftpd $VSFTPD_CONF
