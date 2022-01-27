package main

import (
	sap_api_caller "sap-api-integrations-material-stock-reads/SAP_API_Caller"
	"sap-api-integrations-material-stock-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Material_Stock_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"MaterialStock", "ToMaterialStock",
		}
	}

	caller.AsyncGetMaterialStock(
		inoutSDC.MaterialStock.Material,
		inoutSDC.MaterialStock.Plant,
		inoutSDC.MaterialStock.StorageLocation,
		inoutSDC.MaterialStock.Batch,
		inoutSDC.MaterialStock.Supplier,
		inoutSDC.MaterialStock.Customer,
		inoutSDC.MaterialStock.WBSElementInternalID,
		inoutSDC.MaterialStock.SDDocument,
		inoutSDC.MaterialStock.SDDocumentItem,
		inoutSDC.MaterialStock.InventorySpecialStockType,
		inoutSDC.MaterialStock.InventoryStockType,
		accepter,
	)
}
