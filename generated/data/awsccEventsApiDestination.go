package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventsApiDestination = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The arn of the api destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn_for_policy": {
        "computed": true,
        "description": "The arn of the api destination to be used in IAM policies.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_arn": {
        "computed": true,
        "description": "The arn of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "http_method": {
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
      "invocation_endpoint": {
        "computed": true,
        "description": "Url endpoint to invoke.",
        "description_kind": "plain",
        "type": "string"
      },
      "invocation_rate_limit_per_second": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "Name of the apiDestination.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Events::ApiDestination",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEventsApiDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventsApiDestination), &result)
	return &result
}
