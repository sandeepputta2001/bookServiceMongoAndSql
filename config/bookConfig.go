package config

import "os"

type Config struct {
	MongoURI string 
	DatabaseName string 
	SqlURI string
	RepoType string 
}


func  GetConfig() *Config { 

	return &Config{
		MongoURI: GetEnv("MONGO_URI","mongodb://localhost:27017"),
		DatabaseName: GetEnv("DATABASE_NAME","booksstore"),
		SqlURI: GetEnv("SQL_URI",""),
		RepoType: GetEnv("REPO_TYPE","mongo"), 
	}
}

func GetEnv(key string , defaultval string ) string { 

	value , exists := os.LookupEnv(key)
	if exists{
		return value
	} 

	return defaultval

}