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
		Port:            "42980",
		MongoHost:       "testapi-jzige.mongodb.net/",
		MongoUser:       "admin",
		MongoPassword:   "admin",
		MongoDatabase:   "Calanderfy",
		MongoCollection: "profile-v2",
	}
}
