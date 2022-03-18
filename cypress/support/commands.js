// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
Cypress.Commands.add('uploadFirmware', (type, version, size) => {
  const blocks = Array.from(
    {length: size},
    () => Uint8Array.from(
      {length: 1024 * 1024},
      () => Math.floor(Math.random() * 256)
    )
  )

  return cy.request({
    method: 'PUT',
    url: `http://localhost:5000/api/firmware/${type}/${version}`,
    body: blocks,
    headers: {
      'content-type': 'application/octet-stream'
    }
  })
})
//
//
// -- This is a child command --
// Cypress.Commands.add('drag', { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add('dismiss', { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite('visit', (originalFn, url, options) => { ... })
