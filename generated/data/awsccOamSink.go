package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOamSink = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon resource name (ARN) of the ObservabilityAccessManager Sink",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the ObservabilityAccessManager Sink.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description": "The policy of this ObservabilityAccessManager Sink.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to apply to the sink",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Oam::Sink",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOamSinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOamSink), &result)
	return &result
}
