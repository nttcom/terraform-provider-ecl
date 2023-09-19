/*
Package qos_options provides information of several service
in the Enterprise Cloud Compute service

Example to List QoS Options

	listOpts := qos_options.ListOpts{
		QoSType: "guarantee",
	}

	allPages, err := qos_options.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allQoSOptions, err := qos_options.ExtractQoSOptions(allPages)
	if err != nil {
		panic(err)
	}

	for _, qosOption := range allQoSOptions {
		fmt.Printf("%+v", qosOption)
	}

Example to Show QoS Option

	id := "02dc9a22-129c-4b12-9936-4080f6a7ae44"
	qosOption, err := qos_options.Get(client, id).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Print(qosOption)

*/
package qos_options
