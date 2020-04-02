const express = require('express')
const {contract} = require('../contract')
routes = express.Router()

routes.get('/test',(req,res)=>{
    console.log(req.headers)
    contract("QUERY",["GetTest",req.headers.test_id,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json(JSON.parse(payload))
        }
    })
})
routes.get('/report',(req,res)=>{
    console.log(req.headers)
    contract("QUERY",["GetReport",req.headers.report_id,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json(JSON.parse(payload))
        }
    })
})

routes.get('/treatment',(req,res)=>{
    console.log(req.headers)
    contract("QUERY",["GetTreatment",req.headers.treatment_id,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json(JSON.parse(payload))
        }
    })
})

routes.get('/drug',(req,res)=>{
    console.log(req.headers)
    contract("QUERY",["GetDrugs",req.headers.drug_id,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json(JSON.parse(payload))
        }
    })
})
module.exports = routes