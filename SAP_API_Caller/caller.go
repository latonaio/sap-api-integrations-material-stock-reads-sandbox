package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
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

func (c *SAPAPICaller) AsyncGetMaterialStock(Material, Plant, StorageLocation, Batch string) {
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		c.MaterialStockStorageLocation(Material, Plant, StorageLocation)
		wg.Done()
	}()
	go func() {
		c.MaterialStockBatch(Material, Plant, StorageLocation, Batch)
		wg.Done()
	}()
	
	wg.Wait()
}

func (c *SAPAPICaller) MaterialStockStorageLocation(Material, Plant, StorageLocation string) {
	res, err := c.callMaterialStockSrvAPIRequirementBatch("A_MaterialStock('{Material}')/to_MatlStkInAcctMod", Material, Plant, StorageLocation)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) MaterialStockBatch(Material, Plant, StorageLocation, Batch string) {
	res, err := c.callMaterialStockSrvAPIRequirementBatch("A_MaterialStock('{Material}')/to_MatlStkInAcctMod", Material, Plant, StorageLocation, Batch)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callMaterialStockSrvAPIRequirement(api, Material, Plant, StorageLocation, Batch string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_MATERIAL_STOCK_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Material, Plant, StorageLocation, Batch")
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s' and StorageLocation eq '%s' and Batch eq '%s'", Material, Plant, StorageLocation, Batch))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}