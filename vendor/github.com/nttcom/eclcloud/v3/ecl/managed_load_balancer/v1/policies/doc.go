/*
Package policies contains functionality for working with ECL Managed Load Balancer resources.

Example to list policies

	listOpts := policies.ListOpts{}

	allPages, err := policies.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allPolicies, err := policies.ExtractPolicies(allPages)
	if err != nil {
		panic(err)
	}

	for _, policy := range allPolicies {
		fmt.Printf("%+v\n", policy)
	}

Example to create a policy

	serverNameIndication1 := policies.CreateOptsServerNameIndication{
		ServerName: "*.example.com",
		InputType: "fixed",
		Priority: 1,
		CertificateID: "fdfed344-e8ab-4f20-bd62-a4039453a389",
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	createOpts := policies.CreateOpts{
		Name: "policy",
		Description: "description",
		Tags: tags,
		Algorithm: "round-robin",
		Persistence: "cookie",
		PersistenceTimeout: 525600,
		IdleTimeout: 600,
		SorryPageUrl: "https://example.com/sorry",
		SourceNat: "enable",
		ServerNameIndications: &[]policies.CreateOptsServerNameIndication{serverNameIndication1},
		CertificateID: "f57a98fe-d63e-4048-93a0-51fe163f30d7",
		HealthMonitorID: "dd7a96d6-4e66-4666-baca-a8555f0c472c",
		ListenerID: "68633f4f-f52a-402f-8572-b8173418904f",
		DefaultTargetGroupID: "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
		BackupTargetGroupID: "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
		TLSPolicyID: "4ba79662-f2a1-41a4-a3d9-595799bbcd86",
		LoadBalancerID: "67fea379-cff0-4191-9175-de7d6941a040",
	}

	policy, err := policies.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", policy)

Example to show a policy

	showOpts := policies.ShowOpts{}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	policy, err := policies.Show(managedLoadBalancerClient, id, showOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", policy)

Example to update a policy

	name := "policy"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	updateOpts := policies.UpdateOpts{
		Name: &name,
		Description: &description,
		Tags: &tags,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	policy, err := policies.Update(managedLoadBalancerClient, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", policy)

Example to delete a policy

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := policies.Delete(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to create staged policy configurations

	serverNameIndication1 := policies.CreateStagedOptsServerNameIndication{
		ServerName: "*.example.com",
		InputType: "fixed",
		Priority: 1,
		CertificateID: "fdfed344-e8ab-4f20-bd62-a4039453a389",
	}
	createStagedOpts := policies.CreateStagedOpts{
		Algorithm: "round-robin",
		Persistence: "cookie",
		PersistenceTimeout: 525600,
		IdleTimeout: 600,
		SorryPageUrl: "https://example.com/sorry",
		SourceNat: "enable",
		ServerNameIndications: &[]policies.CreateStagedOptsServerNameIndication{serverNameIndication1},
		CertificateID: "f57a98fe-d63e-4048-93a0-51fe163f30d7",
		HealthMonitorID: "dd7a96d6-4e66-4666-baca-a8555f0c472c",
		ListenerID: "68633f4f-f52a-402f-8572-b8173418904f",
		DefaultTargetGroupID: "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
		BackupTargetGroupID: "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
		TLSPolicyID: "4ba79662-f2a1-41a4-a3d9-595799bbcd86",
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	policyConfigurations, err := policies.CreateStaged(managedLoadBalancerClient, id, createStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", policyConfigurations)

Example to show staged policy configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	policyConfigurations, err := policies.ShowStaged(managedLoadBalancerClient, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", policyConfigurations)

Example to update staged policy configurations

	serverNameIndication1ServerName := "*.example.com"
	serverNameIndication1InputType := "fixed"
	serverNameIndication1Priority := 1
	serverNameIndication1CertificateID := "fdfed344-e8ab-4f20-bd62-a4039453a389"
	serverNameIndication1 := policies.UpdateStagedOptsServerNameIndication{
		ServerName: &serverNameIndication1ServerName,
		InputType: &serverNameIndication1InputType,
		Priority: &serverNameIndication1Priority,
		CertificateID: &serverNameIndication1CertificateID,
	}

	algorithm := "round-robin"
	persistence := "cookie"
	persistenceTimeout := 525600
	idleTimeout := 600
	sorryPageUrl := "https://example.com/sorry"
	sourceNat := "enable"
	certificateID := "f57a98fe-d63e-4048-93a0-51fe163f30d7"
	healthMonitorID := "dd7a96d6-4e66-4666-baca-a8555f0c472c"
	listenerID := "68633f4f-f52a-402f-8572-b8173418904f"
	defaultTargetGroupID := "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
	backupTargetGroupID := "f1a117f1-f8df-ce07-6c8c-4bbf103059b6"
	tlsPolicyID := "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
	updateStagedOpts := policies.UpdateStagedOpts{
		Algorithm: &algorithm,
		Persistence: &persistence,
		PersistenceTimeout: &persistenceTimeout,
		IdleTimeout: &idleTimeout,
		SorryPageUrl: &sorryPageUrl,
		SourceNat: &sourceNat,
		ServerNameIndications: &[]policies.UpdateStagedOptsServerNameIndication{serverNameIndication1},
		CertificateID: &certificateID,
		HealthMonitorID: &healthMonitorID,
		ListenerID: &listenerID,
		DefaultTargetGroupID: &defaultTargetGroupID,
		BackupTargetGroupID: &backupTargetGroupID,
		TLSPolicyID: &tlsPolicyID,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	policyConfigurations, err := policies.UpdateStaged(managedLoadBalancerClient, updateStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", policyConfigurations)

Example to cancel staged policy configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := policies.CancelStaged(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package policies
