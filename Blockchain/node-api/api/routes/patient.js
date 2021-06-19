const express = require('express')
const {contract} = require('../contract')

const routes = express.Router()

routes.post('/register',(req,res)=>{
    contract("INVOKE",["RegisterPatient",req.body.aadhaar,req.body.consenter],(err,payload)=>{
        if (err){
            console.log(err)
            res.status(500).json(err)
        }
        else{
            console.log(payload)
            res.status(200).json({
                "message":"successfully registered"
            })
        }
    })
})

routes.put('/giveconsent',(req,res)=>{
    contract("INVOKE",["UpdateTempConsent",req.body.aadhaar,req.body.type,req.body.to,req.body.till],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json({
                "message":`successfully given consent to ${req.body.to} till ${req.body.till}`
            })
        }
    })
})

routes.put('/permconsent',(req,res)=>{
    contract("INVOKE",["UpdatePermConsent",req.body.aadhaar,req.body.type,req.body.to],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json({
                "message":`successfully given consent to ${req.body.to} till ${req.body.till}`
            })
        }
    })
})

module.exports = routes