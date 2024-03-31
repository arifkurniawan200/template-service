package config

type Config struct {
	DB          DatabaseConfig   `yaml:"db"`
	Cache       CacheConfig      `yaml:"cache"`
	API         APIConfig        `yaml:"api"`
	GRPC        GRPCConfig       `yaml:"grpc"`
	Environment EnvConfig        `yaml:"env"`
	ThirdParty  ThirdPartyConfig `yaml:"third-party"`
}

type DatabaseConfig struct {
	Driver     string `yaml:"driver"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Name       string `yaml:"name"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	ActivePool bool   `yaml:"active_pool"`
	MaxPool    int    `yaml:"max_pool"`
	MinPool    int    `yaml:"min_pool"`
}

type CacheConfig struct {
	Driver string `yaml:"driver"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
}

type APIConfig struct {
	Port int `yaml:"port"`
}

type GRPCConfig struct {
	Port int `yaml:"port"`
}

type EnvConfig struct {
	SecretKey  string `yaml:"secret_key"`
	EncryptKey string `yaml:"encrypt_key"`
}

type ThirdPartyConfig struct {
	UserService UserServiceConfig `yaml:"user-service"`
}

type UserServiceConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
