class FirmwareStore {
  constructor(db) {
    this.db = db
  }

  async getAll() {
    return await this.db.find()
      .project({'data': 0})
      .sort({'type': 1, 'version': 1})
      .toArray()
  }

  async getAllByType(type) {
    return await this.db.find({type})
      .project({'data': 0})
      .sort({'type': 1, 'version': 1})
      .toArray()
  }

  async getAllTypes() {
    const types = await this.db.find()
      .project({'_id': 0, 'type': 1})
      .sort({'type': 1})
      .toArray()
    const uniqueTypes = new Set(types.map(type => type.type))
    return Array.from(uniqueTypes)
  }

  async get(type, version) {
    return await this.db.findOne({type, version})
  }

  async add(type, version, data) {
    const size = data.length
    const result = await this.db.insertOne({type, version, size, data})
    return result.acknowledged
  }

  async delete(type, version) {
    const result = await this.db.deleteOne({type, version})
    return result.deletedCount === 1
  }
}

module.exports = FirmwareStore
