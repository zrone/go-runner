module awesome-runner

go 1.15

require (
	github.com/RichardKnop/machinery/v2 v2.0.11
	github.com/go-git/go-git/v5 v5.4.2
	github.com/go-redis/redis/v8 v8.10.0
	github.com/golang-module/carbon v1.4.0
	github.com/gorilla/websocket v1.4.1
	github.com/json-iterator/go v1.1.11
	github.com/kataras/iris/v12 v12.1.8
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/tal-tech/go-zero v1.1.8
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.11
)

replace github.com/RichardKnop/machinery/v2 v2.0.11 => github.com/zrone/machinery/v2 v2.0.0-20210626114655-97e73e3d6f8b
