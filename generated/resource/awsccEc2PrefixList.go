package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2PrefixList = `{
  "block": {
    "attributes": {
      "address_family": {
        "description": "Ip Version of Prefix List.",
        "description_kind": "plain",
        "required": true,
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
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_entries": {
        "computed": true,
        "description": "Max Entries of Prefix List.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Name of Prefix List.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for Prefix List",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "version": {
        "computed": true,
        "description": "Version of Prefix List.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Resource schema of AWS::EC2::PrefixList Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2PrefixListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2PrefixList), &result)
	return &result
}
