package config

type DefaultConfig struct {
	Apps            Apps            `mapstructure:"apps"`
	Server          Server          `mapstructure:"server"`
	GRPCServer      GRPCServer      `mapstructure:"GRPCServer"`
	PasswordHashing PasswordHashing `mapstructure:"passwordHashing"`
	JWT             JWT             `mapstructure:"jwt"`
}

type Apps struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type GRPCServer struct {
	UserMicroservice UserMicroservice `mapstructure:"userMicroservice"`
}

type UserMicroservice struct {
	BaseURL string `mapstructure:"baseURL"`
}

type PasswordHashing struct {
	HashSalt int `mapstructure:"hashSalt"`
}

type JWT struct {
	SecretKey string `mapstructure:"secretKey"`
}
