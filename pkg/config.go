package pkg

type Config struct {
	MongoDBURI     string `env:"MONGODB_URI"`
	CollectionName int    `env:"COLLECTION_NAME" envDefault:"records"`
}
