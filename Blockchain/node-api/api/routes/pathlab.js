const express = require('express')
const {contract} = require('../contract')

routes = express.Router()

routes.put('/dotest',(req,res)=>{
    contract("INVOKE",["DoTest",req.headers.test_id,req.body.test_result,req.body.supervisor,req.body.no_of_mediafile],(err,payload)=>{
        if (err){
            res.status(500).json(err)
        }
        else{
            res.status(200).json(JSON.parse(payload))
        }
    })
})

module.exports = routes