/// <reference types="cypress" />

describe('firmware UI', () => {
  beforeEach(() => {
    cy.request('PUT', 'http://localhost:5000/firmware/firmware-a/1.0.0')
    cy.request('PUT', 'http://localhost:5000/firmware/firmware-c/0.0.1')
    cy.request('PUT', 'http://localhost:5000/firmware/firmware-a/0.0.1')
    cy.request('PUT', 'http://localhost:5000/firmware/firmware-b/0.20.1')
    cy.request('PUT', 'http://localhost:5000/firmware/firmware-b/0.2.1')
  })

  context('clicking the firmware tab', () => {
    it('displays the list of firmware', () => {
      cy.visit('http://localhost:5000')
      cy.get('a').contains('Firmware').click()
      cy.get('th').contains('Type')
      cy.get('th').contains('Version')
      cy.get('th').contains('Size')
      cy.get('th').contains('Actions')
    })
  })

  context('visiting the firmware link', () => {
    it('displays the list of firmware', () => {
      cy.visit('http://localhost:5000/#/firmware')
      cy.get('th').contains('Type')
      cy.get('th').contains('Version')
      cy.get('th').contains('Size')
      cy.get('th').contains('Actions')
    })
  })

  it('sorts the devices by type and version', () => {
    cy.intercept('http://localhost:5000/firmware').as('firmwareRequest')
    cy.visit('http://localhost:5000/#/firmware')
    cy.wait('@firmwareRequest')

    let previousType, previousVersion
    cy.get('#firmware-table tbody tr').each((row, index) => {
      const type = row.children('td:nth-child(0)').text()
      const version = row.children('td:nth-child(1)').text()
      if (index > 0) {
        if (type == previousType) {
          expect(version.localeCompare(previousVersion)).to.be.gte(0)
        } else {
          expect(type.localeCompare(previousType)).to.be.lte(0)
        }
      }
      previousType = type
      previousVersion = version
    })
  })

  // context('when a new firmware binary is uploaded', () => {
  //   beforeEach(() => {
  //     cy.request('DELETE', 'http://localhost:5000/firmware/firmware-d/1.2.3')
  //   })

  //   it('displays that new firmware', () => {
  //     cy.intercept('http://localhost:5000/firmware').as('firmwareRequest')
  //     cy.visit('http://localhost:5000/#/firmware')
  //     cy.wait('@firmwareRequest')

  //     cy.get('#firmware-table tbody tr').then((startingRows) => {
  //       cy.request('PUT', 'http://localhost:5000/firmware/firmware-d/1.2.3')
  //       cy.wait('@firmwareRequest')
  
  //       const updatedRows = cy.get('#firmware-table tbody tr')
  //       expect(updatedRows).to.have.length(startingRows.length+1)
  //     })
  //   })
  // })
})
