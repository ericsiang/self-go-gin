package env

// mapstructure 是用来讀取 yaml 文件字段名 tag
type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	DBName   string `mapstructure:"dbName" json:"dbName"`
	Username string `mapstructure:"userName" json:"userName"`
	Password string `mapstructure:"password" json:"password"`
}

type MongoDBConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	DBName   string `mapstructure:"dbName" json:"dbName"`
	Username string `mapstructure:"userName" json:"userName"`
	Password string `mapstructure:"password" json:"password"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	DBName   string `mapstructure:"dbName" json:"dbName"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	APP_Mode  string        `mapstructure:"APP_Mode" json:"APP_Mode"`
	Port      int           `mapstructure:"Port" json:"Port"`
	JwtSecret string        `mapstructure:"JwtSecret" json:"JwtSecret"`
	MysqlDB   MysqlConfig   `mapstructure:"Mysql" json:"Mysql"`
	Redis     RedisConfig   `mapstructure:"Redis" json:"Redis"`
	MongoDB   MongoDBConfig `mapstructure:"MongoDB" json:"MongoDB"`
}
