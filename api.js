const bodyParser = require('body-parser')
const { Int32, Binary } = require('bson')
const router = require('express').Router()
const status = require('http-status')

module.exports = (db) => {
  router.get('/', async (req, res) => {
    const projection = { data: 0 }
    const sort = { type: 1, version: 1 }
    const firmwareList = await db.find({}, { projection, sort }).toArray()
    res.json(firmwareList)
  })

  router.get('/types', async (req, res) => {
    const projection = { _id: 0, type: 1 }
    const sort = { type: 1 }
    const types = await db.find({}, { projection, sort }).toArray()

    const uniqueTypes = new Set(types.map(type => type.type))
    res.json(Array.from(uniqueTypes))
  })

  router.get('/:type/', async (req, res) => {
    const type = req.params.type

    const projection = { data: 0 }
    const sort = { type: 1, version: 1 }
    const firmware = await db.find({ type }, { projection, sort }).toArray()
    if (firmware.length == 0) {
      return res.status(status.NOT_FOUND).send(`no firmware found for type "${req.params.type}"`)
    }
    res.json(firmware)
  })

  router.get('/:type/:version([0-9a-zA-Z-._]+)', async (req, res) => {
    const type = req.params.type
    const version = req.params.version
    const firmware = await db.findOne({ type, version })
    if (!firmware) {
      return res.status(status.NOT_FOUND).send(`no firmware found for type "${req.params.type}" with version "${req.params.version}"`)
    }
    res.json(firmware)
  })

  router.get('/:type/:version([0-9a-zA-Z-._]+)/data', async (req, res) => {
    const type = req.params.type
    const version = req.params.version
    const firmware = await db.findOne({ type, version })
    if (!firmware) {
      return res.status(status.NOT_FOUND).send(`no firmware found for type "${req.params.type}" with version "${req.params.version}"`)
    }
    res.setHeader('Content-Disposition', `attachment; filename="${type}-${version}.bin"`)
    res.setHeader('Content-Type', 'application/octet-stream')

    res.send(firmware.data.value())
  })

  router.put('/:type/:version([0-9a-zA-Z-._]+)', bodyParser.raw({ limit: '5mb' }), async (req, res) => {
    const type = req.params.type
    const version = req.params.version
    const data = new Binary(req.body)
    const size = new Int32(req.body.length)

    const result = await db.insertOne({ type, version, size, data })
    if (result.acknowledged) {
      res.sendStatus(status.CREATED)
    } else {
      // TODO: Determine what the error was, is this a 500 or a 400?
      res.sendStatus(status.IM_A_TEAPOT)
    }
  })

  router.delete('/:type/:version([0-9a-zA-Z-._]+)', async (req, res) => {
    const type = req.params.type
    const version = req.params.version
    const result = await db.deleteOne({ type, version })
    if (result.deletedCount === 1) {
      res.sendStatus(status.OK)
    } else {
      res.status(status.NOT_FOUND).send(`no firmware found for type "${req.params.type}" with version "${req.params.version}"`)
    }
  })

  return router
}
