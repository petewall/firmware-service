# Firmware Service

This service is for managing a store of firmware binaries.

## API

### `GET /`

### `GET /types`

### `GET /<type>`

### `GET /<type>/<version>`

### `GET /<type>/<version>/data`

### `PUT /<type>/<version>`

### `DELETE /<type>/<version>`

## TODO

* Test in-memory firmware store
* Implement file system firmware store
* Implement cli switches to choose firmware store
* Implement FS-based integration tests
* Get testing in CI
