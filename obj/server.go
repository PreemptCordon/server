package obj

type ServerSettings struct {
	PostgresURI  string `mapstructure:"PostgresURI"`
	RedisURI     string `mapstructure:"RedisURI"`
	Domain       string `mapstructure:"Domain"`
	Country      string
	Languages    string
	EmailUser    string `mapstructure:"EmailUser"`
	SMTPPassword string `mapstructure:"SMTPPassword"`
	SMTPHost     string `mapstructure:"SMTPHost"`
	SMTPPort     string `mapstructure:"SMTPPort"`
	AppKey       string `mapstructure:"AppKey"`
}
