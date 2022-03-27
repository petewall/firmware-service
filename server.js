const express = require('express')
const app = express()
const debug = require('debug')
const log = debug('firmware-service')
debug.enable('firmware-service')
const morgan = require('morgan')
const { MongoClient } = require('mongodb')

app.use(morgan('dev', { stream: { write: msg => log(msg) } }))

async function connectToFirmwareDB(config) {
  const connectionURI = `mongodb://${config.db.username}:${config.db.password}@${config.db.host}:${config.db.port}`
  const connectionURIRedacted = `mongodb://${config.db.username}:<redacted>@${config.db.host}:${config.db.port}`
  log(`Connecting to MongoDB: ${connectionURIRedacted}...`)
  const mongodb = new MongoClient(connectionURI)
  await mongodb.connect()

  return mongodb
}

async function initializeCollection(mongodb, config) {
  const collectionName = 'firmware'
  log(`Connecting to db: ${config.db.name}...`)
  const db = mongodb.db(config.db.name)

  if (!db.listCollections({name: collectionName}).hasNext()) {
    log(`Creating collection: ${collectionName}`)
    await db.createCollection(collectionName)
  }
  const collection = db.collection('firmware')

  // log('checking for index')
  // if (!await collection.indexExists('type_1_version_1')) {
  //   log('Creating index...')
  //   await collection.createIndex(['type', 'version'])
  // }

  log('Database is ready')
  return collection
}

module.exports = async (config) => {
  const db = await connectToFirmwareDB(config)
  const collection = await initializeCollection(db, config)
  const firmwareAPI = require('./api.js')(collection)
  app.use(firmwareAPI)
  app.shutdown = async () => {
    await db.close()
  }
  return app
}
