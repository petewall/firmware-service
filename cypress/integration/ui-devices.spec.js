/// <reference types="cypress" />

describe('the Devices UI', () => {
  // beforeEach(() => {
  //   // Cypress starts out with a blank slate for each test
  //   // so we must tell it to visit our website with the `cy.visit()` command.
  //   // Since we want to visit the same URL at the start of all our tests,
  //   // we include it in our beforeEach function so that it runs before each test
  //   cy.visit('https://localhost:5000/api/devices')
  // })
  it('displays the list of devices', () => {
    cy.visit('http://localhost:5000')
    cy.get('th').contains('MAC')
    cy.get('th').contains('Last Update Request')
    cy.get('th').contains('Current Firmware')
    cy.get('th').contains('Assigned Firmware')
  })

  context('clicking the devices tab', () => {
    it('displays the list of devices', () => {
      cy.visit('http://localhost:5000/#/firmware')
      cy.get('a').contains('Devices').click()
      cy.get('th').contains('MAC')
      cy.get('th').contains('Last Update Request')
      cy.get('th').contains('Current Firmware')
      cy.get('th').contains('Assigned Firmware')
    })
  })

  context('visiting the devices link', () => {
    it('displays the list of devices', () => {
      cy.visit('http://localhost:5000/#/devices')
      cy.get('th').contains('MAC')
      cy.get('th').contains('Last Update Request')
      cy.get('th').contains('Current Firmware')
      cy.get('th').contains('Assigned Firmware')
    })
  })

  // it('sorts the devices by MAC address', () => {})

  // context('when a new device requests for an update', () => {
  //   it('displays that new device', () => {})
  // })
})
