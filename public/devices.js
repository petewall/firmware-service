$(() => {
  $.get('/devices', (devices) => {
    $('#device-table tbody').empty()
    devices.map((device) => {
      const currentFirmware = `${device.currentFirmware.type} ${device.currentFirmware.version}`
      const assignedFirmwareType = device?.assignedFirmware?.type || ''

      $('#device-table tbody').append(
        $('<tr>').append(
          $('<th>', {text: device.mac}),
          $('<td>'),
          $('<td>', {text: currentFirmware}),
          $('<td>', {text: assignedFirmwareType}),
        )
      )
    })

    // $('#device-table').tablesort() // TODO: Need to import the plugin https://fomantic-ui.com/collections/table.html#sortable
  })
})
