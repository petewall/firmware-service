class Device {
  constructor(mac, currentFirmware, assignedFirmware) {
    this.mac = mac
    this.currentFirmware = currentFirmware
    this.assignedFirmware = assignedFirmware
  }

  json() {
    return {
      mac: this.mac,
      currentFirmware: this.currentFirmware?.json(),
      assignedFirmware: this.assignedFirmware?.json()
    }
  }
}

export default Device
