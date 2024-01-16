package resource

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
      "connection_arn": {
        "description": "The arn of the connection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http_method": {
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
      "invocation_endpoint": {
        "description": "Url endpoint to invoke.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "invocation_rate_limit_per_second": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "Name of the apiDestination.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Events::ApiDestination.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEventsApiDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventsApiDestination), &result)
	return &result
}
