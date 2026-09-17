package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicediscoveryInstance = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_attributes": {
        "description": "A string map that contains information for the service that is specified in ServiceId.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "An identifier that you want to associate with the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_id": {
        "description": "The ID or Amazon Resource Name (ARN) of the service that you want to use for settings for the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ServiceDiscovery::Instance",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicediscoveryInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicediscoveryInstance), &result)
	return &result
}
