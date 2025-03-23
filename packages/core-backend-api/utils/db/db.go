package db

type Db interface {
	Close()
	Collection(collection string) Collection
}

type Collection interface {
	CreateOne(document interface{})
	FindMany(filter interface{}, documents interface{})
}
