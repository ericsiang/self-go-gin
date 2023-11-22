# self-go-gin
golang gin  framwork   設計自用模板

### 檔案結構
api -> controller
middleware
router 
validate
response
model
migrate
test
log
util
db connect
redis connect

### 使用到的 package
* github.com/spf13/viper  Viper是一個配置設定文件、環境變量
* go.uber.org/zap Zap 是一個快速、結構化、級別化的日誌庫，由 Uber 開發
* github.com/gin-contrib/zap  Gin 框架的 zap 日誌中間件，zap 是一個快速、結構化、級別化的日誌庫
* github.com/lestrrat-go/file-rotatelogs  Go 語言的日誌文件切割和彙整庫
* golang.org/x/crypto/bcrypt 密碼加密核對
* gorm.io/gorm Go 語言 ORM 庫，它支持 MySQL、PostgreSQL、SQLite 和 SQL Server 數據庫
* github.com/go-sql-driver/mysql  MySQL 驅動，連接 MySQL 數據庫
* github.com/dgrijalva/jwt-go  JSON Web Token (JWT) 庫
* github.com/go-playground/validator 驗證器用於驗證結構體和個別的數據
* github.com/gin-contrib/cors 跨域請求的中間件
