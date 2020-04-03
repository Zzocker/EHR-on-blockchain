
peer channel create -f ./channel.tx -o orderer:7050 -c test
peer channel join -b test.block

echo $(peer channel list)