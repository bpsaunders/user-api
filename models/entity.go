package models

// UserDao describes a user database entity
type UserDao struct {
	ID        string `bson:"_id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
	Country   string `bson:"country"`
}
