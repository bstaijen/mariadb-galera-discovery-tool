# Discovery tool for Discovering MariaDB Server
This is a small golang application build for querying the [consul](https://github.com/gliderlabs/docker-consul) catalog API and returning a comma-seperated-list containing service addresses.

## Parameters
- `-address` - The address were consul is running. Eq: `consul:8500`
- `-service` - Name of the service you are searching. Eq: `galera-db`
- `-debug` - Adds debug info to the output.
- `-version` - Enables you to only show the version of the tool.

## Returns
A comma-separated list containing the service addresses
Example: `10.0.0.1,10.0.0.2,10.0.0.3`

## How to use it
Example 1
`mariadb-disover-tool -address=consul:8500 -service=galera-db`

Example 2
`mariadb-disover-tool -address=192.168.99.100:8500 -service=galera-db -debug`

Example 3
`mariadb-disover-tool -version`

## To Do
- Add [godep](https://github.com/tools/godep) for dependency management
- Add extra discovery backends.