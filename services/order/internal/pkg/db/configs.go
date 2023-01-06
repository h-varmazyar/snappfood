package db

type Configs struct {
	MySqlDSN      string `yaml:"mysql_dsn" mapstructure:"mysql_dsn"`
	RedisAddress  string `yaml:"redis_address" mapstructure:"redis_address"`
	RedisPassword string `yaml:"redis_password" mapstructure:"redis_password"`
	RedisLockDB   int    `yaml:"redis_lock_db" mapstructure:"redis_lock_db"`
}
