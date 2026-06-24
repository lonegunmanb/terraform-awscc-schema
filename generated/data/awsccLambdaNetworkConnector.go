package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaNetworkConnector = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the network connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description": "The network configuration for the connector. Specify a VpcEgressConfiguration to enable outbound traffic routing through your VPC.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vpc_egress_configuration": {
              "computed": true,
              "description": "The VPC egress configuration for the network connector. Specifies the subnets, security groups, and network protocol for routing outbound traffic through your VPC.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "associated_compute_resource_types": {
                    "computed": true,
                    "description": "The types of Lambda compute resources that can use this connector. Currently, only MicroVm is supported.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "network_protocol": {
                    "computed": true,
                    "description": "The network protocol for the connector. Specify IPv4 for IPv4-only networking, or DualStack for both IPv4 and IPv6.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "security_group_ids": {
                    "computed": true,
                    "description": "The IDs of the VPC security groups to attach to the ENIs. Specify 0 to 5 security groups. All security groups must be in the same VPC as the subnets.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_ids": {
                    "computed": true,
                    "description": "The IDs of the VPC subnets where Lambda provisions elastic network interfaces (ENIs). Specify 1 to 16 subnets. All subnets must be in the same VPC.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A unique name for the network connector within your account and Region. Must be 1 to 64 alphanumeric characters, hyphens, or underscores.",
        "description_kind": "plain",
        "type": "string"
      },
      "operator_role": {
        "computed": true,
        "description": "The ARN of the IAM role that Lambda assumes to manage elastic network interfaces in your VPC. This role must have permissions for ec2:CreateNetworkInterface and related describe operations.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the network connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the network connector. Use tags to categorize network connectors for cost allocation, access control, or operational management.",
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
      }
    },
    "description": "Data Source schema for AWS::Lambda::NetworkConnector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLambdaNetworkConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaNetworkConnector), &result)
	return &result
}
