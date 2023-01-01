package main

type ServerSettings struct {
	PostgresURI  string `mapstructure:"PostgresURI"`
	RedisURI     string `mapstructure:"RedisURI"`
	Domain       string `mapstructure:"Domain"`
	Country string
	Languages string
	EmailUser    string `mapstructure:"EmailUser"`
	SMTPPassword string `mapstructure:"SMTPPassword"`
	SMTPHost     string `mapstructure:"SMTPHost"`
	SMTPPort     string `mapstructure:"SMTPPort"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper..AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
ServerConfig := LoadConfig(".")