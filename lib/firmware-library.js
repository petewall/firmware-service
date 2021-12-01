// const Firmware = require('./firmware.js')

class FirmwareLibrary {
  constructor() {
    this.allFirmware = []
  }

  async getAll() {
    return this.allFirmware
  }

  async getAllTypes() {
    let types = {}
    for (const firmware of this.allFirmware) {
      console.log(`type: ${firmware.type}`)
      types[firmware.type] = true
    }
    return Object.keys(types).sort()
  }

  async add(firmware) {
    if (!this.has(firmware.type, firmware.version)) {
      this.allFirmware.push(firmware)
      return true
    }
    return false
  }

  async delete(type, version) {
    this.allFirmware = this.allFirmware.filter((firmware) => {
      return (firmware.type != type || firmware.version != version)
    })
  }

  has(type, version) {
    for (const firmware of this.allFirmware) {
      if (firmware.type == type && firmware.version == version) {
        return true
      }
    }
    return false
  }
}

module.exports = FirmwareLibrary
