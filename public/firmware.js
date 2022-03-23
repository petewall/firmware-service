function deleteFirmware(row, type, version) {
  const yesDelete = confirm(`Are you sure you want to delete the firmware ${type} ${version}?\nThis action cannot be undone!`)
  if (yesDelete) {
    $.ajax(`/api/firmware/${type}/${version}`, {
      method: 'DELETE',
    }).done(() => {
      row.remove()
    })
  }
}

function byTypeAndVersion(a, b) {
  if (a.type == b.type) {
    return -a.version.localeCompare(b.version)
  }
  return a.type.localeCompare(b.type)
}

$(() => {
  $.get('/api/firmware', (firmwareLibrary) => {
    firmwareLibrary.sort(byTypeAndVersion)

    $('#firmware-table tbody').empty()
    firmwareLibrary.map((firmware) => {
      const row = $('<tr>')
      const deleteButton = $('<button class="ui icon button">').append(
        $('<i class="trash icon">')
      ).on('click', () => { deleteFirmware(row, firmware.type, firmware.version) })

      $('#firmware-table tbody').append(
        row.append(
          $('<td>', {text: firmware.type}),
          $('<td>', {text: firmware.version}),
          $('<td>', {text: firmware.size}),
          $('<td>').append(deleteButton)
        )
      )
    })
  })
})
