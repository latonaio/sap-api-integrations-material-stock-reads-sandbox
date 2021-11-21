package main

import (
	sap_api_caller "sap-api-integrations-material-stock-reads/SAP_API_Caller"
	"sap-api-integrations-material-stock-reads/file_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := file_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Material_Stock.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetMaterialStock(
		inoutSDC.Material.Plant.StorageLocation.Batch,
	)
}
