const {Wallets} = require('fabric-network')
const fabricCAService =require('fabric-ca-client')
const fs = require('fs')
const WALLETPATH = './wallet'

// console.log(cert)
// console.log(pk)

const main = async ()=>{
    try {
        const wallet = await Wallets.newFileSystemWallet(WALLETPATH)
        const admin = await wallet.get('admin')
        if (admin){
            console.log("Admin is already enrolled")
            return
        }

        const ca = new fabricCAService("http://localhost:7054")
        
        const enrollment = await ca.enroll({enrollmentID:"admin",enrollmentSecret:'adminpw'},)
        const x509Identity = {
            credentials: {
                certificate: enrollment.certificate,
                privateKey: enrollment.key.toBytes(),
            },
            mspId: 'DevMSP',
            type: 'X.509',
        }
        await wallet.put("admin",x509Identity)
        console.log(`Successfully enrolled admin and imported it into the wallet`)
    } catch (error) {
        console.log(error)
    }
}
main()