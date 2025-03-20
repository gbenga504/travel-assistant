package db

type Db interface {
	Close()
	Collection(collection string) Collection
}

type Collection interface {
	CreateOne()
}
