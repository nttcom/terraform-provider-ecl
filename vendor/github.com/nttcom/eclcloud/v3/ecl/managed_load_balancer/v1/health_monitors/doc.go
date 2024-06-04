/*
Package health_monitors contains functionality for working with ECL Managed Load Balancer resources.

Example to list health monitors

	listOpts := health_monitors.ListOpts{}

	allPages, err := health_monitors.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allHealthMonitors, err := health_monitors.ExtractHealthMonitors(allPages)
	if err != nil {
		panic(err)
	}

	for _, healthMonitor := range allHealthMonitors {
		fmt.Printf("%+v\n", healthMonitor)
	}

Example to create a health monitor


	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	createOpts := health_monitors.CreateOpts{
		Name: "health_monitor",
		Description: "description",
		Tags: tags,
		Port: 80,
		Protocol: "http",
		Interval: 5,
		Retry: 3,
		Timeout: 5,
		Path: "/health",
		HttpStatusCode: "200-299",
		LoadBalancerID: "67fea379-cff0-4191-9175-de7d6941a040",
	}

	healthMonitor, err := health_monitors.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", healthMonitor)

Example to show a health monitor

	showOpts := health_monitors.ShowOpts{}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	healthMonitor, err := health_monitors.Show(managedLoadBalancerClient, id, showOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", healthMonitor)

Example to update a health monitor

	name := "health_monitor"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	updateOpts := health_monitors.UpdateOpts{
		Name: &name,
		Description: &description,
		Tags: &tags,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	healthMonitor, err := health_monitors.Update(managedLoadBalancerClient, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", healthMonitor)

Example to delete a health monitor

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := health_monitors.Delete(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to create staged health monitor configurations

	createStagedOpts := health_monitors.CreateStagedOpts{
		Port: 80,
		Protocol: "http",
		Interval: 5,
		Retry: 3,
		Timeout: 5,
		Path: "/health",
		HttpStatusCode: "200-299",
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	healthMonitorConfigurations, err := health_monitors.CreateStaged(managedLoadBalancerClient, id, createStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", healthMonitorConfigurations)

Example to show staged health monitor configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	healthMonitorConfigurations, err := health_monitors.ShowStaged(managedLoadBalancerClient, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", healthMonitorConfigurations)

Example to update staged health monitor configurations

	port := 80
	protocol := "http"
	interval := 5
	retry := 3
	timeout := 5
	path := "/health"
	httpStatusCode := "200-299"
	updateStagedOpts := health_monitors.UpdateStagedOpts{
		Port: &port,
		Protocol: &protocol,
		Interval: &interval,
		Retry: &retry,
		Timeout: &timeout,
		Path: &path,
		HttpStatusCode: &httpStatusCode,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	healthMonitorConfigurations, err := health_monitors.UpdateStaged(managedLoadBalancerClient, updateStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", healthMonitorConfigurations)

Example to cancel staged health monitor configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := health_monitors.CancelStaged(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package health_monitors
