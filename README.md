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
