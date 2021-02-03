package edenrepo

type StockTransferRequest struct {
	RecognitionDate    string                     `json:"recognition_date" valid:"required"`
	OriginID           string                     `json:"origin_id" valid:"required"`
	DestinationID      string                     `json:"destination_id" valid:"required"`
	DestinationAddress string                     `json:"destination_address"`
	EtaDate            string                     `json:"eta_date" valid:"required"`
	EtaTime            string                     `json:"eta_time" valid:"required"`
	Note               string                     `json:"note"`
	TotalCost          float32                    `json:"total_cost" valid:"required"`
	StockTransferItems []StockTransferItemRequest `json:"stock_transfer_items" valid:"required"`
}

type StockTransferItemRequest struct {
	ID          string  `json:"id"`
	GroupID     string  `json:"group_id" valid:"required"`
	DeliveryQty float32 `json:"deliver_qty" valid:"required"`
	ReceiveQty  float32 `json:"receive_qty"`
	ReceiveNote string  `json:"receive_note"`
	Note        string  `json:"note"`
	ItemWeight  float32 `json:"item_weight"`
	UnitCost    float32 `json:"unit_cost" valid:"required"`
	Subtotal    float32 `json:"subtotal" valid:"required"`
}
