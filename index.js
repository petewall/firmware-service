#!/usr/local/bin/node

const express = require('express')
const app = express()
const port = process.env.PORT || 5000
const server = require('http').createServer(app)
const socketServer = require('socket.io')(server)

async function connectToFirmwareDB() {
  const firmwareDBHost = process.env.FIRMWARE_DB_HOST || 'localhost'
  const firmwareDBPort = process.env.FIRMWARE_DB_PORT || 27017
  const firmwareDBUsername = process.env.FIRMWARE_DB_USERNAME || 'mongoadmin'
  const firmwareDBPassword = process.env.FIRMWARE_DB_PASSWORD || 'secret'

  const { MongoClient } = require('mongodb')
  const mongodb = new MongoClient(`mongodb://${firmwareDBUsername}:${firmwareDBPassword}@${firmwareDBHost}:${firmwareDBPort}`)
  await mongodb.connect()
  return mongodb
}

(async function main() {
  const mongodb = await connectToFirmwareDB()
  const firmwareDB = mongodb.db('firmware-service')

  const Devices = require('./lib/devices.js')
  const devices = new Devices()

  const FirmwareLibrary = require('./lib/firmware-library.js')
  const firmwareLibrary = new FirmwareLibrary(firmwareDB.collection('firmware'))

  app.use(require('morgan')('combined'))
  app.use('/api/devices', require('./api/devices.js')(devices, firmwareLibrary, socketServer))
  app.use('/api/firmware', require('./api/firmware.js')(firmwareLibrary, socketServer))
  app.use('/api/update', require('./api/update.js')(devices, firmwareLibrary, socketServer))

  app.use(express.static('public'))
  app.use('/lib/jquery', express.static('node_modules/jquery/dist'))
  app.use('/lib/jquery-address', express.static('node_modules/jquery-address/src'))
  app.use('/lib/socket.io', express.static('node_modules/socket.io/client-dist'))

  socketServer.on('connection', () => {
    console.log('New user connected')
  })

  server.listen(port)
})()
