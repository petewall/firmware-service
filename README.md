# Firmware Service

## Todo

* Get things working with local memory
* Get things working with MongoDB

## Tests

### API test

#### Firmware

* GET /api/firmware returns the list of firmware
  * Before: put at least one firmware
* GET /api/firmware/types returns the list of firmware types
  * Before: put at least two firmware types
* PUT /api/firmware/`type`/`version` adds the new firmware
  * Before: ensure the type is missing
  * test: put returns OK, ensure the type exists
* DELETE /api/firmware/`type`/`version` removes the firmware

#### Device

* GET /api/devices returns the list of devices
  * before: send update to register one device
* GET /api/devices/`mac` gets detail about a device, including history
* DELETE /api/devices/`mac` forgets a device

* POST /api/devices/`mac` with new firmware type
* POST /api/devices/`mac` with new firmware type and version

* GET /api/update sends updated firmware to a device
  * Device has the same firmware, and same version - No update
  * Device has the same firmware, but older version - Send update
  * Device has the same firmware, but newer version - No update
  * Device has a different firmware - send update
  * New device, has known firmware - set assignment to that firmware with latest
  * New device, unknown firmware - do not assign firmware

### UI tests

Device list should have an assigned firmware type, defaulting to bootloader, and pinned version list, defaulting to "latest".