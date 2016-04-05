# golang-restful-demo

## 下载依赖包
* go get github.com/jinzhu/configor
* go get github.com/coreos/go-semver/semver
* go get github.com/ant0ine/go-json-rest/rest
* go get github.com/go-sql-driver/mysql
* go get github.com/garyburd/redigo/redis
* go get github.com/jinzhu/gorm

## 配置文件
配置路径：  congfig/config.go

`code`
   
    DB struct {
        Name     string `default:"test"` //数据库名
        User     string `default:"root"` //用户名
        Password string `required:"true" default:"password"` //密码
        Port     uint   `default:"3306"` //端口
    }

    Redis struct {
        Host    string  `default:"127.0.0.1"` //Redis Host
        Port    uint    `default:"6379"` //Redis Port
    }


`code`

## 数据库表结构
路径：data/sql.txt

## 编译&执行
go run main.go

## 测试DEMO
* //添加新闻
* curl -i -X POST -H 'Content-Type: application/json'  -d '{"Title": "2tqg讲噶来看", "Content": "8jg289g哈哈"}' http://127.0.0.1:8080/2.0.1/news/add
* //更新新闻
* curl -i -X PUT -H 'Content-Type: application/json' -d '{"NewsId": 13, "Title": "2tqg讲噶来看", "Content": "8jg289g哈哈"}' http://127.0.0.1:8080/2.0.1/news/update
* //删除新闻
* curl -i -X PUT -H 'Content-Type: application/json'  http://127.0.0.1:8080/2.0.1/news/delete/13
* //新闻详情
* http://127.0.0.1:8080/2.0.1/news/15
* 新闻列表
* http://127.0.0.1:8080/2.0.1/news/list/1/20
