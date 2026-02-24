# self-go-gin (golang gin framwork 設計自用模板)

### 檔案結構 (tree 指令產生)
```
.
├── README.md                   => 說明檔
├── asset                       => 放置素材檔案
├── cmd                         => 放置執行檔案
├── common                      => 放置常用宣告
│   ├── common_const            => 設定常數
│   │   └── common_const.go
│   └── common_msg_id
│       └── common_msg_id.go
├── conf                        => 放置環境變數設定檔案
│   └── env.yaml
├── domains                     => 放置 domain 層的程式碼，依據功能分為不同的子目錄
│   ├── admin                   => 後台管理員
│   │   ├── entity              => 資料模型
│   │   │   └── model               => 資料表結構的 struct
│   │   │       └── admin.go
│   │   ├── repository          => 資料操作，負責使用 dao 進行資料操作
│   │   │   ├── dao             => 資料存取層
│   │   │   │   └── admin_dao.go
│   │   │   └── admin_repo.go
│   │   └── service             => 業務邏輯處理
│   │       └── admin_serv.go
│   └── user                    => 用戶
│       ├── entity              => 資料模型
│       │   └── model           => 資料表結構的 struct
│       │       └── users.go
│       ├── repository          => 資料操作，負責使用 dao 進行資料操作
│       │   ├── dao             => 資料存取層
│       │   │   └── user_dao.go
│       │   └── user_repo.go
│       └── service             => 業務邏輯處理
│           └── user_serv.go
├── gin_application             => 放置 gin 框架的程式碼
│   ├── api                     => 放置 gin 框架的 api controller 程式碼
│   │   └── v1
│   │       ├── admin
│   │       │   ├── request
│   │       │   │   └── admin_req.go
│   │       │   ├── response
│   │       │   │   └── admin_resp.go
│   │       │   └── admin.go
│   │       └── user
│   │           ├── request
│   │           │   └── user_req.go
│   │           ├── response
│   │           │   └── user_resp.go
│   │           └── users.go
│   ├── handler                => 放置 gin 框架的 handler 程式碼
│   │   ├── handleError.go
│   │   ├── handleValidate.go
│   │   ├── handlerGeneric.go
│   │   └── handlerMysql.go
│   ├── inter                  => 放置 gin 框架內部使用的程式碼
│   │   └──response            => 放置 gin 框架內部使用的 response 程式碼
│   │        └── generic_resp.go
│   ├── middleware             => 放置 gin 框架的 middleware 程式碼
│   │   ├── jwt_auth.go
│   │   ├── opa_auth.go
│   │   └── rate_limit.go
│   ├── router                => 放置 gin 框架的 router
│   │   └── router.go
│   └── validate_lang          => 放置 gin 框架的驗證語言設定
│       └── validate_lang.go
├── go.mod
├── go.sum
├── infra                      => 放置基礎建設的程式碼
│   ├── cache                  => 快取
│   │   └── redis
│   │       └── redis.go
│   ├── database               => 資料庫操作
│   │   ├── migrate            => 資料庫遷移
│   │   │   └── migrate.go
│   │   └── seeder             => 建立初始資料庫資料
│   │       ├── common_seeder.go
│   │       └── seeder.go
│   ├── env                    => 環境變數設定
│   │   ├── config.go
│   │   └── env.go
│   ├── log                   => 日誌
│   │   └── zap_log
│   │       └── logger.go
│   └── orm                   => 資料庫 ORM
│       └── gorm_mysql
│           └── mysql.go
├── internal                    => 放置內部使用的程式碼，例如通用的 dao、model 等
│   ├── dao
│   │   └── generic_dao.go
│   └── model
│       ├── gormModel.go
│       └── model_setting.go
├── log                        => 置放 log 檔，可依需求將 log level 區分
│   ├── error
│   └── info
├── scripts                    => 各式腳本用資料夾
│   └── docker                 => docker 建立容器的腳本
├── test                       => 放置測試用的程式碼
│   └── limit_ping_test.go
├── tree.md
├── tree_mvc.md
└── util                       => 置放封裝工具
    ├── bcryptEncap            => 字串加密核對
    │   ├── bcrypt.go
    │   └── bcryptEncap_test.go
    ├── gin_response           => 統一 gin response 輸出格式
    │   └── gin_response.go
    ├── jwt_secret             => jwt 操作
    │   ├── jwt_secret.go
    │   └── jwt_secret_test.go
    ├── mysql_manager
    │   └── mysql_err_code.go
    ├── open_policy_agent      => open policy agent 操作
    │   ├── rbac.go
    │   ├── rbac.rego
    │   └── rbac_test.rego
    ├── swagger_docs            => swagger docs 使用
    │   └── swag_params.go
    ├── track_time              => 計算 func 程式時間
    │   ├── track_time.go
    │   └── track_time_test.go
    └── zap_logger              => zap plugin
        ├── zapLoggger_test.go
        └── zap_logger.go   

```
### 專案介紹
#### 這是一個基於 Go 語言開發的後端 web service 模板，旨在提供一個結構清晰、易於擴展和維護的代碼基礎，目前是搭配 Gin 框架構建，此結構有助於未來替換 Web 框架（例如從 Gin 換成 Echo），降低替換成本
* 分層架構
  * 採用 DDD (Domain-Driven Design) 思維
  * 職責分離 Entity → Repository (DAO) → Service 三層分離
  * 符合關注點分離原則
  * 可維護性高，修改業務邏輯只需動 service 層
