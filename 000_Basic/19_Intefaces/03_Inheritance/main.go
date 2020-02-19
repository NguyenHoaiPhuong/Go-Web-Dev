package main

// IKPI : KPI interface
type IKPI interface {
	init()
	// calculate()
	getKPI() *KPI
}

// KPI structure
type KPI struct {
	Name  string  `json:"Name" bson:"Name"`
	Value float64 `json:"Value" bson:"Value"`
}

func (kpi *KPI) getKPI() *KPI {
	return kpi
}

// InventoryCostKPI : Inventory Cost (in $) KPI
type InventoryCostKPI struct {
	KPI
}

func (kpi *InventoryCostKPI) init() {
	kpi.Name = "InventoryCost"
}

// ScrapCostKPI : Scrap Cost (in $) KPI
type ScrapCostKPI struct {
	KPI
}

func main() {
	var invKPI InventoryCostKPI
	invKPI.Name = "Inventory"
	invKPI.Value = 100

	var scrapKPI ScrapCostKPI
	scrapKPI.Name = "Scrap"
	scrapKPI.Value = 200

	KPIs := make(map[string][]IKPI)
	KPIs["Inventory"] = append(KPIs["Inventory"], &invKPI)
}
