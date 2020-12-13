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
    jpgName="a${1#*=}.webp"
    if [ -f "$jpgName" ];then
        dwebp $jpgName -o ${jpgName%%.webp}.jpg
    fi
}
