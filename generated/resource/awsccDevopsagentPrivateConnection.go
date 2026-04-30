package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsagentPrivateConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Private Connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate": {
        "computed": true,
        "description": "Certificate for the Private Connection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_expiry_time": {
        "computed": true,
        "description": "The expiry time of the certificate associated with the Private Connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_configuration": {
        "description": "The connection configuration for the Private Connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "self_managed": {
              "computed": true,
              "description": "Configuration for a self-managed Private Connection.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "resource_configuration_id": {
                    "computed": true,
                    "description": "The ARN of the Resource Configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "service_managed": {
              "computed": true,
              "description": "Configuration for a service-managed Private Connection.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "host_address": {
                    "computed": true,
                    "description": "IP address or DNS name of the target resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ip_address_type": {
                    "computed": true,
                    "description": "IP address type of the service-managed Resource Gateway.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ipv_4_addresses_per_eni": {
                    "computed": true,
                    "description": "Number of IPv4 addresses in each ENI for the service-managed Resource Gateway.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "port_ranges": {
                    "computed": true,
                    "description": "TCP port ranges that a consumer can use to access the resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "security_group_ids": {
                    "computed": true,
                    "description": "Security groups to attach to the service-managed Resource Gateway.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_ids": {
                    "computed": true,
                    "description": "Subnets that the service-managed Resource Gateway will span.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "vpc_id": {
                    "computed": true,
                    "description": "VPC to create the service-managed Resource Gateway in.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Unique name for this Private Connection within the account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the Private Connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::DevOpsAgent::PrivateConnection",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDevopsagentPrivateConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsagentPrivateConnection), &result)
	return &result
}
