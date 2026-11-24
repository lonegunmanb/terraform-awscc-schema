package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalDiscoveryEnablement = `{
  "block": {
    "attributes": {
      "discovery_enablement_id": {
        "computed": true,
        "description": "The ARN of the account discovery configuration or organization discovery configuration to run discovery for.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::APIPortal::DiscoveryEnablement",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApiportalDiscoveryEnablementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalDiscoveryEnablement), &result)
	return &result
}
