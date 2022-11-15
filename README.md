# Trust-chains
**Internship project** 


TrustChain is a new system for managing software licenses in networked environments that allows software vendors to delegate the sale, 
management and accounting of licenses to smart contracts running on the permissioned Fabric's Blockchain.
TrustChain allows you to leverage the decentralization and security of existing Blockchain infrastructures to manage
pay-as-you-grow accounting models in an automated and flexible way without relying on third parties. 
The system is based on Docker microservices that communicate via REST API with Hyperledger Fabric Blockchain.

## Technologies used:
-Hyperledger Fabric
-Golang
-Docker

# Start Project :
__What is hyperledger Fabric?__

https://hyperledger-fabric.readthedocs.io/en/latest/

__Installation guide__ 

https://hyperledger-fabric.readthedocs.io/en/latest/getting_started.html

***I advise you to install golang correctly as indicated on the official golang site like this /go/src/github.com/github-username/repo_where_we_are_working***

**This project was tested & developed on Linux Ubuntu 18.04**

**We will simulate the consumer and the producer of the 5g network with two terminals on the same network for simplicity.**




## Start network & monitor network 

**go to commercial-paper directory & start a terminal to monitor network **

cd commercial-paper
 ./network-starter.sh
 /organization/magnetocorp/configuration/cli/monitordocker.sh fabric_test

## Setup the first node(Producer):

### Terminal 2 Magnetocorp

**go to commercial-paper/organization/magnetocorp directory & start magnetocorp.sh**

cd /go/src/github.com/marco-spagnuolo/trust-chains/commercial-paper/organization/magnetocorp

  ./magnetocorp.sh

## Setup the second node(Consumer):

### Terminale 3 digibank 

**go to commercial-paper/organization/digibank directory & start digibank.sh**

cd /go/src/github.com/marco-spagnuolo/trust-chains/commercial-paper/organization/digibank

 ./digibank.sh

### Terminal 2 Magnetocorp 
 **Start Magnetocorp node**
 
 ./magnetocorpStart.sh

### Terminal 3 digibank 
**Start Digibank node**

./digibankStart.sh

### Terminal 2 Magnetocorp

./commit.sh

**Ready to go**
## SUBMIT ISSUE example :

cd go2 
go run app.go id 

webserver with REST is on  --> now wait for get/post request (I suggest to use POSTMAN)


**Terminale 3 digibank** 

cd go2 
go run app.go id 

**webserver with REST is on  --> now wait for get/post request (I suggest to use POSTMAN)**

__see pptx file in this repo for more info__


## Troubleshooting: 

error: failed to get network: Failed to create new channel client: event service creation failed: could not get chConfig cache reference: QueryBlockConfig failed: QueryBlockConfig failed: queryChaincode failed: Transaction processing for endorser [localhost:9051]: gRPC Transport Status Code: (2) Unknown. Description: error validating proposal: access denied: channel [mychannel] creator org [Org2MSP]\
exit status 1
-->
**Delete only the wallet dir IN IDENTITY/USER/OPERATOR/WALLET**

error: failed to populate wallet contents: keystore folder should have contain one file\
exit status 1
-->
**DELETE THE SECOND KEY IN THE DIRECTORY OF THE RELEVANT ORG (IN THE EXAMPLE ORG1 = DIGIBANK)
/trust-chains/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore**
