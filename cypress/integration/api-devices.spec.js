/// <reference types="cypress" />

describe('/devices', () => {
  context('GET', () => {
    it('returns the list of devices', () => {
      cy.request('http://localhost:5000/devices').then((response) => {
        expect(response.status).to.eq(200)
        expect(response.headers['content-type']).to.include('application/json')
      })
    })
  })
})
