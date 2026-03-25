/*
Package security_group_rules contains functionality for working with ECL Security Group Rule resources.

Security Group Rules define specific ingress and egress traffic rules for Security Groups.

Example to List Security Group Rules

	listOpts := security_group_rules.ListOpts{
		SecurityGroupID: "security-group-id",
	}

	allPages, err := security_group_rules.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allRules, err := security_group_rules.ExtractSecurityGroupRules(allPages)
	if err != nil {
		panic(err)
	}

	for _, rule := range allRules {
		fmt.Printf("%+v\n", rule)
	}

Example to Create a Security Group Rule

	createOpts := security_group_rules.CreateOpts{
		Direction:       "ingress",
		SecurityGroupID: "security-group-id",
		Ethertype:       "IPv4",
		Protocol:        "tcp",
		PortRangeMin:    &[]int{22}[0],
		PortRangeMax:    &[]int{22}[0],
		RemoteIPPrefix:  &[]string{"0.0.0.0/0"}[0],
	}

	rule, err := security_group_rules.Create(networkClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Security Group Rule

	ruleID := "rule-id"
	err := security_group_rules.Delete(networkClient, ruleID).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package security_group_rules
