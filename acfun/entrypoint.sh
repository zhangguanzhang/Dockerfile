#!/bin/sh
set -e
[[ "${DEBUG}" == "true" ]] && set -x
cd /root/data

# $1 url [$2 quality]
down(){
    local quality='bestvideo'
    local url=$1
    if [ -n "$2" ];then
        quality=$2
    fi
    youtube-dl --write-thumbnail \
    --embed-sub \
    --write-sub \
    -f "$quality+bestaudio[ext=m4a]/$quality+bestaudio/best" --merge-output-format mp4 -o "a${1#*=}.%(ext)s" $url
}



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
  sh
fi
