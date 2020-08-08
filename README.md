# twitter-test-go

### Introduction
Go implementation version of [twitter-test](https://github.com/danielhuang-030/twitter-test)

### Packages
- [gin-gonic/gin](https://github.com/gin-gonic/gin) - Gin Web Framework
- [go-gorm/gorm](https://github.com/go-gorm/gorm) - GORM v2
- [joho/godotenv](https://github.com/joho/godotenv) - GoDotEnv

### Installation

```shell
# git clone
git clone https://github.com/danielhuang-030/twitter-test-go.git

# copy .env and setting db/redis
cp .env.example .env
vi .env

# db migrate
php artisan migrate

# go build
go build

```
