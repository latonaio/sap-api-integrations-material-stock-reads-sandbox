package responses

type MRPArea struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product                       string  `json:"Product"`
			Plant                         string  `json:"Plant"`
			MRPArea                       string  `json:"MRPArea"`
			MRPType                       string  `json:"MRPType"`
			MRPResponsible                string  `json:"MRPResponsible"`
			MRPGroup                      string  `json:"MRPGroup"`
			ReorderThresholdQuantity      string  `json:"ReorderThresholdQuantity"`
			PlanningTimeFence             int     `json:"PlanningTimeFence"`
			LotSizingProcedure            string  `json:"LotSizingProcedure"`
			LotSizeRoundingQuantity       float64 `json:"LotSizeRoundingQuantity"`
			MinimumLotSizeQuantity        float64 `json:"MinimumLotSizeQuantity"`
			MaximumLotSizeQuantity        float64 `json:"MaximumLotSizeQuantity"`
			MaximumStockQuantity          float64 `json:"MaximumStockQuantity"`
			ProcurementSubType            string  `json:"ProcurementSubType"`
			DfltStorageLocationExtProcmt  string  `json:"DfltStorageLocationExtProcmt"`
			MRPPlanningCalendar           string  `json:"MRPPlanningCalendar"`
			SafetyStockQuantity           string  `json:"SafetyStockQuantity"`
			SafetyDuration                string  `json:"SafetyDuration"`
			FixedLotSizeQuantity          string  `json:"FixedLotSizeQuantity"`
			PlannedDeliveryDurationInDays int     `json:"PlannedDeliveryDurationInDays"`
			StorageLocation               string  `json:"StorageLocation"`
			IsMarkedForDeletion           string  `json:"IsMarkedForDeletion"`
		} `json:"results"`
	} `json:"d"`
}
