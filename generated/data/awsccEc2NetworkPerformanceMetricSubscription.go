package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2NetworkPerformanceMetricSubscription = `{
  "block": {
    "attributes": {
      "destination": {
        "computed": true,
        "description": "The target Region or Availability Zone for the metric to subscribe to.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metric": {
        "computed": true,
        "description": "The metric type to subscribe to.",
        "description_kind": "plain",
        "type": "string"
      },
      "source": {
        "computed": true,
        "description": "The starting Region or Availability Zone for metric to subscribe to.",
        "description_kind": "plain",
        "type": "string"
      },
      "statistic": {
        "computed": true,
        "description": "The statistic to subscribe to.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::NetworkPerformanceMetricSubscription",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2NetworkPerformanceMetricSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2NetworkPerformanceMetricSubscription), &result)
	return &result
}
