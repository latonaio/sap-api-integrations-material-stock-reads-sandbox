package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-material-stock-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToMaterialStock(raw []byte, l *logger.Logger) (*MaterialStock, error) {
	pm := &responses.MaterialStock{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to MaterialStock. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &MaterialStock{
        Material:                     data.Material,
        Plant:                        data.Plant,
        StorageLocation:              data.StorageLocation,
        Batch:                        data.Batch,
        Supplier:                     data.Supplier,
        Customer:                     data.Customer,
        WBSElementInternalID:         data.WBSElementInternalID,
        SDDocument:                   data.SDDocument,
        SDDocumentItem:               data.SDDocumentItem,
        InventorySpecialStockType:    data.InventorySpecialStockType,
        InventoryStockType:           data.InventoryStockType,
        MaterialBaseUnit:             data.MaterialBaseUnit,
        MatlWrhsStkQtyInMatlBaseUnit: data.MatlWrhsStkQtyInMatlBaseUnit,
	}, nil
}

func ConvertToToMaterialStock(raw []byte, l *logger.Logger) (*ToMaterialStock, error) {
	tms := &responses.ToMaterialStock{}
	err := json.Unmarshal(raw, tms)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToMaterialStock. unmarshal error: %w", err)
	}

	return &ToMaterialStock{
        Material:                  tms.D.Material,
        Plant:                     tms.D.Plant,
        StorageLocation:           tms.D.StorageLocation,
        Batch:                     tms.D.Batch,
        Supplier:                  tms.D.Supplier,
        Customer:                  tms.D.Customer,
        WBSElementInternalID:      tms.D.WBSElementInternalID,
        SDDocument:                tms.D.SDDocument,
        SDDocumentItem:            tms.D.SDDocumentItem,
        InventorySpecialStockType: tms.D.InventorySpecialStockType,
        InventoryStockType:        tms.D.InventoryStockType,
        ToMaterialStock:           tms.D.ToMaterialStock,
	}, nil
}
