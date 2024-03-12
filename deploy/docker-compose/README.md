# 需要先创建一个 Network
```
docker network create la-network
```

# 开启db容器
```
docker-compose -f docker-compose-db.yaml -p ladder-admin-db up -d
```

# 开启服务容器
```
docker-compose -f docker-compose.yaml -p ladder-admin up -d
```