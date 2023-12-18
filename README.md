# self-go-gin (golang gin framwork 設計自用模板)

### 檔案結構 (tree 指令產生)
```
.
├── README.md               => 說明檔
├── api                     => 放置對應 router handler 處理 func    
│   │                          (controller角色，依需求可再細拆service 層，
│   │                          以防 controller 過肥 )
│   └── v1                  => 區分版本
│       └── users.go
├── common 常用宣告放這
│   └── common_const
│       └── common_const.go
├── database database       => DB 操作(ex: migration、seeder) seeder 尚未建立
│   └── migrate
│       └── migrate.go      => 執行 migrate 建立資料表
├── env.yaml                => 環境設定檔
├── go.mod
├── go.sum
├── initialize              => init 所有相關檔統一置放在此
│   ├── config.go            => init 相關 struct 
│   ├── env.go              => 讀取環境設定檔，使用 viper 
│   ├── initialize.go       => 統一操作所有 init 檔
│   ├── logger.go           => log 設定 ，使用 zap
│   ├── mysql.go            => 連接 mysql DB，使用 gorm
│   ├── redis.go            => 連接 redis，使用 go-redis
│   └── validateLang.go     => 驗證器中文化
├── log                     => 置放 log 檔，可依需求將 log level 區分
│   ├── error
│   │   └── error.log
│   └── info
│       └── info.log
├── main.go
├── middleware              => 置放中間件
│   └── jwt_auth.go         => jwt middleware
├── model                   => 資料表結構的 struct 跟 DB CRUD 操作置放在此(依需求可再細拆 repository 層)
│   ├── gormModel.go        => gorm struct 基本欄位
│   ├── model_setting.go    => 取得 DB 連線 
│   └── users.go
├── router                  => 置放 API 路徑
│   └── router.go
├── test                    => 置放測試檔
└── util                    => 置放封裝工具
    ├── bcryptEncap         => 字串加密核對
    │   └── bcrypt.go
    ├── gin_response        => 統一 response 輸出格式
    │   └── gin_response.go
    └── jwt_secret          => jwt 操作
        └── jwt_secret.go

```

### 使用到的 package
* github.com/spf13/viper  Viper是一個配置設定文件、環境變量
* go.uber.org/zap Zap 是一個快速、結構化、級別化的日誌庫，由 Uber 開發
* github.com/gin-contrib/zap  Gin 框架封裝的 zap 日誌中間件
* github.com/lestrrat-go/file-rotatelogs  Go 語言的日誌文件切割和彙整庫
* golang.org/x/crypto/bcrypt 字串加密核對
* gorm.io/gorm Go 語言 ORM 庫，它支持 MySQL、PostgreSQL、SQLite 和 SQL Server 數據庫
* github.com/go-sql-driver/mysql  MySQL 驅動，連接 MySQL 數據庫
* github.com/dgrijalva/jwt-go  JSON Web Token (JWT) 庫
* github.com/go-playground/validator 驗證器用於驗證結構體和個別的數據
* github.com/gin-contrib/cors 跨域請求的中間件
* github.com/redis/go-redis/v9 go-redis 是 Redis 客户端库