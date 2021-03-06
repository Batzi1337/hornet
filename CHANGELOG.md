# HORNET Changelog

All notable changes to this project will be documented in this file.

## [0.2.12] - 04.01.2020

### Fixed

    - Fixes broken ARM7 build with enabled CGO

## [0.2.11] - 04.01.2020

### Added

    - Seperate config file for neighbor settings
    - MQTT broker plugin
    - IOTA Tangle Visualiser plugin
    - Print HORNET version at startup
    - getLedgerDiffExt webapi call for debug purposes

### Removed

    - Almost all command line flags were removed (use the config file instead)
    - Removed "default" profile (use "auto" instead)

### Changed

    - Switched to hive.go packages to reduce codebase
    - Several speed improvements (binary/trinary conversion) due to latest iota.go version

### Fixed

    - Fixes possible panic with reattached milestones
    - Issue were milestoneSolidifierWorkerPool could block processing of tx
    - Fixes concurrent writes to the host blacklist
    - Fixes wrong order of bundles checks in solidifier

### Config file changes

New options:

`config.json`
```json
  "graph": {
    "webrootPath": "IOTAtangle/webroot",
    "socketiopath": "socket.io-client/dist/socket.io.js",
    "domain": "",
    "host": "127.0.0.1",
    "port": 8083,
    "networkName": "meets HORNET"
  },
  "mqtt": {
    "config": "mqtt_config.json"
  },
```

Now there is a seperate file for the neighbor settings:

`neighbors.json`
```json
{
  "autotetheringenabled": false,
  "maxneighbors": 5,
  "neighbors": [
    {
      "identity": "example1.neighbor.com:15600",
      "alias": "Example Neighbor 1",
      "preferIPv6": false
    },
    {
      "identity": "example2.neighbor.com:15600",
      "alias": "Example Neighbor 2",
      "preferIPv6": false
    },
    {
      "identity": "example3.neighbor.com:15600",
      "alias": "Example Neighbor 3",
      "preferIPv6": false
    }
  ]
}
```

Removed options:

`config.json`
```diff
  "network": {
    "address": "0.0.0.0",
-    "autotetheringenabled": false,
    "preferIPv6": false,
-    "maxneighbors": 5,
-    "neighbors": [
-      {
-        "identity": "example1.neighbor.com:15600",
-        "preferIPv6": false
-      },
-      {
-        "identity": "example2.neighbor.com:15600",
-        "preferIPv6": false
-      },
-      {
-        "identity": "example3.neighbor.com:15600",
-        "preferIPv6": false
-      }
-    ],
    "port": 15600,
    "reconnectattemptintervalseconds": 60
  },
```

## [0.2.10] - 27.12.2019

### Added

    - arm64 and armhv support to the Dockerfile

## [0.2.9] - 20.12.2019

### Fixed

    - `addNeighbors` deadlock
    - Message logger caused fatal panic

## [0.2.8] - 19.12.2019

### Added

    - Rate limiting for WebSocket sends
    - Show address balance even if no txs are available (Dashboard - Explorer)
    - Show spent state (Dashboard - Explorer)
    - Port configuration for Monitor plugin
    - Config to prefer IPv6 (addNeighbors)
    - Alternative addNeighbors command

### Changed

    - Release archives now contain a dir which wraps all files
    - API errors
    - TPS chart for better visibility of input and output (Dashboard)

### Fixed

    - Check wasSyncedBefore in ZMQ and Monitor
    - Wrong ZeroMQ `tx_trytes` response order
    - Deadlock if node is shut down during startup phase
    - Different TX order than IRI (attachToTangle)
    - Log level was ignored

### Config file changes

New options:

```json

  "network": {
    "preferIPv6": false,
  }

  "monitor": {
    "domain": "",
    "host": "127.0.0.1",
    "port": 4434,
    "apiPort": 4433
  }
```

**Changed option (you have to edit it in your config):**

```json
  "node": {
    "loglevel": 127
  }
```

## [0.2.7] - 17.12.2019

### Added

    - Version printout `--version`

### Changed

    - WorkerPools don't get flushed at shutdown by default
    - Import spent addresses in smaller batches
    - Faster syncing

### Fixed

    - RequestQueue never got empty if the cache overflowed
    - Several shutdown problems
    - Issue were only tail tx of a bundle got confirmed
    - Status report was still active during shutdown
    - Future cone solidifier got stuck, causing the node to become unsync

## [0.2.6] - 16.12.2019

### Changed

    - Faster initial spent addresses import

## [0.2.5] - 15.12.2019

### Added

    - More badger options in the profiles
    - "auto" profile chooses best setting based on available system memory

### Changed

    - "compactLevel0OnClose" is now disabled per default
    - Faster shutdown of the node

### Config file changes

New option:

```json
  "useProfile": "auto",
```

## [0.2.4] - 15.12.2019

This release fixes a CRITICAL bug! You have to delete your database folder.

### Fixed

    - Spent addresses were not imported from snapshot file.

## [0.2.3] - 15.12.2019

### Fixed

    - Close on closed channel in "ordered daemon" on shutdown

## [0.2.2] - 15.12.2019

### Added

    - TangleMonitor Plugin
    - Spammer Plugin
    - More detailed log messages at shutdown

### Fixed

    - Do not expose passwords from config file at startup
    - Duplicated neighbors

### Config file changes

New settings:

```json
  "monitor": {
    "tanglemonitorpath": "tanglemonitor/frontend",
    "domain": "",
    "host": "127.0.0.1"
  },
  "spammer": {
    "address": "HORNET99INTEGRATED99SPAMMER999999999999999999999999999999999999999999999999999999",
    "depth": 3,
    "message": "Spamming with HORNET tipselect",
    "tag": "HORNET99INTEGRATED99SPAMMER",
    "tpsratelimit": 0.1,
    "workers": 1
  },
  "zmq": {
    "host": "127.0.0.1",
  }
```

## [0.2.1] - 13.12.2019

### Added

    - Cache Metrics in SPA
    - Profiles to adjust cache sizes and DB opts

### Fixed

    - Remote PoW for Trinity

## [0.2.0] - 12.12.2019

### Added

    - DB version number
    - Configurable zmq host
    - Solidification timestamp of transactions
    - Docker files

### Changed

    - Database layout (breaking change)

### Fixed

    - Trinity compatibility
    - WebAPI CORS headers

## [0.1.0] - 11.12.2019

### Added

    - First beta release
