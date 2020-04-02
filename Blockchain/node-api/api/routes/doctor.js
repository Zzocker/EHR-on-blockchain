const express = require('express')
const {contract} = require('../contract')
routes = express.Router()

routes.put('/report/addcomment',(req,res)=>{
    contract("INVOKE",["AddCommentsToReport",req.headers.report_id,req.body.comment,req.body.ref_doctor],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json({
                "message": `Successfuly added comment to report ${req.headers.report_id}`
            })
        }
    })
})

routes.post('/report/reftest',(req,res)=>{
    contract("INVOKE",["RefTest",req.headers.report_id,req.body.name,req.body.ref_doctor,req.body.type_of_test],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json({
                "message": `Successfuly created test of type ${req.body.type_of_test}`,
                'test_id': payload.toString()
            })
        }
    })
})
routes.post('/report/reftreatment',(req,res)=>{
    contract("INVOKE",["RefTreatment",req.headers.report_id,req.body.ref_doctor,req.body.name],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json({
                "message": `Successfuly created treatment of name ${req.body.name}`,
                'test_id': payload.toString()
            })
        }
    })
})

routes.post('/report/presdrugs',(req,res)=>{
    contract("INVOKE",["PrescribeDrugs",req.headers.report_id,req.body.ref_doctor,req.body.drugs,req.body.doses],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json({
                "message": `Successfuly prescribe drugs`,
                'drug_id': payload.toString()
            })
        }
    })
})


routes.put('/treatment/addcomment',(req,res)=>{
    contract("INVOKE",["AddCommentsToTreatment",req.headers.treatment_id,req.body.supervisor,req.body.comment],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json({
                "message": `Successfuly commented on treatment`
            })
        }
    })
})

routes.put('/treatment/addmediafile',(req,res)=>{
    contract("INVOKE",["AddMediaToTreatment",req.headers.treatment_id,req.body.supervisor,req.body.no_of_mediafile],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json(JSON.parse(payload))
        }
    })
})

module.exports = routes