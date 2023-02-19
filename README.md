# Firmware Service

This service is for managing a store of firmware binaries.

## API

* `GET /` - Get all firmware
* `GET /types` - Get all firmware types
* `GET /<type>` - Get all firmware by type
* `GET /<type>/<version>` - Get firmware by type and version
* `GET /<type>/<version>/data` - Download firmware data
* `PUT /<type>/<version>` - Upload new firmware
* `DELETE /<type>/<version>` - Delete firmware

## TODO

* Holistic testing in filesystem tests?
* Get testing in CI
