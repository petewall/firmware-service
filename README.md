# Firmware Service

Todo:

* Store firmware binaries on the FS (for now)

* Manage firmware entries with redis
  type, version, size, location
* Manage devices with redis
  devices
    mac, currentType, currentVersion, assignedType, assignedVersion (pinning)
  device updates
  mac, time, result

* Add firmware-store to store firmware binaries
