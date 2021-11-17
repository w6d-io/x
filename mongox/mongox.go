package mongox

type Mongo struct {
	// Username for database authentication
	Username string `mapstructure:"username"`
	// Password for database authentication
	Password string `mapstructure:"password"`
	// Name of the database
	Name string `mapstructure:"name"`
	// URL of the database
	URL string `mapstructure:"url"`
	// URL of the database
	AuthSource string `mapstructure:"authSource"`
}
