const express = require('express')
const {contract} = require('../contract')

const routes = express.Router()

routes.post('/createreport',(req,res)=>{
    contract("INVOKE",["CreateNewReport",req.body.patient_id,req.body.ref_doctor],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json(payload.toString())
        }
    })
})

routes.put('/starttreatment',(req,res)=>{
    contract("INVOKE",["StartTreatment",req.headers.treatment_id,req.body.supervisor],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json({
                "message": `successfully started the treatment`
            })
        }
    })
})
module.exports= routes
