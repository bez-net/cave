package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person is the data structure that we will save and receive.
type Person struct {
	ID        primitive.ObjectID     	`json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName  	string               	`json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName   	string                 	`json:"last_name,omitempty" bson:"last_name,omitempty"`
	MiddleName 	string				 	`json:"middle_name,omitempty" bson:"middle_name,omitempty"`
	Username 	string				 	`json:"Username,omitempty" bson:"Username,omitempty"`
	Email      	string                 	`json:"email,omitempty" bson:"email,omitempty"`
	DOB			string                 	`json:"dob,omitempty" bson:"dob,omitempty"`
  	Gender		string                 	`json:"gender,omitempty" bson:"gender,omitempty"`
  	Address		string                 	`json:"address,omitempty" bson:"address,omitempty"`
  	City		string                 	`json:"city,omitempty" bson:"city,omitempty"`
  	State		string                 	`json:"state,omitempty" bson:"state,omitempty"`
  	Country		string                 	`json:"country,omitempty" bson:"country,omitempty"`
  	Zip			string                 	`json:"zip,omitempty" bson:"zip,omitempty"`
  	PhoneNo		string                 	`json:"phone_no,omitempty" bson:"phone_no,omitempty"`
  	Nationality	string                 	`json:"nationality,omitempty" bson:"nationality,omitempty"`
  	Profession	string                 	`json:"profession,omitempty" bson:"profession,omitempty"`
	Data       	map[string]interface{} 	`json:"data,omitempty" bson:"data,omitempty"` // data is a optional fields that can hold anything in key:value format.
}

// NewPerson will return a Person{} instance, Person structure factory function
func NewPerson(firstName, lastName, userName, email string, data map[string]interface{}) *Person {
	return &Person{
		Username:  userName,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Data:      data,
	}
}
