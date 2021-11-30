package responses

type MaterialStock struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
		Material        string   `json:"Material"`
		Plant           string   `json:"Plant"`
		StorageLocation string   `json:"StorageLocation"`
		Batch           string   `json:"Batch"`
		ToMaterialStock string   `json:"to_MaterialStock"`
		} `json:"results"`
	} `json:"d"`
}
