package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityagentPrivateConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the private connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate": {
        "computed": true,
        "description": "The PEM-encoded certificate chain for the private connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_expiry_time": {
        "computed": true,
        "description": "The date and time the connection's certificate expires, in UTC format.",
        "description_kind": "plain",
        "type": "string"
      },
      "dns_resolution": {
        "computed": true,
        "description": "The DNS resolution mode for the resource gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "host_address": {
        "computed": true,
        "description": "The IP address or DNS name of the target resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mode": {
        "computed": true,
        "description": "The configuration for the private connection. Specify either a service-managed or a self-managed mode.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "self_managed": {
              "computed": true,
              "description": "The configuration for a self-managed private connection, where you manage your own resource configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "resource_configuration_id": {
                    "computed": true,
                    "description": "The identifier or ARN of the resource configuration.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "service_managed": {
              "computed": true,
              "description": "The configuration for a service-managed private connection, where the service manages the resource gateway lifecycle.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dns_resolution": {
                    "computed": true,
                    "description": "The DNS resolution mode for the resource gateway.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "host_address": {
                    "computed": true,
                    "description": "The IP address or DNS name of the target resource.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ip_address_type": {
                    "computed": true,
                    "description": "The IP address type of the service-managed resource gateway.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ipv_4_addresses_per_eni": {
                    "computed": true,
                    "description": "The number of IPv4 addresses in each elastic network interface for the service-managed resource gateway.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "port_ranges": {
                    "computed": true,
                    "description": "The TCP port ranges that a consumer can use to access the resource.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "security_group_ids": {
                    "computed": true,
                    "description": "The security groups to attach to the service-managed resource gateway.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_ids": {
                    "computed": true,
                    "description": "The subnets that the service-managed resource gateway spans.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "vpc_id": {
                    "computed": true,
                    "description": "The VPC to create the service-managed resource gateway in.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "private_connection_name": {
        "computed": true,
        "description": "A unique name for the private connection within your account.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_configuration_id": {
        "computed": true,
        "description": "The identifier or ARN of the VPC Lattice resource configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_gateway_id": {
        "computed": true,
        "description": "The identifier or ARN of the VPC Lattice resource gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the private connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags attached to the private connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "type": {
        "computed": true,
        "description": "The type of the private connection, indicating whether it is service-managed or self-managed.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The identifier of the VPC the resource gateway is created in.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityAgent::PrivateConnection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityagentPrivateConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityagentPrivateConnection), &result)
	return &result
}
