#!/usr/local/bin/node

const express = require('express')
const app = express()
const deviceRouter = require('./api/devices.js')
const firmwareRouter = require('./api/firmware.js')
const morgan = require('morgan')
const port = process.env.PORT || 5000

app.use(morgan('combined'))
app.use('/devices', deviceRouter)
app.use('/firmware', firmwareRouter)
app.use(express.static('public'))
app.listen(port)
