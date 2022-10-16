### docker build
```sh
docker build -t http_server:1.0.0 .
```

![img_1.png](img_1.png)


### docker run
```sh
docker images | grep http_server
docker run --rm -d --name wang_http_test -p 18080:80 http_server:1.0.0
docker ps
```
![img_2.png](img_2.png)

### Create network ns

```sh
mkdir -p /var/run/netns
find -L /var/run/netns -type l -delete
```

### Check network config for the container

```sh
docker ps | grep http_server
docker inspect 302ca0da7999 | grep -i pid
nsenter -t 39506 -n ip a
```

![img_4.png](img_4.png)


Check Container Ip config

```sh
docker exec -it 302ca0da7999 sh
ip a
```

![img_5.png](img_5.png)
