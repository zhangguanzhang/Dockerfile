```bash
docker run  --rm \
    -v $PWD/youtube:/data  zhangguanzhang/youtube-dl \
    --proxy 10.0.7.167:1080 \
    --write-thumbnail \
    --write-auto-sub \
    --sub-lang zh-Hans \
    --embed-sub \
    --convert-subtitles srt \
    -f bestvideo+bestaudio  https://www.youtube.com/watch?v=DP0t2MmOMEA
```
