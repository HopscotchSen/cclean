# cclean
A simple service clean tool for Consul.

## How to use?

``` sh
➜  ~ cclean --help

A simple service clean tool for Consul.

Usage:
  cclean [CONSUL_ADDRESS] [flags]

Flags:
      --exclude string     exclude consul node ip (eg: 10.20.0.0/16)
  -h, --help               help for cclean
      --port               consul server port (default 8500)
      --timeout duration   http timeout (default 3s)
```

## Example

``` sh
➜  ~ cclean 127.0.0.1:8500
or
➜  ~ cclean 127.0.0.1:8501 --port 8501
```

If you do not provide consul address default use environment variables.

## Environment

Supported environment variables:

``` go
// HTTPAddrEnvName defines an environment variable name which sets
// the HTTP address if there is no -http-addr specified.
HTTPAddrEnvName = "CONSUL_HTTP_ADDR"

// HTTPTokenEnvName defines an environment variable name which sets
// the HTTP token.
HTTPTokenEnvName = "CONSUL_HTTP_TOKEN"

// HTTPAuthEnvName defines an environment variable name which sets
// the HTTP authentication header.
HTTPAuthEnvName = "CONSUL_HTTP_AUTH"

// HTTPSSLEnvName defines an environment variable name which sets
// whether or not to use HTTPS.
HTTPSSLEnvName = "CONSUL_HTTP_SSL"

// HTTPCAFile defines an environment variable name which sets the
// CA file to use for talking to Consul over TLS.
HTTPCAFile = "CONSUL_CACERT"

// HTTPCAPath defines an environment variable name which sets the
// path to a directory of CA certs to use for talking to Consul over TLS.
HTTPCAPath = "CONSUL_CAPATH"

// HTTPClientCert defines an environment variable name which sets the
// client cert file to use for talking to Consul over TLS.
HTTPClientCert = "CONSUL_CLIENT_CERT"

// HTTPClientKey defines an environment variable name which sets the
// client key file to use for talking to Consul over TLS.
HTTPClientKey = "CONSUL_CLIENT_KEY"

// HTTPTLSServerName defines an environment variable name which sets the
// server name to use as the SNI host when connecting via TLS
HTTPTLSServerName = "CONSUL_TLS_SERVER_NAME"

// HTTPSSLVerifyEnvName defines an environment variable name which sets
// whether or not to disable certificate checking.
HTTPSSLVerifyEnvName = "CONSUL_HTTP_SSL_VERIFY"

// GRPCAddrEnvName defines an environment variable name which sets the gRPC
// address for consul connect envoy. Note this isn't actually used by the api
// client in this package but is defined here for consistency with all the
// other ENV names we use.
GRPCAddrEnvName = "CONSUL_GRPC_ADDR"
```

See also: https://github.com/hashicorp/consul/blob/0e280b5a08bbdcf9a214ba2063b1eb850fba1824/api/api.go#L24

## Use with docker

Run cclean in a docker:

``` sh
$ docker run -e CONSUL_HTTP_ADDR=127.0.0.1 gozap/cclean --help
```
