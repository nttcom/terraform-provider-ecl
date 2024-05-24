# How to configure Managed Load Balancer resources and apply them to an instance.

This provides a template for creating a new Managed Load Balancer instance on Enterprise Cloud 2.0.

First set the required environment variables for the Enterprise Cloud 2.0 provider by

```
source openrc
```

There is a sample of openrc as openrc.sample, so copy, rename it, and fill blank parameters afterward.

Secondary run with a command like this:

```
cd resources
terraform apply
```

The network and the subnet are created, and the load balancer and related resources are configured (like candidate-config) by this command.

Thirdly run with a command like this:

```
cd action
terraform apply
```

Configurations of the load balancer and related resources are applied to the instance by this command.
