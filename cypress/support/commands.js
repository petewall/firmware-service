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

Cypress.Commands.add('updateRequest', (mac, type, version) => {
  return cy.request({
    method: 'GET',
    url: `http://localhost:5000/api/update?firmware=${type}&version=${version}`,
    headers: {
      'x-esp8266-sta-mac': mac
    }
  })
})
