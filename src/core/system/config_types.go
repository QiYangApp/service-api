package system

type Config struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type Database struct {
	Type     string         `mapstructure:"type"`
	Host     string         `mapstructure:"host"`
	Port     int            `mapstructure:"port"`
	Username string         `mapstructure:"username"`
	Password string         `mapstructure:"password"`
	Database string         `mapstructure:"database"`
	Config   DatabaseConfig `mapstructure:"config"`
}

type DatabaseConfig struct {
	MaxOpenConns int32 `mapstructure:"maxOpenConns"`
	MaxIdleConns int32 `mapstructure:"maxIdleConns"`
}

// token 签名

type TokenConfig struct {
	PrivateKey string `mapstructure:"privateKey"`
	PublicKey  string `mapstructure:"publicKey"`
	Method     string `mapstructure:"method"`
	Signing    string `mapstructure:"signing"`
}
