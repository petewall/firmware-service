const API = require('../api.js')
const { Int32, Binary } = require('bson')
const express = require('express')
const status = require('http-status')
const supertest = require('supertest')

function fakeFirmware(type, version, size) {
  return {
    type,
    version,
    size: new Int32(size),
    data: new Binary(`${size} bits of binary data`)
  }
}

describe('Unit tests', () => {
  let app
  let findToArrayReturn = []
  let findOneReturn
  let insertOneReturn
  let deleteOneReturn
  const db = {
    find: jest.fn(() => {
      return {
        toArray: jest.fn(() => Promise.resolve(findToArrayReturn))
      }
    }),
    findOne: jest.fn(() => Promise.resolve(findOneReturn)),
    insertOne: jest.fn(() => Promise.resolve(insertOneReturn)),
    deleteOne: jest.fn(() => Promise.resolve(deleteOneReturn))
  }

  beforeEach(() => {
    app = express()
    app.use(API(db))
    db.find.mockClear()
    db.findOne.mockClear()
    db.insertOne.mockClear()
    db.deleteOne.mockClear()
  })

  describe('GET /', () => {
    beforeEach(() => {
      findToArrayReturn = [
        fakeFirmware('bootloader', '1.2.3', 100),
        fakeFirmware('lightswitch', '1.2.3', 100)
      ]
    })

    it('returns all firmware', async () => {
      const response = await supertest(app).get('/')

      expect(response.status).toEqual(status.OK)
      expect(response.body).toEqual([{
        type: 'bootloader',
        version: '1.2.3',
        size: 100,
        data: Buffer.from('100 bits of binary data', 'utf8').toString('base64')
      }, {
        type: 'lightswitch',
        version: '1.2.3',
        size: 100,
        data: Buffer.from('100 bits of binary data', 'utf8').toString('base64')
      }])

      expect(db.find.mock.calls.length).toEqual(1)
      expect(db.find.mock.calls[0][0]).toEqual({})
      expect(db.find.mock.calls[0][1]).toEqual({
        projection: {data: 0},
        sort: {type: 1, version: 1}
      })
    })
  })

  describe('GET /types', () => {
    beforeEach(() => {
      findToArrayReturn = [
        {type: 'bootloader'},
        {type: 'bootloader'},
        {type: 'lightswitch'}
      ]
    })

    it('returns a sorted list of types', async () => {
      const response = await supertest(app).get('/types')

      expect(response.status).toEqual(status.OK)
      expect(response.body).toEqual(['bootloader', 'lightswitch'])

      expect(db.find.mock.calls.length).toEqual(1)
      expect(db.find.mock.calls[0][0]).toEqual({})
      expect(db.find.mock.calls[0][1]).toEqual({
        projection: {_id: 0, type: 1},
        sort: {type: 1}
      })
    })
  })

  describe('GET /<type>', () => {
    beforeEach(() => {
      findToArrayReturn = [
        fakeFirmware('bootloader', '1.0.0', 100),
        fakeFirmware('bootloader', '1.2.3', 100)
      ]
    })

    it('returns all firmware for a given type', async () => {
      const response = await supertest(app).get('/bootloader')

      expect(response.status).toEqual(status.OK)
      expect(response.body).toEqual([{
        type: 'bootloader',
        version: '1.0.0',
        size: 100,
        data: Buffer.from('100 bits of binary data', 'utf8').toString('base64')
      }, {
        type: 'bootloader',
        version: '1.2.3',
        size: 100,
        data: Buffer.from('100 bits of binary data', 'utf8').toString('base64')
      }])

      expect(db.find.mock.calls.length).toEqual(1)
      expect(db.find.mock.calls[0][0]).toEqual({
        type: 'bootloader'
      })
      expect(db.find.mock.calls[0][1]).toEqual({
        projection: {data: 0},
        sort: {type: 1, version: 1}
      })
    })
  })

  describe('GET /<type> (not found)', () => {
    beforeEach(() => {
      findToArrayReturn = []
    })

    it('returns 404', async () => {
      const response = await supertest(app).get('/nothing')

      expect(response.status).toEqual(status.NOT_FOUND)
      expect(response.error.text).toEqual('no firmware found for type "nothing"')

      expect(db.find.mock.calls.length).toEqual(1)
      expect(db.find.mock.calls[0][0]).toEqual({
        type: 'nothing'
      })
      expect(db.find.mock.calls[0][1]).toEqual({
        projection: {data: 0},
        sort: {type: 1, version: 1}
      })
    })
  })

  describe('GET /<type>/<version>', () => {
    beforeEach(() => {
      findOneReturn = fakeFirmware('bootloader', '1.2.3', 100)
    })

    it('returns a specific firmware', async () => {
      const response = await supertest(app).get('/bootloader/1.2.3')

      expect(response.status).toEqual(status.OK)
      expect(response.body).toEqual({
        type: 'bootloader',
        version: '1.2.3',
        size: 100,
        data: Buffer.from('100 bits of binary data', 'utf8').toString('base64')
      })

      expect(db.findOne.mock.calls.length).toEqual(1)
      expect(db.findOne.mock.calls[0][0]).toEqual({
        type: 'bootloader',
        version: '1.2.3'
      })
    })
  })

  describe('GET /<type>/<version> (not found)', () => {
    beforeEach(() => {
      findOneReturn = null
    })

    it('returns 404', async () => {
      const response = await supertest(app).get('/nothing/9.9.9-rc1')

      expect(response.status).toEqual(status.NOT_FOUND)
      expect(response.error.text).toEqual('no firmware found for type "nothing" with version "9.9.9-rc1"')

      expect(db.findOne.mock.calls.length).toEqual(1)
      expect(db.findOne.mock.calls[0][0]).toEqual({
        type: 'nothing',
        version: '9.9.9-rc1'
      })
    })
  })

  describe('GET /<type>/<version>/data', () => {
    beforeEach(() => {
      findOneReturn = fakeFirmware('bootloader', '1.2.3', 100)
    })

    it('downloads the firmware data', async () => {
      const response = await supertest(app).get('/bootloader/1.2.3/data')

      expect(response.status).toEqual(status.OK)
      expect(response.headers['content-disposition']).toEqual('attachment; filename="bootloader-1.2.3.bin"')
      expect(response.headers['content-type']).toEqual('application/octet-stream; charset=utf-8')
      expect(response.body).toEqual(
        Buffer.from('100 bits of binary data')
      )

      expect(db.findOne.mock.calls.length).toEqual(1)
      expect(db.findOne.mock.calls[0][0]).toEqual({
        type: 'bootloader',
        version: '1.2.3'
      })
    })
  })

  describe('GET /<type>/<version>/data (not found)', () => {
    beforeEach(() => {
      findOneReturn = null
    })

    it('downloads the firmware data', async () => {
      const response = await supertest(app).get('/nothing/0.0.1-beta1/data')

      expect(response.status).toEqual(status.NOT_FOUND)
      expect(response.error.text).toEqual('no firmware found for type "nothing" with version "0.0.1-beta1"')

      expect(db.findOne.mock.calls.length).toEqual(1)
      expect(db.findOne.mock.calls[0][0]).toEqual({
        type: 'nothing',
        version: '0.0.1-beta1'
      })
    })
  })

  describe('PUT /<type>/<version>', () => {
    beforeEach(() => {
      insertOneReturn = {
        acknowledged: true
      }
    })

    it('uploads a new firmware', async () => {
      const response = await supertest(app).put('/bootloader/2.0.0')
        .set('Content-type', 'application/octet-stream')
        .send('this is the new firmware content')

      expect(response.status).toEqual(status.CREATED)

      expect(db.insertOne.mock.calls.length).toEqual(1)
      expect(db.insertOne.mock.calls[0][0]).toEqual({
        type: 'bootloader',
        version: '2.0.0',
        size: new Int32('this is the new firmware content'.length),
        data: new Binary('this is the new firmware content')
      })
    })
  })

  describe('PUT /<type>/<version> (fails)', () => {
    beforeEach(() => {
      insertOneReturn = {
        acknowledged: false
      }
    })

    it('responds with an error', async () => {
      const response = await supertest(app).put('/bootloader/2.0.0')
        .set('Content-type', 'application/octet-stream')
        .send('this is the new firmware content')

      expect(response.status).toEqual(status.IM_A_TEAPOT)

      expect(db.insertOne.mock.calls.length).toEqual(1)
      expect(db.insertOne.mock.calls[0][0]).toEqual({
        type: 'bootloader',
        version: '2.0.0',
        size: new Int32('this is the new firmware content'.length),
        data: new Binary('this is the new firmware content')
      })
    })
  })

  describe('DELETE /<type>/<version>', () => {
    beforeEach(() => {
      deleteOneReturn = {
        deletedCount: 1
      }
    })

    it('removes a firmware', async () => {
      const response = await supertest(app).delete('/bootloader/2.0.0')

      expect(response.status).toEqual(status.OK)

      expect(db.deleteOne.mock.calls.length).toEqual(1)
      expect(db.deleteOne.mock.calls[0][0]).toEqual({
        type: 'bootloader',
        version: '2.0.0'
      })
    })
  })

  describe('DELETE /<type>/<version> (not found)', () => {
    beforeEach(() => {
      deleteOneReturn = {
        deletedCount: 0
      }
    })

    it('removes a firmware', async () => {
      const response = await supertest(app).delete('/bootloader/2.0.0')

      expect(response.status).toEqual(status.NOT_FOUND)
      expect(response.error.text).toEqual('no firmware found for type "bootloader" with version "2.0.0"')

      expect(db.deleteOne.mock.calls.length).toEqual(1)
      expect(db.deleteOne.mock.calls[0][0]).toEqual({
        type: 'bootloader',
        version: '2.0.0'
      })
    })
  })
})
