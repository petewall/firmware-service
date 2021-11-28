class Firmware {
  constructor(type, version) {
    this.type = type
    this.version = version
  }

  json() {
    return {
      type: this.type,
      version: this.version
    }
  }
}

export default Firmware
