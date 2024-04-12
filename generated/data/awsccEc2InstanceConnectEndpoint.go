package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2InstanceConnectEndpoint = `{
  "block": {
    "attributes": {
      "client_token": {
        "computed": true,
        "description": "The client token of the instance connect endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_connect_endpoint_id": {
        "computed": true,
        "description": "The id of the instance connect endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "preserve_client_ip": {
        "computed": true,
        "description": "If true, the address of the instance connect endpoint client is preserved when connecting to the end resource",
        "description_kind": "plain",
        "type": "bool"
      },
      "security_group_ids": {
        "computed": true,
        "description": "The security group IDs of the instance connect endpoint.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "subnet_id": {
        "computed": true,
        "description": "The subnet id of the instance connect endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags of the instance connect endpoint.",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::EC2::InstanceConnectEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2InstanceConnectEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2InstanceConnectEndpoint), &result)
	return &result
}
