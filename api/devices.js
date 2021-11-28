import Devices from '../lib/devices.js'
import { Router } from 'express'
const router = Router()

const devices = new Devices()

router.get('/', async (req, res) => {
  const deviceList = await devices.getAll()
  res.json(deviceList.map((device) => device.json()))
})

export default router
