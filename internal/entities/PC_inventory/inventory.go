package entity

type Inventory struct {
	Dep_name        string `json:"dep_name"`
	Inventory_id    int    `json:"inventory_id"`
	Fk_dep_id       int    `json:"fk_dep_id"`
	Full_name       string `json:"full_name"`
	Pc_id           string `json:"pc_id"`
	Inventory_num   string `json:"inventory_num"`
	Invent_monitors string `json:"invent_monitors"`
	Invent_printer  string `json:"invent_printer"`
	Invent_mfu      string `json:"invent_mfu"`
	Invent_laptop   string `json:"invent_laptop"`
	Invent_other    string `json:"invent_other"`
}
