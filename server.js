#!/usr/local/bin/node

const express = require('express')
const app = express()
const bodyParser = require('body-parser')
const server = require('http').createServer(app)
const FirmwareStore = require('./firmware-store.js')
const status = require('http-status')

const port = process.env.PORT || 5000

app.use(require('morgan')('combined'))

async function connectToFirmwareDB() {
  const firmwareDBHost = process.env.FIRMWARE_DB_HOST || 'localhost'
  const firmwareDBPort = process.env.FIRMWARE_DB_PORT || 27017
  const firmwareDBUsername = process.env.FIRMWARE_DB_USERNAME || 'mongoadmin'
  const firmwareDBPassword = process.env.FIRMWARE_DB_PASSWORD || 'secret'

  const { MongoClient } = require('mongodb')
  const mongodb = new MongoClient(`mongodb://${firmwareDBUsername}:${firmwareDBPassword}@${firmwareDBHost}:${firmwareDBPort}`)
  await mongodb.connect()
  return mongodb.db('firmware-service')
}

(async function main() {
  const db = await connectToFirmwareDB()
  const firmwareStore = new FirmwareStore(db.collection('firmware'))
  const firmwareAPI = require('./api.js')(firmwareStore)
  app.use(firmwareAPI)
  server.listen(port)
})()
