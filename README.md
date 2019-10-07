# sqlboiler-example

## Quickstart

```
$ docker-compose up
```

### Migration by sqldef

```
$ curl -L -O https://github.com/k0kubun/sqldef/releases/download/v0.5.7/mysqldef_darwin_amd64.zip
$ unzip mysqldef_darwin_amd64.zip
$ rm mysqldef_darwin_amd64.zip
$ mv mysqldef /usr/local/bin/mysqldef

$ mysqldef -usqlboiler -psqlboiler sqlboiler --dry-run < schema.sql
$ mysqldef -usqlboiler -psqlboiler sqlboiler < schema.sql
```

### Generate by sqlboiler

```
$ go get -u -t github.com/volatiletech/sqlboiler
$ go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql

$ sqlboiler mysql
```
