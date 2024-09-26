package entity

type Inventory struct {
	Dep_id          string `json:"fk_dep_id"`
	Inventory_id    int    `json:"inventory_id"`
	Workplace_num   string `json:"workplace_num"`
	Full_name       string `json:"full_name"`
	Pc_id           string `json:"pc_id"`
	Inventory_num   string `json:"inventory_num"`
	Invent_monitors string `json:"invent_monitors"`
	Invent_printer  string `json:"invent_printer"`
	Invent_mfu      string `json:"invent_mfu"`
	Invent_laptop   string `json:"invent_laptop"`
	Invent_other    string `json:"invent_other"`
}