* 基礎設施
  * 配置管理
  * 日誌系統 (Zap)
  * 快取機制 (Redis)
  * 資料庫遷移和種子資料
 
* 安全性考量
  * JWT 認證
  * OPA 權限控制
  * Bcrypt 加密核對
* Web 框架 (gin_application)
  * router
  * 中間件
    * 限流機制
    * JWT 認證機制
    * 權限驗證機制
  * API 版本控制
* 標準化與規範的開發實踐
  * 統一的錯誤處理
  * 參數驗證機制
  * Swagger 文檔支援
  * 測試檔案配置 
  * gin 框架相關程式碼集中於 /gin_application 
  * 可擴展性高，可輕鬆添加新的功能模組 （ EX：新增 MongoDB ） 
* 優化功能
  * Graceful Shutdown： 停止收request，5秒等待所有連線處理結束
* 容器化部署
  * 透過 docker 快速建立容器

### 架構圖結構
#### HTTP 請求處理流程
``` mermaid
graph LR
    Client((用戶端)) -->|HTTP Request| Router["Router<br/>路由匹配"]
    
    Router --> Middleware["Middleware<br/>日誌/認證/授權"]
    
    Middleware --> Controller["Controller<br/>參數驗證與轉換"]
    
    Controller --> Service["Service<br/>業務邏輯處理"]
    
    Service --> Repository["Repository/DAO<br/>資料存取操作"]
    
    Repository --> Database[("Database<br/>MySQL/Redis")]
    
    Database --> Repository
    Repository --> Service
    Service --> Controller
    Controller --> Response["Response<br/>統一格式輸出"]
    Response --> Client
    
```

#### 框架可替換性設計
``` mermaid
graph TB
    subgraph Current ["目前架構 (使用 Gin)"]
        GinApp["gin_application"]
    end
    
    subgraph Core ["核心業務層 (框架無關)"]
        DomainCore["domains"]
        InfraCore["infra"]
    end
    
    subgraph Future ["未來可替換 (例如 Echo)"]
        EchoApp["echo_application"]
    end
    
    GinApp -.->|調用| DomainCore
    EchoApp -.->|調用| DomainCore
    DomainCore --> InfraCore
    
    Replace["🔄 替換 Web 框架<br/>只需修改 Web 框架層<br/>Domain 和 Infrastructure 層無需變動"]
    
    GinApp -.-> Replace
    Replace -.-> EchoApp
    
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
        <td> <a href="./asset/markdown/zap.md" target="_blank">open</a>  </td>
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
        <td><a href="https://github.com/golang-jwt/jwt" target="_blank">golang-jwt</a></td>
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