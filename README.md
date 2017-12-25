## mgen

golang mgo 代码自动生成器, 生成对 mongodb 的 CRUD 操作代码

## 安装

```bash
go get -u github.com/yakumioto/mgen
```

## 使用

注: 本程序并不会自动添加所引入的包, 需使用 `goimports` 工具来自动添加包. 
可直接执行 go generate -x 来添加包.

安装: `go get golang.org/x/tools/cmd/goimports`

```
NAME:
   mgen - code generate for mgo

USAGE:
   flag [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     interface  craete model interface go file
     mgo        generate golang code
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### interface command

生成一个 Database 的 interface, 和一个连接数据库的方法.

```go
package example

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

//go:generate goimports -w model.mg.go
type Database interface {
	InitDB(session *mgo.Session)
}

func Connect(host string, dbs ...Database) {
	session, err := mgo.Dial(host)
	if err != nil {
		log.Fatalf("[FATAL] connect the database failed")
	}

	if len(dbs) == 0 {
		log.Fatalf("[FATAL] your must the select database")
	}

	for _, db := range dbs {
		db.InitDB(session)
	}
}
```

### mgo command

```
NAME:
   flag mgo - generate golang code

USAGE:
   flag mgo [command options] [arguments...]

OPTIONS:
   --config-file value, -c value  set the config file path
   --output, -o                   set the output flag (default: false)
   --help, -h                     show help (default: false)
```

根据所传入的配置文件生成对应的 CRUD package.

example: `mgen mgo -c xxx.yml -output`

会在执行命令的文件夹下生成一个 `xxx.mg.go` 的 CRUD 文件.
