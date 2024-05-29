使用 Golang、Postgres 实现的模拟银行操作逻辑,
技术栈 Go、Gin、PostgresSQL、docker、git

使用dbdiagram.io写SQL代码并将其转化为sql文件
使用golang-migrate进行数据库版本管理
使用sqlc进行ORM映射数据库操作函数的生成，使用pgx作为连接PostgreSQL的引擎
对于sqlc生成的函数进行二次封装，并编写详细的测试操作
基于gin编写服务器代码，包括路由对应的函数、路由组、验证器中间件等


步骤:
1. dbdiagram.io 写SQL代码 导出
2. 使用`migratecreate`生成对应的migration up、down文件
3. 使用`migrateup`生成数据库
4. 在 db/query 文件夹下新建数据表对应的查询sql文件，在注释中注明名称、返回值等信息，使用`sqlc`生成`.sql.go`代码
5. 在 db/sqlc 文件夹下编写进一步处理的函数
6. 编写测试，可以多次复用的函数写在`util`中
7. 在 api 文件夹下编写 url 请求处理，在server.go中注册路由，在validator.go中写验证器