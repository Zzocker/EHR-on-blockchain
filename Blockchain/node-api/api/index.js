const express = require('express')
const bodyParser = require('body-parser')

const patient = require('./routes/patient')
const hospital = require('./routes/hospital')
const doctor = require('./routes/doctor')
const pathlab = require('./routes/pathlab')
const pharmacies = require('./routes/pharmacies')
const general = require('./routes/general')

PORT = process.env.PORT || 3000


const app = express()
const logger = (req,res,next)=>{
    console.log(`${req.protocol}://${req.get('host')}${req.originalUrl}`)
    next()
}
// app.use(express.json())
app.use(bodyParser.urlencoded({ extended: false }))
app.use(bodyParser.json())
app.use(logger)

app.use('/patient',patient)
app.use('/hospital',hospital)
app.use('/doctor',doctor)
app.use('/pathlab',pathlab)
app.use('/pharmacies',pharmacies)
app.use('/',general)

app.listen(PORT,()=>{
    console.log(`listening on port: ${PORT}`)
})