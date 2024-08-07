# TrustChain Project
**Internship Project**

TrustChain is a new system for managing software licenses in networked environments. It allows software vendors to delegate the sale, management, and accounting of licenses to smart contracts running on the permissioned Hyperledger Fabric Blockchain. TrustChain leverages the decentralization and security of existing Blockchain infrastructures to manage pay-as-you-grow accounting models in an automated and flexible way without relying on third parties. The system is based on Docker microservices that communicate via REST API with Hyperledger Fabric Blockchain.

## Technologies Used
- Hyperledger Fabric
- Golang
- Docker

## Starting the Project

### What is Hyperledger Fabric?
[Hyperledger Fabric Documentation](https://hyperledger-fabric.readthedocs.io/en/latest/)

### Installation Guide
**Please refer only to the official guide:**
[Getting Started with Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/latest/getting_started.html)

**Note:** Ensure that Golang is installed correctly as indicated on the official Golang site, e.g., `/go/src/github.com/github-username/repo_where_we_are_working`.

**This project was tested & developed on Linux Ubuntu 18.04.**

**We will simulate the consumer and producer of the 5G network with two terminals on the same network for simplicity.**

## Starting and Monitoring the Network

1. **Go to the `commercial-paper` directory and start a terminal to monitor the network:**

    ```sh
    cd commercial-paper
    ./network-starter.sh
    /organization/magnetocorp/configuration/cli/monitordocker.sh fabric_test
    ```

## Setup the First Node (Producer)

### Terminal 2 - Magnetocorp

2. **Go to `commercial-paper/organization/magnetocorp` directory and start `magnetocorp.sh`:**

    ```sh
    cd /go/src/github.com/marco-spagnuolo/trust-chains/commercial-paper/organization/magnetocorp
    ./magnetocorp.sh
    ```

## Setup the Second Node (Consumer)

### Terminal 3 - Digibank

3. **Go to `commercial-paper/organization/digibank` directory and start `digibank.sh`:**

    ```sh
    cd /go/src/github.com/marco-spagnuolo/trust-chains/commercial-paper/organization/digibank
    ./digibank.sh
    ```

### Terminal 2 - Magnetocorp

4. **Start the Magnetocorp node:**

    ```sh
    ./magnetocorpStart.sh
    ```

### Terminal 3 - Digibank

5. **Start the Digibank node:**

    ```sh
    ./digibankStart.sh
    ```

### Terminal 2 - Magnetocorp

6. **Commit the transaction:**

    ```sh
    ./commit.sh
    ```

**Ready to go!**

## Submitting an Issue Example

### Terminal 3 - Digibank

7. **Run the application:**

    ```sh
    cd go2
    go run app.go id
    ```

**Webserver with REST is on – now wait for GET/POST requests (using POSTMAN is suggested).**

### Terminal 3 - Digibank

8. **Run the application:**

    ```sh
    cd go2
    go run app.go id
    ```

**Webserver with REST is on – now wait for GET/POST requests (using POSTMAN is suggested).**

For more information, see the PPTX file in this repository.

## Troubleshooting

- **Error: failed to get network: Failed to create new channel client: event service creation failed: could not get chConfig cache reference: QueryBlockConfig failed: QueryBlockConfig failed: queryChaincode failed: Transaction processing for endorser [localhost:9051]: gRPC Transport Status Code: (2) Unknown. Description: error validating proposal: access denied: channel [mychannel] creator org [Org2MSP]**

    **Solution:** Delete only the `wallet` directory in `IDENTITY/USER/OPERATOR/WALLET`.

- **Error: failed to populate wallet contents: keystore folder should contain one file**

    **Solution:** Delete the second key in the directory of the relevant organization (e.g., Org1 = Digibank):

    ```sh
    /trust-chains/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore
    ```


## Additional Setup and Running the Network

### Detailed Steps for Running the Network

1. **Navigate to the `commercial-paper` directory and monitor the network:**

    ```sh
    cd commercial-paper
    ./network-starter.sh
    /organization/magnetocorp/configuration/cli/monitordocker.sh fabric_test
    ```

2. **Setting up the Producer Node (Magnetocorp):**

    **Terminal 2:**

    ```sh
    cd /go/src/github.com/marco-spagnuolo/trust-chains/commercial-paper/organization/magnetocorp
    ./magnetocorp.sh
    ./magnetocorpStart.sh
    ./commit.sh
    ```

3. **Setting up the Consumer Node (Digibank):**

    **Terminal 3:**

    ```sh
    cd /go/src/github.com/marco-spagnuolo/trust-chains/commercial-paper/organization/digibank
    ./digibank.sh
    ./digibankStart.sh
    ```

### Submitting an Issue Example

**Terminal 3 - Digibank:**

1. **Run the application in the `go2` directory:**

    ```sh
    cd go2
    go run app.go id
    ```

    **Webserver with REST is on – now wait for GET/POST requests (using POSTMAN is suggested).**

### Additional Resources

- For more detailed instructions, refer to the PPTX file included in the repository.
- Always ensure that you are using the correct directories and following the sequence of commands precisely to avoid setup issues.

## Troubleshooting Guide

- **Error: failed to get network: Failed to create new channel client: event service creation failed: could not get chConfig cache reference: QueryBlockConfig failed: QueryBlockConfig failed: queryChaincode failed: Transaction processing for endorser [localhost:9051]: gRPC Transport Status Code: (2) Unknown. Description: error validating proposal: access denied: channel [mychannel] creator org [Org2MSP]**

    **Solution:** Delete only the `wallet` directory in `IDENTITY/USER/OPERATOR/WALLET`.

    ```sh
    rm -rf /trust-chains/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore
    ```

- **Error: failed to populate wallet contents: keystore folder should contain one file**

    **Solution:** Delete the second key in the directory of the relevant organization (e.g., Org1 = Digibank):

    ```sh
    rm -rf /trust-chains/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/*
    ```

### Common Issues and Solutions

- **Issue:** Network monitoring script fails to start.

    **Solution:** Ensure Docker is running and all necessary containers are up. Restart Docker if necessary and rerun the network starter script.

- **Issue:** Smart contract deployment fails.

    **Solution:** Check the logs for detailed error messages, ensure all dependencies are correctly installed, and verify that the environment variables are set properly.

- **Issue:** REST API does not respond.

    **Solution:** Confirm that the web server is running, check network configurations, and use tools like POSTMAN to send test requests. Ensure that ports are not blocked by a firewall.

## Final Notes

- Regularly consult the official Hyperledger Fabric documentation for updates and best practices.
- Maintain a clean and organized directory structure to avoid path-related issues.
- Keep your system and all dependencies updated to the latest stable versions to ensure compatibility and security.

For any additional help or to report issues, refer to the project's issue tracker on GitHub.
