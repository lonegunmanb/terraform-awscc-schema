package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVpclatticeResourceConfiguration = `{
  "block": {
    "attributes": {
      "allow_association_to_sharable_service_network": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_verification_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "port_ranges": {
        "computed": true,
        "description": "The port ranges the resource configuration listens on, as single ports or ranges (for example, 80 or 1001-2000). Applies to SINGLE, GROUP, CHILD, and CIDR resource configurations. For a CIDR resource configuration with ProtocolType TCP_UDP, these ranges cover both TCP and UDP traffic, so include port 53 to carry DNS to the network segment.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "protocol_type": {
        "computed": true,
        "description": "The protocol the resource configuration accepts. SINGLE and GROUP resource configurations support TCP. CIDR resource configurations support TCP and TCP_UDP; use TCP_UDP to carry UDP traffic such as DNS on port 53.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_configuration_auth_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_configuration_definition": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn_resource": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "cidr_resource": {
              "computed": true,
              "description": "The CIDR ranges that make up a CIDR resource configuration. Only valid when ResourceConfigurationType is CIDR, and the resource gateway must have ResourceConfigDnsResolution set to IN_VPC.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr_ranges": {
                    "computed": true,
                    "description": "The IPv4 CIDR ranges reachable through this resource configuration, in CIDR notation (for example, 10.0.0.0/24).",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "dns_resource": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ip_address_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ip_resource": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_configuration_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_configuration_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_configuration_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::VpcLattice::ResourceConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccVpclatticeResourceConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVpclatticeResourceConfiguration), &result)
	return &result
}
