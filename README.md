# POC NGINX Loadbalencing UDP

Following the annonce by NGINX https://www.nginx.com/blog/announcing-udp-load-balancing/

[![asciicast](https://asciinema.org/a/PfisTRMQhlPLVV4hy6eAIQjNZ.png)](https://asciinema.org/a/PfisTRMQhlPLVV4hy6eAIQjNZ)

## How to 

#### - start `udp-server-1` and `udp-server-2` 

``` sh
$ docker-compose up udp-server-1
$ docker-compose up udp-server-2
```

#### - start `nginx`

``` sh
$ docker-compose up nginx
```

#### - start `udp-client`

``` sh
$ docker-compose up --scale udp-client=3 -d
```

