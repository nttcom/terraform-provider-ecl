/*
Package rules contains functionality for working with ECL Managed Load Balancer resources.

Example to list rules

	listOpts := rules.ListOpts{}

	allPages, err := rules.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allRules, err := rules.ExtractRules(allPages)
	if err != nil {
		panic(err)
	}

	for _, rule := range allRules {
		fmt.Printf("%+v\n", rule)
	}

Example to create a rule

	condition := rules.CreateOptsCondition{
		PathPatterns: []string{"^/statics/"},
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	createOpts := rules.CreateOpts{
		Name: "rule",
		Description: "description",
		Tags: tags,
		Priority: 1,
		TargetGroupID: "29527a3c-9e5d-48b7-868f-6442c7d21a95",
		PolicyID: "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
		Conditions: &condition,
	}

	rule, err := rules.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", rule)

Example to show a rule

	showOpts := rules.ShowOpts{}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	rule, err := rules.Show(managedLoadBalancerClient, id, showOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", rule)

Example to update a rule

	name := "rule"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	updateOpts := rules.UpdateOpts{
		Name: &name,
		Description: &description,
		Tags: &tags,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	rule, err := rules.Update(managedLoadBalancerClient, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", rule)

Example to delete a rule

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := rules.Delete(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to create staged rule configurations

	condition := rules.CreateStagedOptsCondition{
		PathPatterns: []string{"^/statics/"},
	}
	createStagedOpts := rules.CreateStagedOpts{
		Priority: 1,
		TargetGroupID: "29527a3c-9e5d-48b7-868f-6442c7d21a95",
		Conditions: &condition,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	ruleConfigurations, err := rules.CreateStaged(managedLoadBalancerClient, id, createStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ruleConfigurations)

Example to show staged rule configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	ruleConfigurations, err := rules.ShowStaged(managedLoadBalancerClient, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ruleConfigurations)

Example to update staged rule configurations

	condition := rules.UpdateStagedOptsCondition{
		PathPatterns: &[]string{"^/statics/"},
	}

	priority := 1
	targetGroupID := "29527a3c-9e5d-48b7-868f-6442c7d21a95"
	updateStagedOpts := rules.UpdateStagedOpts{
		Priority: &priority,
		TargetGroupID: &targetGroupID,
		Conditions: &condition,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	ruleConfigurations, err := rules.UpdateStaged(managedLoadBalancerClient, updateStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ruleConfigurations)

Example to cancel staged rule configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := rules.CancelStaged(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package rules
