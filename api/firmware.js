const Firmware = require('../lib/firmware.js')
const FirmwareLibrary = require('../lib/firmware-library.js')
const Router = require('express').Router
const router = Router()
const status  = require('http-status')

const firmwareLibrary = new FirmwareLibrary()

router.get('/', async (req, res) => {
  const firmwareList = await firmwareLibrary.getAll()
  res.json(firmwareList.map((firmware) => firmware.json()))
})

router.get('/types', async (req, res) => {
  res.json(await firmwareLibrary.getAllTypes())
})

router.put('/:type/:version([0-9a-zA-Z-._]+)', async (req, res) => {
  const firmware = new Firmware(req.params.type, req.params.version, 100)
  const created = await firmwareLibrary.add(firmware)
  if (created) {
    res.sendStatus(status.CREATED)
  } else {
    res.sendStatus(status.OK)
  }
})

router.delete('/:type/:version([0-9a-zA-Z-._]+)', async (req, res) => {
  await firmwareLibrary.delete(req.params.type, req.params.version)
  res.sendStatus(status.OK)
})

module.exports = router
