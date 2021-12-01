class Firmware {
  constructor(type, version, size) {
    this.type = type
    this.version = version
    this.size = size
  }

  json() {
    return {
      type: this.type,
      version: this.version,
      size: this.size
    }
  }
}

module.exports = Firmware
