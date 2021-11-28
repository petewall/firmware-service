import Device from './device.js'
import Firmware from './firmware.js'

class Devices {
  constructor() {
    this.devices = {
      'b8:e8:56:44:fd:20': new Device('b8:e8:56:44:fd:20', new Firmware('lightswitches', '1.2.3'), null)
    }
  }

  async getAll() {
    return Object.values(this.devices)
  }
}

export default Devices
