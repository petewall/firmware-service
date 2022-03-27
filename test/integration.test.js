const { MongoClient } = require('mongodb')
const status = require('http-status')
const supertest = require('supertest')

const config = {
  db: {
    host: process.env.FIRMWARE_DB_HOST || 'localhost',
    port: process.env.FIRMWARE_DB_PORT || 27017,
    name: process.env.FIRMWARE_DB_NAME || 'firmware-service-test',
    username: process.env.FIRMWARE_DB_USERNAME || 'mongoadmin',
    password: process.env.FIRMWARE_DB_PASSWORD || 'secret'
  }
}

async function dbClean() {
  const connectionURI = `mongodb://${config.db.username}:${config.db.password}@${config.db.host}:${config.db.port}`
  const mongodb = new MongoClient(connectionURI)
  await mongodb.connect()

  const db = mongodb.db(config.db.name)
  await db.dropDatabase()

  await mongodb.close()
}

describe('Integration tests', () => {
  let app, request
  beforeAll(async () => {
    await dbClean()
    app = await require('../server.js')(config)
    request = supertest(app)
  })

  afterAll(async () => {
    // await dbClean()
    app.shutdown()
  })

  describe('Empty database', () => {
    describe('GET /', () => {
      it('returns an empty list', async () => {
        const response = await request.get('/')

        expect(response.status).toEqual(status.OK)
        expect(response.header['content-type']).toEqual('application/json; charset=utf-8')
        expect(response.body).toEqual([])
      })
    })
    describe('GET /types', () => {
      it('returns an empty list', async () => {
        const response = await request.get('/types')

        expect(response.status).toEqual(status.OK)
        expect(response.header['content-type']).toEqual('application/json; charset=utf-8')
        expect(response.body).toEqual([])
      })
    })
    describe('GET /testfirmware', () => {
      it('returns not found', async () => {
        const response = await request.get('/testfirmware')

        expect(response.status).toEqual(status.NOT_FOUND)
        expect(response.error.text).toEqual('no firmware found for type "testfirmware"')
      })
    })
    describe('GET /testfirmware/1.2.3', () => {
      it('returns not found', async () => {
        const response = await request.get('/testfirmware/1.2.3')

        expect(response.status).toEqual(status.NOT_FOUND)
        expect(response.error.text).toEqual('no firmware found for type "testfirmware" with version "1.2.3"')
      })
    })
    describe('GET /testfirmware/1.2.3/data', () => {
      it('returns not found', async () => {
        const response = await request.get('/testfirmware/1.2.3/data')

        expect(response.status).toEqual(status.NOT_FOUND)
        expect(response.error.text).toEqual('no firmware found for type "testfirmware" with version "1.2.3"')
      })
    })
    describe('DELETE /testfirmware/1.2.3', () => {
      it('returns not found', async () => {
        const response = await request.delete('/testfirmware/1.2.3')

        expect(response.status).toEqual(status.NOT_FOUND)
        expect(response.error.text).toEqual('no firmware found for type "testfirmware" with version "1.2.3"')
      })
    })
  })

  describe('Uploading firmware', () => {
    describe('PUT /<types>/<versions>', () => {
      it('uploads the firmware', async () => {
        let putResponse = await request.put('/testfirmware/1.2.3')
          .set('Content-type', 'application/octet-stream')
          .send('this is the firmware content for version 1.2.3')

        expect(putResponse.status).toEqual(status.CREATED)

        putResponse = await request.put('/testfirmware/2.3.4')
          .set('Content-type', 'application/octet-stream')
          .send('this is the firmware content for version 2.3.4')

        expect(putResponse.status).toEqual(status.CREATED)

        putResponse = await request.put('/anotherfirmware/1.0.0')
          .set('Content-type', 'application/octet-stream')
          .send('this is the firmware content for another firmware')

        expect(putResponse.status).toEqual(status.CREATED)
      })
    })

    // describe('PUT /<types>/<versions> (already exists)', () => {
    //   it('')
    // })
  })

  describe('Populated database', () => {
    describe('GET /', () => {
      it('returns the list of firmware', async () => {
        const response = await request.get('/')

        expect(response.status).toEqual(status.OK)
        expect(response.header['content-type']).toEqual('application/json; charset=utf-8')
        expect(response.body.length).toEqual(3)
      })
    })

    describe('GET /types', () => {
      it('returns the sorted list of types', async () => {
        const response = await request.get('/types')

        expect(response.status).toEqual(status.OK)
        expect(response.header['content-type']).toEqual('application/json; charset=utf-8')
        expect(response.body).toEqual(['anotherfirmware', 'testfirmware'])
      })
    })

    describe('GET /testfirmware', () => {
      it('returns the list of firmware for a given type', async () => {
        const response = await request.get('/testfirmware')

        expect(response.status).toEqual(status.OK)
        expect(response.body.length).toEqual(2)
      })
    })

    describe('GET /testfirmware/1.2.3', () => {
      it('returns the detail and data for a single firmware', async () => {
        const response = await request.get('/testfirmware/1.2.3')

        expect(response.status).toEqual(status.OK)
        expect(response.body.type).toEqual('testfirmware')
        expect(response.body.version).toEqual('1.2.3')
        expect(response.body.size).toEqual('this is the firmware content for version 1.2.3'.length)
        expect(response.body.data).toEqual(
          Buffer.from('this is the firmware content for version 1.2.3', 'utf8').toString('base64')
        )
      })
    })

    describe('GET /testfirmware/1.2.3/data', () => {
      it('downloads the firmware data', async () => {
        const response = await request.get('/testfirmware/1.2.3/data')

        expect(response.headers['content-type']).toEqual('application/octet-stream; charset=utf-8')
        expect(response.status).toEqual(status.OK)
        expect(response.body).toEqual(
          Buffer.from('this is the firmware content for version 1.2.3')
        )
      })
    })
  })

  describe('Deleting firmware', () => {
    describe('DELETE /testfirmware/1.2.3', () => {
      it('removes the firmware', async () => {
        let getResponse = await request.get('/testfirmware')
        expect(getResponse.status).toEqual(status.OK)
        expect(getResponse.body.length).toEqual(2)

        const deleteResponse = await request.delete('/testfirmware/1.2.3')
        expect(deleteResponse.status).toEqual(status.OK)

        getResponse = await request.get('/testfirmware')
        expect(getResponse.status).toEqual(status.OK)
        expect(getResponse.body.length).toEqual(1)
      })
    })
  })
})
