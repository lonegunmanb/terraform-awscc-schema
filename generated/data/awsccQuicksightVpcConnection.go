package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightVpcConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the VPC connection.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "\u003cp\u003eThe time that the VPC connection was created.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "dns_resolvers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "\u003cp\u003eThe time that the VPC connection was last updated.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_interfaces": {
        "computed": true,
        "description": "\u003cp\u003eA list of network interfaces.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description": "\u003cp\u003eThe availability zone that the network interface resides in.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "error_message": {
              "computed": true,
              "description": "\u003cp\u003eAn error message.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "network_interface_id": {
              "computed": true,
              "description": "\u003cp\u003eThe network interface ID.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subnet_id": {
              "computed": true,
              "description": "\u003cp\u003eThe subnet ID associated with the network interface.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "\u003cp\u003eTag key.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "vpc_connection_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon EC2 VPC ID associated with the VPC connection.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QuickSight::VPCConnection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightVpcConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightVpcConnection), &result)
	return &result
}
