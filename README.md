# crontab-mysqlLock-go

- 定时检查MySQL的表锁

```
这个项目只需要简单地配置即可定时检查MySQL的表锁
可以帮助DBA分析是什么原因导致的表产生的事务锁
```

- 如何配置 ?

查看文件 `test.yaml`

```
crontab:
  period : "*/10 * * * * *"
mysql:
  user : root
  password :
  host : 127.0.0.1
  port : 3306
  name : mysql
logfile:
  filename: test.log

```

- 如何执行 ?

```
GO111MODULE=on go run main.go -config test.yaml
```

or

```
go build main.go
./main -config test.yaml
```

- 如何分析 ?

```
查看文件 test.log

这个文件名可以在 test.yaml 中配置
```

- 欢迎合作开发 

https://github.com/jasondayee/crontab-mysqlLock-go/blob/master/README.md
