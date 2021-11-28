#!/usr/local/bin/node

import express from 'express'
const app = express()
import deviceRouter from './api/devices.js'
import morgan from 'morgan'
const port = process.env.PORT || 5000

// app.get("/healthcheck", (req, res) => {
//   res.sendStatus(status.OK)
// })

app.use(morgan('combined'))
app.use('/devices', deviceRouter)
app.use(express.static('public'))
app.listen(port)
