# sap-api-integrations-material-stock-reads  
sap-api-integrations-material-stock-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で品目在庫データを取得するマイクロサービスです。  
sap-api-integrations-material-stock-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-material-stock-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_MATERIAL_STOCK_SRV/overview  

## 動作環境
sap-api-integrations-material-stock-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。 
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）  
・ AION のリソース （推奨)  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-material-stock-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-material-stock-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_MATERIAL_STOCK_SRV/overview  
* APIサービス名(=baseURL): API_MATERIAL_STOCK_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-material-stock-reads には、次の API をコールするためのリソースが含まれています。  

* A_MatlStkInAcctMod（品目在庫データ）

## API への 値入力条件 の 初期値
sap-api-integrations-material-stock-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.MaterialStock.Material（品目）
* inoutSDC.MaterialStock.Plant（プラント）
* inoutSDC.MaterialStock.StorageLocation（保管場所）
* inoutSDC.MaterialStock.Batch（ロット）
* inoutSDC.MaterialStock.Supplier（仕入先）
* inoutSDC.MaterialStock.Customer（得意先）
* inoutSDC.MaterialStock.WBSElementInternalID（WBS要素）
* inoutSDC.MaterialStock.SDDocument（販売伝票）
* inoutSDC.MaterialStock.SDDocumentItem（販売伝票明細）
* inoutSDC.MaterialStock.InventorySpecialStockType（特殊在庫タイプ）
* inoutSDC.MaterialStock.InventoryStockType（在庫タイプ）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"MaterialStock" が指定されています。    
  
```
	"api_schema": "/A_MatlStkInAcctMod",
	"accepter": ["MaterialStock"],
	"material_code": "FG29",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "/A_MatlStkInAcctMod",
	"accepter": ["All"],
	"material_code": "FG29",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetMaterialStock(material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "MaterialStock":
			func() {
				c.MaterialStock(material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType)
				wg.Done()
			}()
		case "ToMaterialStock":
			func() {
				c.ToMaterialStock(material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType)
				wg.Done()
			}()

		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 品目在庫データ が取得された結果の JSON の例です。  
以下の項目のうち、"Batch" ～ "WBSElementInternalID" は、/SAP_API_Output_Formatter/type.go 内 の Type Product {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"Batch": "0000000109",
	"Customer": "",
	"InventorySpecialStockType": "",
	"InventoryStockType": "01",
	"Material": "FG29",
	"MaterialBaseUnit": "BT",
	"MatlWrhsStkQtyInMatlBaseUnit": "100",
	"Plant": "1710",
	"SDDocument": "",
	"SDDocumentItem": "0",
	"StorageLocation": "171A",
	"Supplier": "",
	"WBSElementInternalID": "000000000000000000000000",
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-material-stock-reads/SAP_API_Caller/caller.go#L50",
	"function": "sap-api-integrations-material-stock-reads/SAP_API_Caller.(*SAPAPICaller).MaterialStock",
	"level": "INFO",
	"time": "2021-12-02T17:56:29.953557+09:00"
}
```
