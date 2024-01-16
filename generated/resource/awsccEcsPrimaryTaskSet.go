package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsPrimaryTaskSet = `{
  "block": {
    "attributes": {
      "cluster": {
        "description": "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
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
      "service": {
        "description": "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "task_set_id": {
        "description": "The ID or full Amazon Resource Name (ARN) of the task set.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "A pseudo-resource that manages which of your ECS task sets is primary.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcsPrimaryTaskSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsPrimaryTaskSet), &result)
	return &result
}
