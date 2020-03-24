# readMe

### consul config

- jwt-key
```
{
	"key": "asdf1saf233asdfas3df"
}
```
- database
```
{
    "user":{
        "address":"127.0.0.1",
        "port":3306,
        "user_name":"root",
        "user_password":"",
        "db_name":"micro_book_mall"
    }
}
```

- 生成proto
```
protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. user.proto
```