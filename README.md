# Firmware Service

This is a replacement for the OTA service.

The architecture:

1. Dependency on mongodb for firmware status
2. Dependency on minio for firmware storage?

Tested in cypress
UI designed with fomantic



UI design:
List on the left with binaries that have registrered with the service
List on the right, displayed in collapsable categories by firmware type, expanded to show versions.
