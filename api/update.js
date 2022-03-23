const httpStatus = require('http-status')
const firmware = require('./firmware')

const Router = require('express').Router
const router = Router()

module.exports = (devices, firmwareLibrary, socketServer) => {
  router.get('/update', async (req, res) => {
    let mac = req.get('x-esp8266-sta-mac')
    const currentType = req.query.firmware
    const currentVersion = req.query.version

    const device = await devices.update(mac, firmwareLibrary.get(currentType, currentVersion))
    const newFirmware = await firmwareLibrary.getUpdatedFirmware(device)
    if (!newFirmware) {
      return res.sendStatus(httpStatus.NOT_MODIFIED)
    }

    return res.send(newFirmware.data)
  })

  return router
}
