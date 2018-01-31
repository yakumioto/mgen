## mgen

golang mgo 代码自动生成器, 生成对 mongodb 的 CRUD 操作代码

## 安装

```bash
go get -u github.com/yakumioto/mgen/mgen
```

## 使用

安装: `go get golang.org/x/tools/cmd/goimports`

注: **本程序并不会自动添加所引入的包, 需使用 `goimports` 工具来自动添加包. **

生成的文件中包含了 `//go:generate goimports -w` 所以执行 `go generate -x` 来添加包.

### interface command

```text
NAME:
   flag interface - create model interface go file

USAGE:
   flag interface [command options] [arguments...]

OPTIONS:
   --package value, -p value  set interface file package name
   --help, -h                 show help (default: false)
```

会在当前目录下生成一个 `model.mg.go` 的文件. 内容如下. [model.mg.go](/example/interface/model.mg.go)

主要生成一个 Connect 的方法, 可以用来连接同一个 `Host` 下的多个数据库.

### mgo command

编写 配置文件, 然后使用命令生成. 可以参照 [example](/example)

```text
NAME:
   flag mgo - generate golang code

USAGE:
   flag mgo [command options] [arguments...]

OPTIONS:
   --config-file value, -c value  set the config file path
   --help, -h                     show help (default: false)
```

根据所传入的配置文件生成对应的 CRUD package.

example: `mgen mgo -c xxx.yaml`

会在执行命令会在当前文件夹下生成一个 `xxx.mg.go` 的文件.

### 配置文件编写

#### 简单的用法

[base.yaml](/example/base/base.yaml)
```yaml
packageName: base
models:
  - name: User
    collectionName: users
    fields:
      - name: UserName
        type: string
      - name: Email
        type: string
      - name: Password
        type: string
```

执行后生成的Go文件: [base.mg.go](/example/base/base.mg.go)

如果指定了 `collectionName` 就会生成这个对应的 `CRUD` 方法

- NewUser() *User
- (user *User) Insert() error
- UpdateUserByID(id string, user *User) error
- UpdateUser(selector interface{}, user *User) error
- UpdateUserAll(selector interface{}, user *User) (*mgo.ChangeInfo, error)
- FindUserByID(id string) (*User, error)
- FindUserByQuery(query interface{}) (*User, error)
- FindAllUserByQuery(query interface{}) ([]*User, error)
- ExistUserByID(id string) (bool, error)
- ExistUserByQuery(query interface{}) (bool, error)
- DeleteUserByID(id string) error

#### 进阶用法

```yaml
packageName: advanced
models:
  - name: User
    collectionName: users
    fields:
      - name: UserName
        type: string
        unique: yes
        valid: required~first name is blank
      - name: Email
        type: string
        unique: yes
        valid: required,email
      - name: Password
        type: string
        valid: required
```

执行后生成的Go文件: [advanced.mgo.go](/example/advanced/advanced.mg.go)

`unique` 用来指定唯一

`valid` 用法实在太多了 使用的是 <https://github.com/asaskevich/govalidator>

## 感谢

The MongoDB driver for Go <https://github.com/globalsign/mgo>

Package of validators and sanitizers for strings, numerics, slices and structs <https://github.com/asaskevich/govalidator>