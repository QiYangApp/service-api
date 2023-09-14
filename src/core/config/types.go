package config

type Config struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
	Token    `mapstructure:"token"`
	Cache    `mapstructure:"cache"`
}

type Server struct {
	Domain string `mapstructure:"domain"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Env    string `mapstructure:"env"`
}

type DatabaseOperation struct {
	Host     string         `mapstructure:"host"`
	Port     int            `mapstructure:"port"`
	Username string         `mapstructure:"username"`
	Password string         `mapstructure:"password"`
	Database string         `mapstructure:"database"`
	Config   DatabaseConfig `mapstructure:"config"`
}

type DatabaseRedis struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	Database     int    `mapstructure:"database"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
	MaxActive    int    `mapstructure:"maxActive"`
	IdleTimeout  int    `mapstructure:"idleTimeout"`
}

type Database struct {
	Type  string            `mapstructure:"type"`
	Cache DatabaseRedis     `mapstructure:"cache"`
	Read  DatabaseOperation `mapstructure:"read"`
	Write DatabaseOperation `mapstructure:"write"`
}

type DatabaseConfig struct {
	MaxOpenConns int32 `mapstructure:"maxOpenConns"`
	MaxIdleConns int32 `mapstructure:"maxIdleConns"`
}

// token 签名
type Token struct {
	PrivateKey  string `mapstructure:"privateKey"`
	PublicKey   string `mapstructure:"publicKey"`
	Signing     string `mapstructure:"signing"`
	ExpiresTime int    `mapstructure:"expiresTime"`
}

type Cache struct {
	Driver  string     `mapstructure:"driver"`
	Expires int        `mapstructure:"expires"`
	Redis   CacheRedis `mapstructure:"redis"`
}

type CacheRedis struct {
	Database int `mapstructure:"database"`
}
