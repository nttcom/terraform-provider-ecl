# eclcloud: an Enterprise Cloud SDK for Go

eclcloud is an Enterprise Cloud Go SDK.

## Useful links

* [Effective Go](https://golang.org/doc/effective_go.html)

## How to install

Before installing, you need to ensure that your [GOPATH environment variable](https://golang.org/doc/code.html#GOPATH)
is pointing to an appropriate directory where you want to install eclcloud:

```bash
mkdir $HOME/go
export GOPATH=$HOME/go
```

## Getting started

### Credentials

Because you'll be hitting an API, you will need to retrieve your Enterprise Cloud credentials and either store them as environment variables or in your local Go
files. The first method is recommended because it decouples credential
information from source code, allowing you to push the latter to your version
control system without any security risk.

You will need to retrieve the following:

* APIKey(equivalent of keystone user name)
* API Secret Key(equivalent of keystone password)
* a valid Keystone identity URL

### Authentication

Once you have access to your credentials, you can begin plugging them into
eclcloud. The next step is authentication, and this is handled by a base
"Provider" struct. To get one, you can either pass in your credentials
explicitly, or tell eclcloud to use environment variables:

```go
import (
  "github.com/nttcom/eclcloud"
  "github.com/nttcom/eclcloud"
  "github.com/nttcom/eclcloud/utils"
)

// Option 1: Pass in the values yourself
opts := eclcloud.AuthOptions{
  IdentityEndpoint: "https://{your keystone url}/v3",
  Username: "{apikey}",
  Password: "{api secret key}",
}

// Option 2: Use a utility function to retrieve all your environment variables
opts, err := ecl.AuthOptionsFromEnv()
```

Once you have the `opts` variable, you can pass it in and get back a
`ProviderClient` struct:

```go
provider, err := ecl.AuthenticatedClient(opts)
```

The `ProviderClient` is the top-level client that all of your Enterprise Cloud services derive from. The provider contains all of the authentication details that allow
your Go code to access the API - such as the base URL and token ID.

### Provision a server

Once we have a base Provider, we inject it as a dependency into each Enterprise Cloud service. In order to work with the Compute API, we need a Compute service client; which can be created like so:

```go
client, err := ecl.NewComputeV2(provider, eclcloud.EndpointOpts{
  Region: os.Getenv("OS_REGION_NAME"),
})
```

We then use this `client` for any Compute API operation we want. In our case,
we want to provision a new server - so we invoke the `Create` method and pass
in the flavor ID (hardware specification) and image ID (operating system) we're
interested in:

```go
import "github.com/nttcom/eclcloud/compute/v2/servers"

server, err := servers.Create(client, servers.CreateOpts{
  Name:      "My new server!",
  FlavorRef: "flavor_id",
  ImageRef:  "image_id",
}).Extract()
```

The above code sample creates a new server with the parameters, and embodies the
new resource in the `server` variable (a
[`servers.Server`](http://godoc.org/github.com/nttcom/eclcloud) struct).

## Advanced Usage

Have a look at the [FAQ](./docs/FAQ.md) for some tips on customizing the way eclcloud works.

## Backwards-Compatibility Guarantees

None. Vendor it and write tests covering the parts you use.

## Contributing

See the [contributing guide](./.github/CONTRIBUTING.md).

## Help and feedback

If you're struggling with something or have spotted a potential bug, feel free to submit an issue to our [bug tracker](https://github.com/nttcom/eclcloud/issues).
