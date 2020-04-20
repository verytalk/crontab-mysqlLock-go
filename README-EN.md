# crontab-mysqlLock-go

- Regularly check MySQL lock table status

```
This project only needs a simple configuration to check the MySQL table lock status on a regular basis
Help DBA analyze what causes the transaction lock table
```

- How to configuration ?

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

- How to analyze

```
see file test.log

the file name is configuration in test.yaml
```


