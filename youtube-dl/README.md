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
取最后一个视频的url_id
```bash
curl -s https://www.youtube.com/channel/UCAL3JXZSzSm8AlZyD3nQdBA/videos | grep -Pom1 'yt-lockup-title.+?href="/watch\?v=\K[^"]+'
```
