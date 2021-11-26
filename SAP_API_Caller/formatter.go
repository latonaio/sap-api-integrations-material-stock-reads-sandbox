package sap_api_caller

type MaterialStockReads struct {
	 ConnectionKey string `json:"connection_key"`
	 Result        bool   `json:"result"`
	 RedisKey      string `json:"redis_key"`
	 Filepath      string `json:"filepath"`
	 APISchema     string `json:"api_schema"`
	 Material      string `json:"material_code"`
	 Deleted       string `json:"deleted"`
}

type MaterialStock   struct {
     Material        string `json:"Material"`
     Plant           string `json:"Plant"`
     StorageLocation string `json:"StorageLocation"`
     Batch           string `json:"Batch"`
     ToMaterialStock int    `json:"to_MaterialStock"`
}
