# Create your first serverless function & workflow 


Some commands:

#Creates Resource Group
az group create --resource-group rg-1-serverless-function-and-workflow --subscription 9697bd55-1f87-4128-a7c7-c82bc0afa083 --location centralus

#Creates Storage account.
az storage account create -n <Storage Account Name> -g <Resource Group Name>

#Initializes Function App in directory
func init --worker-runtime custom

#Creates new Azure Function
func new -l Custom -t HttpTrigger -n <Function Name>

#Start Azure Runtime and Function into local sand box
func start

#Create Function App in Azure
az functionapp create --consumption-plan-location centralus --name anandxmj-product-function-app --os-type Linux --resource-group rg-1-serverless-function-and-workflow  --runtime custom --storage-account openhackserverless2022 --functions-version 4

#Publish current directory to Function App
func azure functionapp publish <Azure FunctionApp Name>

Some Notes:
- While creating a Function App, pay attention to properly setting up ostype properly.
- Logic App will not show the Azure Function in Workflow Designer if Custom routing ( “route” attribute in  function json) is used.
- Directory where you run func init is the representative of Function App
- Every function created from that directory with func new becomes /api/<function name> endpoint
- When creating Function App, you need Storage Account. Hence a storage account needs to be created first
 




