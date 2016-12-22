# Discovery tool for Discovering MariaDB Server
This is a small golang application build for querying a [consul](https://github.com/gliderlabs/docker-consul) instance. This application is not production ready.

## Parameters
- `-h` - Location of the consul instance with port. Eq: `consul:8500`
- `-servicename` - Name of the service you're seraching. Eq: `galera-db`

## Returns
A comma-separated list of IPs. 
Example: `10.0.0.1,10.0.0.2,10.0.0.3`

## How to use it
`mariadb-disover-tool -h=consul:8500 -servicename=galera-db`