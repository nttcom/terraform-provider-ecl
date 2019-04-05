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
                                "region": "%[1]s",
                                "region_id": "%[1]s",
                                "url": "%[2]sv2/01234567890123456789abcdefabcdef"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcde1",
                                "interface": "internal",
                                "region": "%[1]s",
                                "region_id": "%[1]s",
                                "url": "%[2]sv2/01234567890123456789abcdefabcdef"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcde2",
                                "interface": "public",
                                "region": "%[1]s",
                                "region_id": "%[1]s",
                                "url": "%[2]sv2/01234567890123456789abcdefabcdef"
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
                                "region": "%[1]s",
                                "region_id": "%[1]s",
                                "url": "%[2]s"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcde5",
                                "interface": "internal",
                                "region": "%[1]s",
                                "region_id": "%[1]s",
                                "url": "%[2]s"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcde6",
                                "interface": "admin",
                                "region": "%[1]s",
                                "region_id": "%[1]s",
                                "url": "%[2]s"
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
                                "region": "%[1]s",
                                "region_id": "%[1]s",
                                "url": "%[2]sv3"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcde9",
                                "interface": "internal",
                                "region": "%[1]s",
                                "region_id": "%[1]s",
                                "url": "%[2]sv3"
                            },
                            {
                                "id": "1234567890abcdef1234567890abcdea",
                                "interface": "public",
                                "region": "%[1]s",
                                "region_id": "%[1]s",
                                "url": "%[2]sv3"
                            }
                        ],
                        "id": "1234567890abcdef1234567890abcdeb",
                        "name": "keystone",
                        "type": "identity"
                    },
                    {
                        "endpoints":[
                            {
                                "region_id": "%[1]s",
                                "url": "%[2]sv2/01234567890123456789abcdefabcdef",
                                "region": "%[1]s",
                                "interface": "public",
                                "id": "1234567890abcdef1234567890abcdec"
                            },
                            {
                                "region_id": "%[1]s",
                                "url": "%[2]sv2/01234567890123456789abcdefabcdef",
                                "region": "%[1]s",
                                "interface": "admin",
                                "id": "1234567890abcdef1234567890abcded"
                            },
                            {
                                "region_id": "%[1]s",
                                "url": "%[2]sv2/01234567890123456789abcdefabcdef",
                                "region": "%[1]s",
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
                                "region_id": "%[1]s",
                                "url": "%[2]sv1/01234567890123456789abcdefabcdef",
                                "region": "%[1]s",
                                "interface": "public",
                                "id": "1234567890abcdef1234567890abcdd0"
                            },
                            {
                                "region_id": "%[1]s",
                                "url": "%[2]sv1/01234567890123456789abcdefabcdef",
                                "region": "%[1]s",
                                "interface": "internal",
                                "id": "1234567890abcdef1234567890abcdd1"
                            },
                            {
                                "region_id": "%[1]s",
                                "url": "%[2]sv1/01234567890123456789abcdefabcdef",
                                "region": "%[1]s",
                                "interface": "admin",
                                "id": "1234567890abcdef1234567890abcdd2"
                            }
                        ],
                        "type": "volume",
                        "id": "1234567890abcdef1234567890abcdd3",
                        "name": "cinder"
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
