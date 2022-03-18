class Device {
  constructor(mac, currentFirmware, assignedFirmware) {
    this.mac = mac
    this.currentFirmware = currentFirmware
    this.assignedFirmware = assignedFirmware
    this.pinnedVersion
  }

  json() {
    return {
      mac: this.mac,
      currentFirmware: this.currentFirmware?.json(),
      assignedFirmware: this.assignedFirmware?.json()
    }
  }
}

module.exports = Device
