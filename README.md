Terraform Enterprise Cloud Provider
===================================

Maintainers
-----------

This provider plugin is maintained by:

* Keiichi Hikita ([@keiichi-hikita](https://github.com/keiichi-hikita))
* Yuma Maki ([@ymaki](https://github.com/ymaki))
* Rinto Shimizu([@rintooo](https://github.com/rintooo))

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.11x
- [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to anywhere you want: 

```sh
$ git clone https://github.com/nttcom/terraform-provider-ecl 
```

Enter the provider directory and build the provider

```sh
$ cd terraform-provider-ecl
$ make build
```

Using the provider
----------------------
You can browse the documentation within this repo [here](https://github.com/nttcom/terraform-provider-ecl/tree/master/website/docs).

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ cp terraform-provider-ecl $GOPATH/bin/terraform-provider-ecl
...
```
