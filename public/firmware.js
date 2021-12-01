$(() => {
  $.get('/firmware', (firmwareLibrary) => {
    $('#firmware-table tbody').empty()
    firmwareLibrary.map((firmware) => {
      $('#firmware-table tbody').append(
        $('<tr>').append(
          $('<th>', {text: firmware.type}),
          $('<td>', {text: firmware.version}),
          $('<td>', {text: firmware.size}),
          // $('<td>').append($('<input>', {type: 'checkbox'}))
        )
      )
    })

    // $('#firmware-table').tablesort() // TODO: Need to import the plugin https://fomantic-ui.com/collections/table.html#sortable
  })
})
