const express = require('express')
const {contract} = require('../contract')
routes = express.Router()

routes.put('/givedrugs',(req,res)=>{
    console.log(req.body)
    contract("INVOKE",["GiveDrugs",req.headers.drug_id],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json({
                "message":"DONE"
            })
        }
    })
})

module.exports = routes
