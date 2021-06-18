const fs = require('fs')
const yaml = require('js-yaml')
const path = require('path')
const {Wallets,Gateway} = require('fabric-network')

const CONNECTION_PROFILE_PATH = path.resolve(__dirname,"../test-client/connection.yaml")
const WALLET_PATH= path.resolve(__dirname,"../test-client/wallet")

const IDENTITY_NAME = "client"
const CHANNEL_NAME = "test"
const CONTRACT_NAME="health"


const contract =  async (type,inputs,callback) =>{
   const gateway = new Gateway()
       try {
         const ccp = yaml.safeLoad(fs.readFileSync(CONNECTION_PROFILE_PATH))
         const wallet = await Wallets.newFileSystemWallet(WALLET_PATH)
         await gateway.connect(ccp,{wallet:wallet,identity:IDENTITY_NAME,discovery: { enabled: false, asLocalhost: true }})
         const network = await gateway.getNetwork(CHANNEL_NAME)
         const contract = network.getContract(CONTRACT_NAME)
         var res
         if (type == "INVOKE"){
            res = await contract.submitTransaction(...inputs)
         } else if (type == "QUERY"){
            res = await contract.evaluateTransaction(...inputs)
         }
         return callback(null,res)
      } catch (error) {
         //  callback(error,null)
          return callback(error,null)
       } finally{
          gateway.disconnect()
       }
}
module.exports = {contract}