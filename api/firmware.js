const Firmware = require('../lib/firmware.js')
const FirmwareLibrary = require('../lib/firmware-library.js')
const bodyParser = require('body-parser')
const router = require('express').Router()
const status  = require('http-status')

const firmwareLibrary = new FirmwareLibrary()
module.exports = (socketServer) => {
  router.get('/', async (req, res) => {
    const firmwareList = await firmwareLibrary.getAll()
    res.json(firmwareList.map((firmware) => firmware.json()))
  })

  router.get('/types', async (req, res) => {
    res.json(await firmwareLibrary.getAllTypes())
  })

  router.put('/:type/:version([0-9a-zA-Z-._]+)', bodyParser.raw({ limit: '5mb' }), async (req, res) => {
    const firmware = new Firmware(req.params.type, req.params.version, req.body.length)
    const created = await firmwareLibrary.add(firmware, req.body)
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
    await firmwareLibrary.delete(req.params.type, req.params.version)
    res.sendStatus(status.OK)
    socketServer.sockets.emit('firmware', {
      action: 'DELETED',
      type: req.params.type,
      version: req.params.version
    })
  })

  return router
}
