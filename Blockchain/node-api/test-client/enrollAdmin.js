const {Wallets} = require('fabric-network')
const fs = require('fs')
const WALLETPATH = './wallet'
const pk = fs.readFileSync('./Admin/priv_sk')
const cert = fs.readFileSync('./Admin/Admin@devorg-cert.pem').toString()

// console.log(cert)
// console.log(pk)

const main = async ()=>{
    try {
        const wallet = await Wallets.newFileSystemWallet(WALLETPATH)
        const x509Identity = {
            credentials: {
                certificate: cert,
                privateKey: pk.toString(),
            },
            mspId: 'DevMSP',
            type: 'X.509',
        }
        await wallet.put("admin",x509Identity)
        console.log(`Successfully enrolled user admin and imported it into the wallet`)
    } catch (error) {
        console.log(error)
    }
}
main()