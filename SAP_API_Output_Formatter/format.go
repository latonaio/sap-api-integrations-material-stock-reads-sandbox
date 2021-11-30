package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-material-stock-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToMaterialStock(raw []byte, l *logger.Logger) *MaterialStock {
	pm := &responses.MaterialStock{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &MaterialStock{
		Material:            data.Material,
		Plant:               data.Plant,
		StorageLocation:     data.StorageLocation,
		Batch:               data.ValidityStartDate,
		ToMaterialStock:     data.ToMaterialStock,
	}
}
