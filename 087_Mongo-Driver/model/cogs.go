package model

// COGS : Cost of Good Sold
type COGS struct {
	ProductID     string  `json:"ProductID" bson:"ProductID"`
	ScrapCLPPerKg float64 `json:"Scrap CLP/kg" bson:"Scrap CLP/kg"`
	COGSCLPPerKg  float64 `json:"COGS CLP/kg" bson:"COGS CLP/kg"`
}
