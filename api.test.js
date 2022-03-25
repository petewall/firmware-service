const API = require('./api.js')
const express = require('express')
const status = require('http-status')
const supertest = require('supertest')

function fakeFirmware(type, version, size) {
  return {
    type, version, size, data: `${size} bits of binary data`
  }
}

describe('API', () => {
  let app, firmwareStore
  beforeEach(() => {
    firmwareStore = {}
    app = express()
    app.use(API(firmwareStore))
  })

  describe('GET /', () => {
    beforeEach(() => {
      firmwareStore.getAll = jest.fn(() => Promise.resolve([
        fakeFirmware('bootloader', '1.2.3', 100)
      ]))
    })

    it('returns all firmware', async () => {
      const response = await supertest(app).get('/')

      expect(response.status).toEqual(status.OK)
      expect(response.body).toEqual([{
        type: 'bootloader',
        version: '1.2.3',
        size: 100,
        data: '100 bits of binary data'
      }])
    })
  })
})