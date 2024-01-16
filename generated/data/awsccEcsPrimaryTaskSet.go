package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsPrimaryTaskSet = `{
  "block": {
    "attributes": {
      "cluster": {
        "computed": true,
        "description": "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service": {
        "computed": true,
        "description": "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
        "description_kind": "plain",
        "type": "string"
      },
      "task_set_id": {
        "computed": true,
        "description": "The ID or full Amazon Resource Name (ARN) of the task set.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ECS::PrimaryTaskSet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcsPrimaryTaskSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsPrimaryTaskSet), &result)
	return &result
}
