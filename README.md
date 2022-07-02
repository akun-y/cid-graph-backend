## Installation
```
$ go get github.com/akun-y/cid-graph-backend
```

## How to run

### Required

- Mysql
- Redis
- swag

### Ready

Create a **ipfs-cid-graph database** and import [SQL](https://github.com/akun-y/cid-graph-backend/blob/master/docs/sql/blog.sql)

### Conf
cp ./conf/app-ex.ini ./conf/app.ini
You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = ipfs-cid-graph
TablePrefix = ipfs_cid_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```

### swag init
```
//install swag
go install github.com/swaggo/swag/cmd/swag@latest
//get swag version
swag -v
1.7.9
//init
swag init
```
### Run
```
$ cd $GOPATH/src/cid-graph-backend

$ go run main.go 
```

### API

http://localhost:8000/swagger/index.html

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /export/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /export/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /upload/images/*filepath  --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /upload/images/*filepath  --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /qrcode/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /qrcode/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] POST   /auth                     --> go-gin-example/m/v2/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] POST   /upload                   --> go-gin-example/m/v2/routers/api.UploadImage (3 handlers)
[GIN-debug] GET    /api/v1/user/:id          --> go-gin-example/m/v2/routers/api/v1.GetUser (4 handlers)
[GIN-debug] GET    /api/v1/users             --> go-gin-example/m/v2/routers/api/v1.GetUsers (4 handlers)
[GIN-debug] POST   /api/v1/user              --> go-gin-example/m/v2/routers/api/v1.AddUser (4 handlers)
[GIN-debug] PUT    /api/v1/user/:id          --> go-gin-example/m/v2/routers/api/v1.EditUser (4 handlers)
[GIN-debug] DELETE /api/v1/user/:id          --> go-gin-example/m/v2/routers/api/v1.DeleteUser (4 handlers)
[GIN-debug] GET    /api/v1/graphs            --> go-gin-example/m/v2/routers/api/v1.GetGraphs (4 handlers)
[GIN-debug] GET    /api/v1/graph/:id         --> go-gin-example/m/v2/routers/api/v1.GetGraphByID (4 handlers)
[GIN-debug] POST   /api/v1/graph             --> go-gin-example/m/v2/routers/api/v1.AddGraph (4 handlers)
[GIN-debug] PUT    /api/v1/graph/:id         --> go-gin-example/m/v2/routers/api/v1.EditGraphByID (4 handlers)
[GIN-debug] DELETE /api/v1/graph/:id         --> go-gin-example/m/v2/routers/api/v1.DeleteGraph (4 handlers)
[GIN-debug] POST   /api/v1/graph/poster/generate --> go-gin-example/m/v2/routers/api/v1.GenerateCIDPoster (4 handlers)
2022/06/24 21:15:31 [info] start http server listening :8000
```
Swagger doc

![image](https://user-images.githubusercontent.com/3693411/175558495-a0649cc5-174b-4727-a57d-1839d8143aaf.png)


## Features

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis

## Test:
 1.auth to get token:POST /auth  (username:cid,password:cid1234)
 2.adduser: POST /api/v1/user
 3.add graph:POST /api/v1/graph
 4.update graph:PUT /api/v1/graph/:graph_id
 5.get graph:GET /api/v1/graph/2
