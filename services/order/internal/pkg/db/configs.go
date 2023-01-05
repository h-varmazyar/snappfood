package db

type Configs struct {
	MySqlDSN      string `yaml:"postgres_dsn"`
	RedisAddress  string `yaml:"REDIS_ADDRESS"`
	RedisPassword string `yaml:"REDIS_PASSWORD"`
	RedisLockDB   int    `yaml:"REDIS_LOCK_DB"`
}
