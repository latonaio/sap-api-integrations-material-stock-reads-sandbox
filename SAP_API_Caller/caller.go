package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-material-stock-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetMaterialStock(material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType string) {
	wg := &sync.WaitGroup{}

	wg.Add(2)
	func() {
		c.MaterialStock(material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType)
		wg.Done()
	}()
	func() {
		c.ToMaterialStock(material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) MaterialStock(material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType string) {
	data, err := c.callMaterialStockSrvAPIRequirementMaterialStock("A_MatlStkInAcctMod", material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callMaterialStockSrvAPIRequirementMaterialStock(api, material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType string) (*sap_api_output_formatter.MaterialStock, error) {
	url := strings.Join([]string{c.baseURL, "API_MATERIAL_STOCK_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithMaterialStock(req, material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToMaterialStock(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ToMaterialStock(material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType string) {
	data, err := c.callMaterialStockSrvAPIRequirementToMaterialStock(fmt.Sprintf("A_MatlStkInAcctMod(Material='%s',Plant='%s',StorageLocation='%s',Batch='%s',Supplier='%s',Customer='%s',WBSElementInternalID='%s',SDDocument='%s',SDDocumentItem='%s',InventorySpecialStockType='%s',InventoryStockType='%s')/to_MaterialStock", material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType))
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callMaterialStockSrvAPIRequirementToMaterialStock(api string) (*sap_api_output_formatter.ToMaterialStock, error) {
	url := strings.Join([]string{c.baseURL, "API_MATERIAL_STOCK_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToMaterialStock(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithMaterialStock(req *http.Request, material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s' and StorageLocation eq '%s' and Batch eq '%s' and Supplier eq '%s' and Customer eq '%s' and WBSElementInternalID eq '%s' and SDDocument eq '%s' and SDDocumentItem eq '%s' and InventorySpecialStockType eq '%s' and InventoryStockType eq '%s'", material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType))
	req.URL.RawQuery = params.Encode()
}
