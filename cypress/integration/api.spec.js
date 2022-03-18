/// <reference types="cypress" />

describe('/api/firmware', () => {
  context('GET', () => {
    beforeEach(() => {
      cy.uploadFirmware('power-meter', '1.0.0', 100)
    })

    it('returns the list of firmware', () => {
      cy.request('http://localhost:5000/api/firmware').then((response) => {
        expect(response.status).to.eq(status.OK)
        expect(response.headers['content-type']).to.include('application/json')
      })
    })
  })

  context('PUT', () => {
    beforeEach(() => {
      cy.request('DELETE', 'http://localhost:5000/api/firmware/temperature-sensor/1.0.0')
    })

    it('uploads a new firmware', () => {
      let previousLength = -1
      cy.request('http://localhost:5000/api/firmware').then((response) => {
        expect(response.status).to.eq(status.OK)
        expect(response.headers['content-type']).to.include('application/json')
        expect(response.body).to.not.include({
          type: 'temperature-sensor',
          version: '1.0.0',
          size: 100
        })
        previousLength = response.body.length
      })

      cy.uploadFirmware('temperature-sensor', '1.0.0', 100).then((response) => {
        expect(response.status).to.eq(status.CREATED)
      })

      cy.request('http://localhost:5000/api/firmware').then((response) => {
        expect(response.status).to.eq(status.OK)
        expect(response.headers['content-type']).to.include('application/json')
        expect(response.body).to.have.length(previousLength+1)
        expect(response.body).to.deep.include.members([{
          type: 'temperature-sensor',
          version: '1.0.0',
          size: 100
        }])
      })
    })
  })

  context('DELETE', () => {
    beforeEach(() => {
      cy.request('PUT', 'http://localhost:5000/api/firmware/water-meter/2.0.0')
    })

    it('removes a firmware', () => {
      let previousLength = -1
      cy.request('http://localhost:5000/api/firmware').then((response) => {
        expect(response.status).to.eq(status.OK)
        expect(response.headers['content-type']).to.include('application/json')
        expect(response.body).to.deep.include.members([{
          type: 'water-meter',
          version: '2.0.0',
          size: 100
        }])
        previousLength = response.body.length
      })

      cy.request('DELETE', 'http://localhost:5000/api/firmware/water-meter/2.0.0').then((response) => {
        expect(response.status).to.eq(status.OK)
      })

      cy.request('http://localhost:5000/api/firmware').then((response) => {
        expect(response.status).to.eq(status.OK)
        expect(response.headers['content-type']).to.include('application/json')
        expect(response.body).to.have.length(previousLength-1)
        expect(response.body).to.not.deep.include.members([{
          type: 'water-meter',
          version: '2.0.0',
          size: 100
        }])
      })
    })

    it('returns ok if deleting something that does not exist', () => {
      cy.request('DELETE', 'http://localhost:5000/api/firmware/does-not-exist/9.9.9').then((response) => {
        expect(response.status).to.eq(status.OK)
      })
    })
  })

  describe('/firmware/types', () => {
    beforeEach(() => {
      cy.request('PUT', 'http://localhost:5000/api/firmware/power-meter/0.0.1')
      cy.request('PUT', 'http://localhost:5000/api/firmware/water-meter/0.0.1')
      cy.request('PUT', 'http://localhost:5000/api/firmware/temperature-sensor/0.0.1')
    })

    context('GET', () => {
      it('returns the list of types', () => {
        cy.request('http://localhost:5000/api/firmware/types').then((response) => {
          expect(response.status).to.eq(status.OK)
          expect(response.headers['content-type']).to.include('application/json')
          expect(response.body).to.include.members([
            'power-meter',
            'temperature-sensor',
            'water-meter'
          ])
        })
      })
    })
  })
})

describe('/api/devices', () => {
  context('GET', () => {
    it('returns the list of devices', () => {
      cy.request('http://localhost:5000/api/devices').then((response) => {
        expect(response.status).to.eq(200)
        expect(response.headers['content-type']).to.include('application/json')
      })
    })
  })
})