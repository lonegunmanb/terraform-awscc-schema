package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2NetworkPerformanceMetricSubscription = `{
  "block": {
    "attributes": {
      "destination": {
        "description": "The target Region or Availability Zone for the metric to subscribe to.",
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
      "metric": {
        "description": "The metric type to subscribe to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source": {
        "description": "The starting Region or Availability Zone for metric to subscribe to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "statistic": {
        "description": "The statistic to subscribe to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::NetworkPerformanceMetricSubscription",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2NetworkPerformanceMetricSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2NetworkPerformanceMetricSubscription), &result)
	return &result
}
