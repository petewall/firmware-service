const Devices = require('../lib/devices.js')
const Router = require('express').Router
const router = Router()

const devices = new Devices()

router.get('/', async (req, res) => {
  const deviceList = await devices.getAll()
  res.json(deviceList.map((device) => device.json()))
})

module.exports = router
