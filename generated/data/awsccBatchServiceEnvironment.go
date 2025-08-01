package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBatchServiceEnvironment = `{
  "block": {
    "attributes": {
      "capacity_limits": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_unit": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "max_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
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
      "service_environment_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_environment_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_environment_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Batch::ServiceEnvironment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBatchServiceEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBatchServiceEnvironment), &result)
	return &result
}
