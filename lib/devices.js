const Device = require('./device.js')
const Firmware = require('./firmware.js')

class Devices {
  constructor() {
    this.devices = {}
  }

  async getAll() {
    return Object.values(this.devices)
  }

  async update(mac, firmware) {
    let device = this.devices[mac]
    if (!device) {
      device = new Device(mac, firmware, new Firmware(firmware.type, null))
    }

    this.devices[mac] = device
  }
}

module.exports = Devices
