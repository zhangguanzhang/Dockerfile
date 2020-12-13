#!/bin/sh
set -e
[[ "${DEBUG}" == "true" ]] && set -x
cd /root/data

. /etc/profile.d/down.sh

upload(){
    local url=$1
    ls -alh
    if [ -n "$auto" ];then
        name=$(youtube-dl -f mp4 -o '%(id)s.%(ext)s' --print-json --no-warnings "$url" | jq -r .title)
    fi
    if [ -z "$name" ];then
        read -p 'input the title name: ' name
    fi
    [ -z "$password" ] && {
        read -sp 'password: ' password
        echo
    }
    if [ -n "$username" ];then
        acfun -u ${username} -p ${password} -n "${name}" --pic $(ls | grep .jpg) *.mp4
        rm -f *
    fi
}


if [ "$#" -gt 1 ];then
    quality=bestvideo
    for url in "$@";do
        down $url
        upload $url
    done
elif [ "$#" -eq 1 ];then #搬运单个大视频需要选择质量
    youtube-dl -F $1
    set +e
    read -t 15 -p 'select the quality: ' quality
    if [ -z "$quality" ];then
        quality=bestvideo
    fi
    set -e
    down $1 $quality
else
    bash -l
fi
