/*
Package load_balancer_syslog_servers contains functionality for working with
ECL Load Balancer Syslog Server resources.

Example to List Load Balancer Syslog Servers

	listOpts := load_balancer_syslog_servers.ListOpts{
		Status: "ACTIVE",
	}

	allPages, err := load_balancer_syslog_servers.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allLoadBalancerSyslogServers, err := load_balancer_syslog_servers.ExtractLoadBalancerSyslogServers(allPages)
	if err != nil {
		panic(err)
	}

	for _, loadBalancerSyslogServer := range allLoadBalancerSyslogServers {
		fmt.Printf("%+v\n", loadBalancerSyslogServer)
	}


Example to Show Load Balancer Syslog Server

	loadBalancerSyslogServerID := "9ab7ab3c-38a6-417c-926b-93772c4eb2f9"

	loadBalancerSyslogServer, err := load_balancer_syslog_servers.Get(networkClient, loadBalancerSyslogServerID).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", loadBalancerSyslogServer)


Example to Create a Load Balancer Syslog Server

	priority := 20

	createOpts := load_balancer_syslog_servers.CreateOpts{
		AclLogging:                  "DISABLED",
		AppflowLogging:              "DISABLED",
		DateFormat:                  "MMDDYYYY",
		Description:                 "test",
		IPAddress:                   "120.120.120.30",
		LoadBalancerID:              "4f6ebc24-f768-485b-99ef-f308063d0209",
		LogFacility:                 "LOCAL3",
		LogLevel:                    "DEBUG",
		Name:                        "first_syslog_server",
		PortNumber:                  514,
		Priority:                    &priority,
		TcpLogging:                  "ALL",
		TenantID:                    "b58531f716614e82a9bf001571c8bb15",
		TimeZone:                    "LOCAL_TIME",
		TransportType:               "UDP",
		UserConfigurableLogMessages: "NO",
	}

	loadBalancerSyslogServer, err := load_balancer_syslog_servers.Create(networkClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a Load Balancer Syslog Server

	loadBalancerSyslogServerID := "9ab7ab3c-38a6-417c-926b-93772c4eb2f9"
	description := "new_description"

	updateOpts := load_balancer_syslog_servers.UpdateOpts{
		Description:           &description,
	}

	loadBalancerSyslogServer, err := load_balancer_syslog_servers.Update(networkClient, loadBalancerSyslogServerID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Load Balancer Syslog Server

	loadBalancerSyslogServerID := "13762eaf-9564-4c94-a106-98ece9fa189e"
	err := load_balancer_syslog_servers.Delete(networkClient, loadBalancerSyslogServerID).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package load_balancer_syslog_servers
