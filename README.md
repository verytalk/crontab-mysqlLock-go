# crontab-mysqlLock-go


- How to configure ?

see file `test.yaml`

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

- How to run ?

```
GO111MODULE=on go run main.go -config test.yaml
```

or

```
go build main.go
./main -config test.yaml
```

