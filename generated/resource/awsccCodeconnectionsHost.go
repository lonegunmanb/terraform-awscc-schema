package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodeconnectionsHost = `{
  "block": {
    "attributes": {
      "host_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the host.",
        "description_kind": "plain",
        "type": "string"
      },
      "host_id": {
        "computed": true,
        "description": "The server-generated unique identifier for the host.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the host.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provider_endpoint": {
        "description": "The endpoint of the infrastructure where your provider type is installed.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provider_type": {
        "description": "The name of the installed provider to be associated with your connection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the host.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to apply to the host.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "vpc_configuration": {
        "computed": true,
        "description": "The VPC configuration provisioned for the host.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "The ID of the security group or security groups associated with the Amazon VPC.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "The ID of the subnet or subnets associated with the Amazon VPC.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "tls_certificate": {
              "computed": true,
              "description": "The value of the Transport Layer Security (TLS) certificate associated with the infrastructure where your provider type is installed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_id": {
              "computed": true,
              "description": "The ID of the Amazon VPC connected to the infrastructure where your provider type is installed.",
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
    "description": "A resource that represents the infrastructure where a third-party provider is installed. You create one host for all connections to that provider.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCodeconnectionsHostSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodeconnectionsHost), &result)
	return &result
}
