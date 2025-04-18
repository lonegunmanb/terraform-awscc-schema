package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNeptuneDbParameterGroup = `{
  "block": {
    "attributes": {
      "description": {
        "description": "Provides the customer-specified description for this DB parameter group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "family": {
        "description": "Must be ` + "`" + `neptune1` + "`" + ` for engine versions prior to 1.2.0.0, or ` + "`" + `neptune1.2` + "`" + ` for engine version ` + "`" + `1.2.0.0` + "`" + ` and higher.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Provides the name of the DB parameter group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "description": "The parameters to set for this DB parameter group.\n\nThe parameters are expressed as a JSON object consisting of key-value pairs.\n\nChanges to dynamic parameters are applied immediately. During an update, if you have static parameters (whether they were changed or not), it triggers AWS CloudFormation to reboot the associated DB instance without failover.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An optional array of key-value pairs to apply to this DB parameter group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "AWS::Neptune::DBParameterGroup creates a new DB parameter group. This type can be declared in a template and referenced in the DBParameterGroupName parameter of AWS::Neptune::DBInstance",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNeptuneDbParameterGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNeptuneDbParameterGroup), &result)
	return &result
}
