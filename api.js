const bodyParser = require('body-parser')
const router = require('express').Router()
const status = require('http-status')

module.exports = (firmwareStore) => {
  // const firmwareStore = new FirmwareStore(db.collection('firmware'))
  router.get('/', async (req, res) => {
    const firmwareList = await firmwareStore.getAll()
    res.json(firmwareList)
  })

  router.get('/types', async (req, res) => {
    res.json(await firmwareStore.getAllTypes())
  })

  router.get('/:type/', async (req, res) => {
    const firmware = await firmwareStore.getAllByType(req.params.type)
    if (firmware.length == 0) {
      return res.status(status.NOT_FOUND).send(`no firmware found for type "${req.params.type}"`)
    }
    res.json(firmware)
  })

  router.get('/:type/:version([0-9a-zA-Z-._]+)', async (req, res) => {
    const firmware = await firmwareStore.get(req.params.type, req.params.version)
    if (!firmware) {
      return res.status(status.NOT_FOUND).send(`no firmware found for type "${req.params.type}" with version "${req.params.version}"`)
    }
    res.json(firmware)
  })

  router.get('/:type/:version([0-9a-zA-Z-._]+)/data', async (req, res) => {
    const firmware = await firmwareStore.get(req.params.type, req.params.version)
    if (!firmware) {
      return res.status(status.NOT_FOUND).send(`no firmware found for type "${req.params.type}" with version "${req.params.version}"`)
    }
    res.setHeader('Content-Type', 'application/octet-stream')
    res.send(firmware.data)
  })

  router.put('/:type/:version([0-9a-zA-Z-._]+)', bodyParser.raw({ limit: '5mb' }), async (req, res) => {
    const created = await firmwareStore.add(req.params.type, req.params.version, req.body)
    if (created) {
      res.sendStatus(status.CREATED)
    } else {
      res.sendStatus(status.OK)
    }
  })

  router.delete('/:type/:version([0-9a-zA-Z-._]+)', async (req, res) => {
    const deleted = await firmwareStore.delete(req.params.type, req.params.version)
    if (deleted) {
      res.sendStatus(status.OK)
    } else {
      res.sendStatus(status.NOT_FOUND)
    }
  })

  return router
}
