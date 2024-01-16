package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2PrefixList = `{
  "block": {
    "attributes": {
      "address_family": {
        "computed": true,
        "description": "Ip Version of Prefix List.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Prefix List.",
        "description_kind": "plain",
        "type": "string"
      },
      "entries": {
        "computed": true,
        "description": "Entries of Prefix List.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_entries": {
        "computed": true,
        "description": "Max Entries of Prefix List.",
        "description_kind": "plain",
        "type": "number"
      },
      "owner_id": {
        "computed": true,
        "description": "Owner Id of Prefix List.",
        "description_kind": "plain",
        "type": "string"
      },
      "prefix_list_id": {
        "computed": true,
        "description": "Id of Prefix List.",
        "description_kind": "plain",
        "type": "string"
      },
      "prefix_list_name": {
        "computed": true,
        "description": "Name of Prefix List.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for Prefix List",
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
      },
      "version": {
        "computed": true,
        "description": "Version of Prefix List.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::EC2::PrefixList",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2PrefixListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2PrefixList), &result)
	return &result
}
