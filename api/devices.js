const Router = require('express').Router
const httpStatus = require('http-status')
const router = Router()


module.exports = (devices, firmwareList, socketServer) => {
  router.get('/', async (req, res) => {
    const deviceList = await devices.getAll()
    res.json(deviceList.map((device) => device.json()))
  })

  router.get('/:mac', async (req, res) => {
    res.status(httpStatus.IM_A_TEAPOT).send('not yet implemented')
  })

  router.delete('/:mac', async (req, res) => {
    res.status(httpStatus.IM_A_TEAPOT).send('not yet implemented')
  })

  router.post('/:mac', async (req, res) => {
    res.status(httpStatus.IM_A_TEAPOT).send('not yet implemented')
  })

  return router
}
