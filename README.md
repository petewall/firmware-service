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

* Protect against protected words
  * No type should be named "types", so you cannot `PUT /types/1.2.3`
* Determine what to do when uploading firmware with the same type and version
  * Error? firmware is immutable
  * Update? firmware is mutable and updatable
  * Create new record? firmware is immutable, but the newest copy is used
* Add APIs for
  * Get latest firmware for a given type `GET /<type>/latest`. `latest` becomes a protected version
* Better error handling for rejected promises
* Enable integration testing in CI
