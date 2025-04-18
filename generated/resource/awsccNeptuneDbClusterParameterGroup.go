package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNeptuneDbClusterParameterGroup = `{
  "block": {
    "attributes": {
      "description": {
        "description": "Provides the customer-specified description for this DB cluster parameter group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "family": {
        "description": "Must be neptune1 for engine versions prior to 1.2.0.0, or neptune1.2 for engine version 1.2.0.0 and higher.",
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
        "description": "Provides the name of the DB cluster parameter group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "description": "An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The list of tags for the cluster parameter group.",
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
    "description": "The AWS::Neptune::DBClusterParameterGroup resource creates a new Amazon Neptune DB cluster parameter group",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNeptuneDbClusterParameterGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNeptuneDbClusterParameterGroup), &result)
	return &result
}
