package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOmicsRunGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_cpus": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_duration": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_gpus": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_runs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of resource tags",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Omics::RunGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOmicsRunGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOmicsRunGroup), &result)
	return &result
}
