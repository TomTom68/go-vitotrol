Das Paket `go-vitotrol` ermöglicht einen Zugriff auf die Viessmann™
Vitotrol™ cloud API und damit zur Anzeige und Steuerung einer Heizungsanlage.

Siehe auch https://www.viessmann.com/app_vitodata/VIIWebService-1.16.0.0/iPhoneWebService.asmx

Folgende Befehle sind umgesetzt und funktionieren für eine Heizungsanlage vom Typ Vitodens 333-F:
- Login
- GetDevices
- RequestRefreshStatus
- RequestWriteStatus
- GetData
- WriteData
- RefreshData
- GetErrorHistory
- GetTimesheet
- WriteTimesheetData
- GetTypeInfo

## Installation

### The `vitotrol` command

To obtain a `vitotrol` executable in `$GOPATH/bin/` directory if
`GOPATH` environment variable exists or in `$HOME/go/bin` otherwise:

#### starting go 1.18

```
go install github.com/maxatome/go-vitotrol/cmd/vitotrol@master
```

#### before go 1.18

```
go get -u github.com/maxatome/go-vitotrol/cmd/vitotrol
```

### The library
```
go get -u github.com/maxatome/go-vitotrol
```

`-u` became useless since go 1.18.

## Example

See `cmd/vitotrol/*.go` for an example of use.

Executable `vitotrol` usage follows:

```
usage: vitotrol [OPTIONS] ACTION [PARAMS]
  -config string
        login+password config file
  -debug
        print debug information
  -device string
        DeviceID, index, DeviceName, DeviceId@LocationID, DeviceName@LocationName (see `devices' action) (default "0")
  -json
        used by `timesheet' action to display timesheets using JSON format
  -login string
        login on vitotrol API
  -password string
        password on vitotrol API
  -verbose
        print verbose information

ACTION & PARAMS can be:
- devices              list all available devices
- list [attrs|timesheets]  list attribute (default) or timesheet names
- get ATTR_NAME ...    get the value of attributes ATTR_NAME, ... on vitodata
                         server
- get all              get all known attributes on vitodata server
- rget ATTR_NAME ...   refresh then get the value of attributes ATTR_NAME, ...
                         on vitodata server
- rget all             refresh than get all known attributes on vitodata server
- bget ATTR_IDX ...    get the value of attributes ATTR_IDX, ... on vitodata
                         server without checking their validity before (for
                         developing purpose)
- rbget ATTR_IDX ...   refresh then get the value of attributes ATTR_IDX, ...
                         on vitodata server without checking their validity
                         before (for developing purpose)
- set ATTR_NAME VALUE  set the value of attribute ATTR_NAME to VALUE
- timesheet TIMESHEET ...
                       get the timesheet TIMESHEET data
- set_timesheet TIMESHEET '{"wday":[{"from":630,"to":2200},...],...}'
                       replace the whole timesheet TIMESHEET
                       wday is either a day (eg. mon) or a range of days
                       (eg. mon-wed or sat-mon)
                       The JSON content can be in a file with the syntax @file
- errors               get the error history
- remote_attrs         list server available attributes
                         (for developing purpose)
```

The config file is a two lines file containing the LOGIN on the first
line and the PASSWORD on the second, and is named
`$HOME/.vitotrol-api` by default (when all `--config`, `--login` and
`--password` options are missing or empty):

```
LOGIN
PASSWORD
```

## License

go-vitotrol is released under the MIT License.
