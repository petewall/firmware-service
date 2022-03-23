const Firmware = require('../lib/firmware.js')
const bodyParser = require('body-parser')
const router = require('express').Router()
const status = require('http-status')

module.exports = (firmwareLibrary, socketServer) => {
  router.get('/', async (req, res) => {
    const firmwareList = await firmwareLibrary.getAll()
    res.json(firmwareList)
  })

  router.get('/types', async (req, res) => {
    res.json(await firmwareLibrary.getAllTypes())
  })

  router.get('/:type/', async (req, res) => {
    const firmware = await firmwareLibrary.getAllByType(req.params.type)
    if (firmware.length == 0) {
      return res.status(status.NOT_FOUND).send(`no firmware found for type "${req.params.type}"`)
    }
    res.json(firmware)
  })

  router.get('/:type/:version([0-9a-zA-Z-._]+)', async (req, res) => {
    const firmware = await firmwareLibrary.get(req.params.type, req.params.version)
    if (!firmware) {
      return res.status(status.NOT_FOUND).send(`no firmware found for type "${req.params.type}" with version "${req.params.version}"`)
    }
    res.json(firmware)
  })

  router.get('/:type/:version([0-9a-zA-Z-._]+)/data', async (req, res) => {
    const firmware = await firmwareLibrary.get(req.params.type, req.params.version)
    if (!firmware) {
      return res.status(status.NOT_FOUND).send(`no firmware found for type "${req.params.type}" with version "${req.params.version}"`)
    }
    res.setHeader('Content-Type', 'application/octet-stream')
    res.send(firmware.data)
  })

  router.put('/:type/:version([0-9a-zA-Z-._]+)', bodyParser.raw({ limit: '5mb' }), async (req, res) => {
    const created = await firmwareLibrary.add(req.params.type, req.params.version, req.body)
    if (created) {
      res.sendStatus(status.CREATED)
    } else {
      res.sendStatus(status.OK)
    }
    socketServer.sockets.emit('firmware', {
      action: 'CREATED',
      type: req.params.type,
      version: req.params.version
    })
  })

  router.delete('/:type/:version([0-9a-zA-Z-._]+)', async (req, res) => {
    const deleted = await firmwareLibrary.delete(req.params.type, req.params.version)
    if (deleted) {
      res.sendStatus(status.OK)
      socketServer.sockets.emit('firmware', {
        action: 'DELETED',
        type: req.params.type,
        version: req.params.version
      })
    } else {
      res.sendStatus(status.NOT_FOUND)
    }
  })

  return router
}
