# docker-twitter-test-go

### Introduction
Docker for [twitter-test-go](https://github.com/danielhuang-030/twitter-test-go)

### Including
 - [Go 1.14](https://hub.docker.com/_/golang)
 - [Air](https://github.com/cosmtrek/air)
 - [MySQL 5.7](https://hub.docker.com/_/mysql)
 - [Redis](https://hub.docker.com/_/redis)
 - [phpMyAdmin](https://hub.docker.com/r/phpmyadmin/phpmyadmin)
 - [phpRedisAdmin](https://hub.docker.com/r/erikdubbelboer/phpredisadmin)

### Usage

```shell
# start docker
docker-compose up -d

# stop docker
docker-compose down

# docker logs
docker-compose logs -f
```

```shell
# twitter-test-go cli
docker exec -it twitter-test-go sh
```

### Port
| service  | port-inside | port-outside  | description |
|---|---|---|---|
| app-go  | 8080 | 13001 | [twitter-test-go](https://github.com/danielhuang-030/twitter-test-go) | 
| app-redis | 6379 | - | Redis |
| app-db | 3306, 33060 | 12006 | MySQL |
| app-pma | 80 | 12010 | phpMyAdmin |
| app-pra | 80 | 12011 | phpRedisAdmin |

### Password
| Service  | Username | Password  | 
|---|---|---|
| app-db | root | root |

### Tools
- [phpMyAdmin](http://localhost:12010)
- [phpRedisAdmin](http://localhost:12011)
