package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLocationTrackerConsumer = `{
  "block": {
    "attributes": {
      "consumer_arn": {
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
      "tracker_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::Location::TrackerConsumer Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLocationTrackerConsumerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLocationTrackerConsumer), &result)
	return &result
}
