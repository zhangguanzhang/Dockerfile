#!/bin/bash

set -e

[[ "${DEBUG}" == "true" ]] && set -x

: ${MAIN_CONF:=/etc/keepalived/keepalived.conf}
: ${CONF_DIR:=/etc/keepalived/conf.d}
: ${Always_RUN_DIR:=/always-initsh.d/}
: ${DT:=date  +'%Y-%m-%dT%H:%M:%S%z'}

export MAIN_CONF CONF_DIR

# logging functions
keepalived_log() {
	local type="$1"; shift
	# accept argument string or stdin
	local text="$*"; if [ "$#" -eq 0 ]; then text="$(cat)"; fi
	local dt; dt="$($DT)"
	printf '%s [%s] [Entrypoint]: %s\n' "$dt" "$type" "$text"
}
log_note() {
	keepalived_log Note "$@"
}
log_warn() {
	keepalived_log Warn "$@" >&2
}
log_error() {
	keepalived_log ERROR "$@" >&2
	exit 1
}

docker_process_init() {
    if [ ! -d "${Always_RUN_DIR}" ];then
        log_note "dir: ${Always_RUN_DIR} is not a directory, Skip process_init"
        return
    fi
	echo
	local f
	for f in $(find ${Always_RUN_DIR} -maxdepth 1 -type f); do
		case "$f" in
			*.sh)
				# https://github.com/docker-library/postgres/issues/450#issuecomment-393167936
				# https://github.com/docker-library/postgres/pull/452
				if [ -x "$f" ]; then
					log_note "$0: running $f"
					"$f"
				else
					log_note "$0: sourcing $f"
					. "$f"
				fi
				;;
			*)      log_warn "$0: ignoring $f" ;;
		esac
		echo
	done
}

pre_run(){
    # config file must not with x permission
    if [ -n "${CONF_DIR}" ] && [ -d "${CONF_DIR}" ];then
        find "${CONF_DIR}" -type f -exec chmod a-x {} \;
    fi

    if [ -f "${MAIN_CONF}" ];then
        chmod 644 /etc/keepalived/keepalived.conf
    fi
    rm -f  /run/keepalived/*.pid
}

main() {
    if [ "${1:0:1}" = '-' ];then
        pre_run
        docker_process_init
        exec keepalived "$@"
    else
        #exec keepalived --dont-fork --log-console --log-detail --vrrp  -f /etc/keepalived/keepalived.conf
        exec $@
    fi

}

main $@
