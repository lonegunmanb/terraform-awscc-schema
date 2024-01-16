package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDetectiveGraph = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Detective graph ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_enable_members": {
        "computed": true,
        "description": "Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @ ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @ ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::Detective::Graph",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDetectiveGraphSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDetectiveGraph), &result)
	return &result
}
