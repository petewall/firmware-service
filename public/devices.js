$(() => {
  $.get('/devices', (devices) => {
    $('#device-table').empty()
    devices.map((device) => {
      const currentFirmware = `${device.currentFirmware.type} ${device.currentFirmware.version}`
      const assignedFirmwareType = device?.assignedFirmware?.type || ''

      $('#device-table').append(
        $('<tr>').append(
          $('<th>', {text: device.mac}),
          $('<td>'),
          $('<td>', currentFirmware),
          $('<td>', assignedFirmwareType),
        )
      )
    })
  })
})
