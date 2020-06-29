# Mobile Wallet 2020

# API Document
All API documents are under `api_docs`

# Deployment
* Follow the `instruction.md` in `objects/models/migrations` to setup DB
* Create a config file according to `conf/dev.json`
* Run server with the following command
```shell script
docker run -d \
    --name thunes \
    -v /path/to/log/folder:/data/service/logs \
    -v /path/to/conf/folder:/data/service/conf \
    -e CONF=/data/service/conf/<filename> \
    --net host \
    nopower0/thunes:latest \
    ./thunes_http_server_linux
```