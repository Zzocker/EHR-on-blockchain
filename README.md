# Electronic Health Record System on Blockchain
---

This project leverages a hyperledger fabric (v2.0) network maintained by hospital to store a medical record of the patient securely with keeping the patient at the center. That means the medical record of any patient cannot be accessed without the consent of his/her. All the participants of the network like doctors, hospitals, pharmacies and private clinics will be given digital certificates by the Offical Medical board of the country to join the network. Doctors will able to perform all the CRUD operations on their patient records.

This project also helps in cries like the COVID-19 outbreak by providing correct and auditable certificates to the People. This also keeps track of infected during the time of cries/pandemic only.

## Installation

### Install Docker
    > chmod +x docker.sh
    > sudo ./docker.sh
    > usermod -a -G docker ${USER}

### Start Local Test Fabric Network
    > cd Blockchain/test-network
    > docker-compose up -d
    > docker exec -it cli bash
    > cd channel-artifacts && ./joinchannel.sh
    > cd $GOPATH/src/chaincode && go mod vendor
    > cd .. && peer lifecycle chaincode package health.tar.gz --label health -p chaincode
    > peer lifecycle chaincode install health.tar.gz 
    > export CC_PACKAGE=health:----------64hexdigit number----------
    > peer lifecycle chaincode approveformyorg -C test -n health --package-id $CC_PACKAGE -v 1.0 -o orderer:7050 --sequence 1
    > peer lifecycle chaincode commit -C test -n health -v 1.0 -o orderer:7050 --sequence 1

### Start Hyperledger Explorer to view blocks
    > docker-compose up -d




