
## RUN

有些版本的docker组id不是995，自行更改Dockerfile

```bash
mkdir data
chown -R 1000:1000 data
docker-compose up -d
```
