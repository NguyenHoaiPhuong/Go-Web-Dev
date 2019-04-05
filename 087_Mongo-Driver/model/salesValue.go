package model

// SalesValue : sale value
type SalesValue struct {
	CustomerID           string  `json:"CustomerID" bson:"CustomerID"`
	ProductID            string  `json:"ProductID" bson:"ProductID"`
	GrossSalesValuePerKg float64 `json:"GrossSalesValuePerKg" bson:"GrossSalesValuePerKg"`
	NetSalesValuePerKg   float64 `json:"NetSalesValuePerKg" bson:"NetSalesValuePerKg"`
}
