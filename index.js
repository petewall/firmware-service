#!/usr/local/bin/node

const express = require('express')
const app = express()
const port = process.env.PORT || 5000
const server = require('http').createServer(app)
const socketServer = require('socket.io')(server)

socketServer.on('connection', () => {
  console.log('New user connected')
})

app.use(require('morgan')('combined'))
app.use('/devices', require('./api/devices.js'))
app.use('/firmware', require('./api/firmware.js')(socketServer))

app.use(express.static('public'))
app.use('/lib/jquery', express.static('node_modules/jquery/dist'))
app.use('/lib/jquery-address', express.static('node_modules/jquery-address/src'))
app.use('/lib/socket.io', express.static('node_modules/socket.io/client-dist'))

server.listen(port)
