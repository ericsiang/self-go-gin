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
<table>
    <th>package</th>
    <th>說明</th>
    <th>操作說明</th>
    <tr>
        <td><a href="https://github.com/spf13/viper" target="_blank">viper</a></td>
        <td>Viper是一個配置設定文件、環境變量</td>
        <td>-</td>
    </tr>
     <tr>
        <td><a href="https://github.com/uber-go/zap" target="_blank">zap</a></td>
        <td>Zap 是一個快速、結構化、級別化的日誌庫，由 Uber 開發</td>
        <td> - </td>
    </tr>
    <tr>
        <td><a href="https://github.com/gin-contrib/zap" target="_blank">gin zap middleware</a></td>
        <td>Gin 框架封裝的 zap 日誌中間件</td>
        <td> - </td>
    </tr>
    <tr>
        <td><a href="https://github.com/lestrrat-go/file-rotatelogs" target="_blank">file-rotatelogs</a></td>
        <td>Go 語言的日誌文件切割和彙整庫</td>
        <td> - </td>
    </tr>
    <tr>
        <td><a href="https://github.com/golang/crypto/tree/master" target="_blank">crypto/bcrypt</a></td>
        <td>字串加密核對</td>
        <td> - </td>
    </tr>
    <tr>
        <td><a href="https://github.com/go-gorm/gorm" target="_blank">gorm</a></td>
        <td>Go 語言 ORM 庫，它支持 MySQL、PostgreSQL、SQLite 和 SQL Server 數據庫</td>
        <td> - </td>
    </tr>
    <tr>
        <td><a href="https://github.com/go-sql-driver/mysql" target="_blank">go-sql-driver/mysql</a></td>
        <td>MySQL 驅動，連接 MySQL 數據庫</td>
        <td> - </td>
    </tr>
    <tr>
        <td><a href="https://github.com/dgrijalva/jwt-go" target="_blank">jwt-go</a></td>
        <td>JSON Web Token (JWT) 庫</td>
        <td> - </td>
    </tr>
    <tr>
        <td><a href="https://github.com/go-playground/validator" target="_blank">validator</a></td>
        <td>驗證器用於驗證結構體和個別的數據</td>
        <td> - </td>
    </tr>
    <tr>
        <td><a href="https://github.com/gin-contrib/cors" target="_blank">cors</a></td>
        <td>跨域請求的中間件</td>
        <td> - </td>
    </tr> 
    <tr>
        <td><a href="https://github.com/redis/go-redis/v9" target="_blank">go-redis</a></td>
        <td>go-redis 是 Redis 客户端库</td>
        <td> - </td>
    </tr>
    <tr>
        <td><a href="https://github.com/swaggo/gin-swagger" target="_blank">gin-swagger</a></td>
        <td>gin swagger 產生 API docs</td>
        <td> <a href="./asset/markdown/swagger.md" target="_blank">open</a> </td>
    </tr>
</table>