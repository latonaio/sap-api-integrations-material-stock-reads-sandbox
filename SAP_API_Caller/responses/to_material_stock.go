package responses

type ToMaterialStock struct {
	D struct {
		Metadata struct {
			ID   string `json:"id"`
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		Material                  string `json:"Material"`
		Plant                     string `json:"Plant"`
		StorageLocation           string `json:"StorageLocation"`
		Batch                     string `json:"Batch"`
		Supplier                  string `json:"Supplier"`
		Customer                  string `json:"Customer"`
		WBSElementInternalID      string `json:"WBSElementInternalID"`
		SDDocument                string `json:"SDDocument"`
		SDDocumentItem            string `json:"SDDocumentItem"`
		InventorySpecialStockType string `json:"InventorySpecialStockType"`
		InventoryStockType        string `json:"InventoryStockType"`
		ToMaterialStock           string `json:"to_MaterialStock"`
	} `json:"d"`
}
