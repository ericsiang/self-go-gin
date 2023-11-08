package initialize

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
	Port    int           `mapstructure:"port" json:"port"`
	MysqlDB MysqlConfig   `mapstructure:"mysql" json:"mysql"`
	Redis   RedisConfig   `mapstructure:"redis" json:"redis"`
	MongoDB MongoDBConfig `mapstructure:"mongodb" json:"mongodb"`
}