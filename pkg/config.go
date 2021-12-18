package pkg

type Config struct {
	MongoDBURI     string `env:"MONGODB_URI"`
	DatabaseName   string `env:"DATABASE_NAME" envDefault:"getir-case-study"`
	CollectionName string `env:"COLLECTION_NAME" envDefault:"records"`
}
