package resource

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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "max_cpus": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_duration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_gpus": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_runs": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of resource tags",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Definition of AWS::Omics::RunGroup Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOmicsRunGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOmicsRunGroup), &result)
	return &result
}
