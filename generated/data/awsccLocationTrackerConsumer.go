package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLocationTrackerConsumer = `{
  "block": {
    "attributes": {
      "consumer_arn": {
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
      "tracker_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Location::TrackerConsumer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLocationTrackerConsumerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLocationTrackerConsumer), &result)
	return &result
}
