package airportrepository

import (
	"github.com/gbenga504/travel-assistant/utils/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type AirportSchema struct {
	Id      string   `json:"id" bson:"_id"`
	City    string   `json:"city" bson:"city"`
	State   string   `json:"state" bson:"state"`
	Country string   `json:"country" bson:"country"`
	Codes   []string `json:"codes" bson:"codes"`
	Scorer  *float64 `json:"score" bson:"score"`
}

type AirportRepository struct {
	collection db.Collection
}

func NewAirportRepository(db db.Db) *AirportRepository {
	type AirportIndexModel struct {
		City    string
		State   string
		Country string
	}

	collection := db.Collection("airports")
	collection.CreateOneIndex(AirportIndexModel{City: "text", State: "text", Country: "text"})

	return &AirportRepository{
		collection,
	}
}

func (r *AirportRepository) SearchAirports(searchTerm string) (results []AirportSchema) {
	var airports []AirportSchema

	// Filter
	filter := bson.D{{Key: "$text", Value: bson.M{"$search": searchTerm}}}

	// FindMany Options
	var limit int64 = 5
	projections := bson.D{
		{Key: "city", Value: 1},
		{Key: "state", Value: 1},
		{Key: "country", Value: 1},
		{Key: "codes", Value: 1},
		{Key: "score", Value: bson.M{"$meta": "textScore"}},
	}
	sort := bson.D{{Key: "score", Value: bson.M{"$meta": "textScore"}}}

	opts := db.FindManyOptions{
		Sort:       &sort,
		Limit:      &limit,
		Projection: &projections,
	}

	r.collection.FindMany(
		filter,
		&airports,
		&opts,
	)

	if airports == nil {
		return []AirportSchema{}
	}

	return airports
}
