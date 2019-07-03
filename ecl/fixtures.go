package ecl

var fakeKeystonePostTmpl = `
request:
    method: POST
response:
    code: 201
    body: >
        {
            "token": {
                "audit_ids": [
                    "DummyIds123456789abcde"
                ],
                "catalog": [
                    {
                        "endpoints": [
                            {
                                "id": "1234567890abcdef1234567890abcde0",
                                "interface": "admin",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]sv2/01234567890123456789abcdefabcdef"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcde1",
                                "interface": "internal",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]sv2/01234567890123456789abcdefabcdef"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcde2",
                                "interface": "public",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]sv2/01234567890123456789abcdefabcdef"
                            }
                        ],
                        "id": "1234567890abcdef1234567890abcde3",
                        "name": "nova",
                        "type": "compute"
                    },
                    {
                        "endpoints": [
                            {
                                "id": "1234567890abcdef1234567890abcde4",
                                "interface": "public",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]s"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcde5",
                                "interface": "internal",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]s"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcde6",
                                "interface": "admin",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]s"
                            }
                        ],
                        "id": "1234567890abcdef1234567890abcde7",
                        "name": "network",
                        "type": "network"
                    },
                    {
                        "endpoints": [
                            {
                                "id": "1234567890abcdef1234567890abcde8",
                                "interface": "admin",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]sv3"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcde9",
                                "interface": "internal",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]sv3"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcdea",
                                "interface": "public",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]sv3"
                            }
                        ],
                        "id": "1234567890abcdef1234567890abcdeb",
                        "name": "keystone",
                        "type": "identity"
                    },
                    {
                        "endpoints":[
                            {
                                "region_id": "RegionOne",
                                "url": "%[1]sv2/01234567890123456789abcdefabcdef",
                                "region": "RegionOne",
                                "interface": "public",
                                "id": "1234567890abcdef1234567890abcdec"
                            },
                            {
                                "region_id": "RegionOne",
                                "url": "%[1]sv2/01234567890123456789abcdefabcdef",
                                "region": "RegionOne",
                                "interface": "admin",
                                "id": "1234567890abcdef1234567890abcded"
                            },
                            {
                                "region_id": "RegionOne",
                                "url": "%[1]sv2/01234567890123456789abcdefabcdef",
                                "region": "RegionOne",
                                "interface": "internal",
                                "id": "1234567890abcdef1234567890abcdee"
                            }
                        ],
                        "type": "volumev2",
                        "id": "1234567890abcdef1234567890abcdef",
                        "name": "cinder2"
                    },
                    {
                        "endpoints":[
                            {
                                "region_id": "RegionOne",
                                "url": "%[1]sv1/01234567890123456789abcdefabcdef",
                                "region": "RegionOne",
                                "interface": "public",
                                "id": "1234567890abcdef1234567890abcdd0"
                            },
                            {
                                "region_id": "RegionOne",
                                "url": "%[1]sv1/01234567890123456789abcdefabcdef",
                                "region": "RegionOne",
                                "interface": "internal",
                                "id": "1234567890abcdef1234567890abcdd1"
                            },
                            {
                                "region_id": "RegionOne",
                                "url": "%[1]sv1/01234567890123456789abcdefabcdef",
                                "region": "RegionOne",
                                "interface": "admin",
                                "id": "1234567890abcdef1234567890abcdd2"
                            }
                        ],
                        "type": "volume",
                        "id": "1234567890abcdef1234567890abcdd3",
                        "name": "cinder"
                    },
                    {
                        "endpoints": [
                            {
                                "id": "c4c383a719cb489d8210328e17659621",
                                "interface": "public",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]s"
                            },
                            {
                                "id": "c4c383a719cb489d8210328e17659622",
                                "interface": "internal",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]s"
                            },
                            {
                                "id": "c4c383a719cb489d8210328e17659623",
                                "interface": "admin",
                                "region": "RegionOne",
                                "region_id": "RegionOne",
                                "url": "%[1]s"
                            }
                        ],
                        "id": "c4c383a719cb489d8210328e17659620",
                        "name": "virtual-network-appliance",
                        "type": "virtual-network-appliance"
                    }
                ],
                "expires_at": "2018-11-28T02:48:52.111201Z",
                "issued_at": "2018-11-28T01:48:52.111227Z",
                "methods": [
                    "password"
                ],
                "project": {
                    "domain": {
                        "id": "default",
                        "name": "Default"
                    },
                    "id": "01234567890123456789abcdefabcdef",
                    "name": "FakeTenant"
                },
                "roles": [
                    {
                        "id": "0123456789abcdef0123456789abcdef",
                        "name": "_member_"
                    }
                ],
                "user": {
                    "domain": {
                        "id": "default",
                        "name": "Default"
                    },
                    "id": "abcdef0123456789abcdef0123456789",
                    "name": "ThisIsADummyTenantUsername"
                }
            }
        }
`
