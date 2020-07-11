package config

type Configs struct {
	LogMaxSize      int
	LogMaxBackups   int
	LogMaxAge       int
	LogFile         string
	Port            string
	MongoHost       string
	MongoUser       string
	MongoPassword   string
	MongoDatabase   string
	MongoCollection string
}

func GetConfigs() Configs {
	return Configs{
		LogMaxSize:      1,
		LogMaxAge:       3,
		LogMaxBackups:   28,
		LogFile:         "/var/log/transact.log",
		Port:            "4900",
		MongoHost:       "testapi-shard-00-00-jzige.mongodb.net:27017",
		MongoUser:       "admin",
		MongoPassword:   "admin",
		MongoDatabase:   "Calanderfy",
		MongoCollection: "profile-v2",
	}
}
