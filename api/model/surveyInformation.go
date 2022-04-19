package model

type StudentDataStructure struct {
	EMAIL    string `json:"Email" bson:"Email"`
	USERNAME string `json:"Username" bson:"Username"`
	PASSWORD string `json:"Password" bson:"Password"`
}

type LoginDataStructure struct {

	USERNAME string `json:"Username" bson:"Username"`
	PASSWORD string `json:"Password" bson:"Password"`
}
