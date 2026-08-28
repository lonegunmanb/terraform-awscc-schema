package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDocdbDbClusterParameterGroup = `{
  "block": {
    "attributes": {
      "description": {
        "description": "The description for the DB cluster parameter group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "family": {
        "description": "The DB cluster parameter group family name (e.g. docdb5.0).",
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
        "description": "The name of the DB cluster parameter group. If omitted, CloudFormation generates a unique name. The name is stored as lowercase.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "description": "An object containing key-value pairs of parameters to set for the DB cluster parameter group.",
        "description_kind": "plain",
        "required": true,
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::DocDB::DBClusterParameterGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDocdbDbClusterParameterGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDocdbDbClusterParameterGroup), &result)
	return &result
}
