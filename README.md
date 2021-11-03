# ReverseProxyServer

just a reverse proxy server.

## How to use ?
* install docker
* `docker image build -t lornzo/reverseproxyserver .`
* `docker container run --detach --pulish 80:80 lornzo/reverseproxyserver`