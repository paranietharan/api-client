package dto

type Order struct {
	ID    string `json:"id"`
	Items []Item `json:"item"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Qty  string `json:"qty"`
}

type OrderReq struct {
	ID    string    `json:"id"`
	Items []ReqItem `json:"item"`
}

type ReqItem struct {
	Name string `json:"name"`
	Qty  string `json:"qty"`
}
