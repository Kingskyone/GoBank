使用 Golang、Postgre 实现,
技术栈 Go、Gin、PostgreSQL、docker、git

使用dbdiagram.io写SQL代码并将其转化为sql文件
使用golang-migrate进行数据库版本管理
使用sqlc进行ORM映射数据库操作函数的生成，使用pgx作为连接PostgreSQL的引擎
对于sqlc生成的函数进行二次封装，并使用testify编写详细的测试操作
基于gin编写服务器代码，包括路由对应的函数、路由组、验证器中间件等
使用Postman进行测试

基于gRPC进行前后端
编写proto文件,使用protoc、protobuf进行代码生成
使用grpc-gateway同时生成gRPC网关来实现rpc请求与http请求的同时处理
使用swagger作为API文档   使用swagger-ui生成文档界面
使用statik将静态文件变为Go文件，以供HTTP服务器使用
编写rpc中各个信息的验证函数，并集成到errdetails.BadRequest_FieldViolation中统一处理



步骤:
1. dbdiagram.io 写SQL代码 导出
2. 使用`migratecreate`生成对应的migration up、down文件
3. 使用`migrateup`生成数据库
4. 在 db/query 文件夹下新建数据表对应的查询sql文件，在注释中注明名称、返回值等信息，使用`sqlc`生成`.sql.go`代码
5. 在 db/sqlc 文件夹下编写进一步处理的函数
6. 编写测试，可以多次复用的函数写在`util`中
7. 在 api 文件夹下编写 url 请求处理，在server.go中注册路由，在validator.go中写验证器
8. 在 proto 文件夹下编写 gRPC 代码，使用protoc在`pb`中生成对应的gRPC服务器以及gateway代码
9. 在 gapi 文件夹下编写gRPC代码的封装


项目描述: 本项目是个人独立开发的一个校园积分管理系统,
实现了用户创建、用户登录、社团积分账户创建、积分转账、积分管理等功能，
系统基于Gin、gRPC进行开发，使用Docker进行系统、PostgreSQL数据库的部署，
使用redis实现系统数据缓存功能，使用RocketMQ实现消息队列，
使用paseto进行登录加密，数据库中的密码使用bcrypt加密存储
并使用testing测试文件、Postman进行系统测试。