# das

[![License](http://img.shields.io/:license-apache%202.0-brightgreen.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)

## What is das?

**das** is an open source project that provides a reliable autonomous platform for the usage and management of MySQL database. This autonomous platform provides real-time monitoring of the MySQL database, regular inspections, and healthcheck alarms. The platform will centrally manage the MySQL database to assist users in using the MySQL database in a more standardized and efficient manner.

## Deployment

The project can be deployed according to the following steps:

1. Use Percona Monitoring and Management to centrally manage MySQL that needs to be monitored (omitted)
2. Install a local MySQL database to store das data, and the database table structure is in the sql directory of das
3. Install sql tuning tool [soar](https://github.com/XiaoMi/soar) into the bin directory
4. Create configuration files das.yaml and soar.yaml in the config directory, there are template yaml files in the config directory
5. das uses soar as Linux shell command, so on Windows, we need to create a shell environment by installing cygwin or mingw, normally, installing git tool is enough, make sure that the paths of soar.exe and sh.exe are in the PATH environment

## Quick Start

The following software is required to work with the das codebase and build it locally:

-  Go version `Go >= 1.16`

To check the source code and build binaries, you can simply run:

Build
```
go build -o das main.go
```

Start
```
./das start --daemon=true --config=./config/das.yaml
```

Swagger
```
http://127.0.0.1:6090/swagger/index.html
```

## Contributing

<a href="https://github.com/romberli/das/graphs/contributors">
  <img src="https://contributors-img.web.app/image?repo=romberli/das"/>
</a>

## License

Copyright das Authors.
Licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

