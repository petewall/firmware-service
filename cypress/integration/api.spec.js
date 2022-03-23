/// <reference types="cypress" />

const httpStatus = require('http-status')

describe('/api/firmware', () => {
  context('GET', () => {
    beforeEach(() => {
      cy.uploadFirmware('power-meter', '1.0.0', 100)
    })

    it('returns the list of firmware', () => {
      cy.request('http://localhost:5000/api/firmware').then((response) => {
        expect(response.status).to.eq(httpStatus.OK)
        expect(response.headers['content-type']).to.include('application/json')
        expect(response.body).to.deep.include.members([{
          type: 'power-meter',
          version: '1.0.0',
          size: 100
        }])
      })
    })
  })
})

describe('/api/firmware/<type>/<version>', () => {
  context('PUT', () => {
    beforeEach(() => {
      cy.request('DELETE', 'http://localhost:5000/api/firmware/temperature-sensor/1.0.0')
    })

    it('uploads a new firmware', () => {
      let previousLength = -1
      cy.request('http://localhost:5000/api/firmware').then((response) => {
        expect(response.status).to.eq(httpStatus.OK)
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
        expect(response.status).to.eq(httpStatus.OK)
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
      cy.uploadFirmware('water-meter', '1.0.0', 100)
    })

    it('removes a firmware', () => {
      let previousLength = -1
      cy.request('http://localhost:5000/api/firmware').then((response) => {
        expect(response.status).to.eq(httpStatus.OK)
        expect(response.headers['content-type']).to.include('application/json')
        expect(response.body).to.deep.include.members([{
          type: 'water-meter',
          version: '2.0.0',
          size: 100
        }])
        previousLength = response.body.length
      })

      cy.request('DELETE', 'http://localhost:5000/api/firmware/water-meter/2.0.0').then((response) => {
        expect(response.status).to.eq(httpStatus.OK)
      })

      cy.request('http://localhost:5000/api/firmware').then((response) => {
        expect(response.status).to.eq(httpStatus.OK)
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
        expect(response.status).to.eq(httpStatus.OK)
      })
    })
  })  
})

describe('/firmware/types', () => {
  beforeEach(() => {
    cy.uploadFirmware('power-meter', '0.0.1', 100)
    cy.uploadFirmware('water-meter', '0.0.1', 100)
    cy.uploadFirmware('temperature-sensor', '0.0.1', 100)
  })

  context('GET', () => {
    it('returns the list of types', () => {
      cy.request('http://localhost:5000/api/firmware/types').then((response) => {
        expect(response.status).to.eq(httpStatus.OK)
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

describe('/api/devices', () => {
  context('GET', () => {
    beforeEach(() => {
      cy.updateRequest('b8:e8:56:44:fd:20', 'lightswitch', '1.0.0')
    })

    it('returns the list of devices', () => {
      cy.request('http://localhost:5000/api/devices').then((response) => {
        expect(response.status).to.eq(200)
        expect(response.headers['content-type']).to.include('application/json')
        
        expect(response.body).to.deep.include.members([{
          mac: 'b8:e8:56:44:fd:20',
          currentFirmware: {
            type: 'lightswitch',
            version: '1.0.0',
          },
          assignedFirmware: null
        }])
      })
    })
  })
})

describe('/api/devices/<mac>', () => {
  context('GET', () => {
    it('details about a device', () => {
      // cy.request('http://localhost:5000/api/devices').then((response) => {
      //   expect(response.status).to.eq(200)
      //   expect(response.headers['content-type']).to.include('application/json')
      // })
    })
  })

  context('DELETE', () => {
    it('forgets a device', () => {
      
    })
  })

  context('POST', () => {
    context('sending firmware type only', () => {
      it('assigns a firmware type to a device', () => {
        
      })
    })

    context('sending firmware type and version', () => {
      it('assigns a firmware type and version to a device', () => {
        
      })
    })
  })
})

describe('/api/update', () => {
  context('GET', () => {
    context('device has the same firmware and version', () => {
      it('does not send an update', () => {})
    })
    context('device has the same firmware but an older version', () => {
      it('sends the update', () => {})
    })
    context('device has the same firmware but a newer version', () => {
      it('does not send an update', () => {})
    })
    context('device has a different firmware', () => {
      it('sends the update', () => {})
    })
    context('new device, known firmware', () => {
      it('adds the device to the list', () => {})
    })
    context('new device, unknown firmware', () => {
      it('adds the device to the list, does not assign a firmware', () => {})
    })
  })
})
