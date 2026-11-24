package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalDiscoveryEnablement = `{
  "block": {
    "attributes": {
      "discovery_enablement_id": {
        "description": "The ARN of the account discovery configuration or organization discovery configuration to run discovery for.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::APIPortal::DiscoveryEnablement. Starts API discovery for an AWS::APIPortal::DiscoveryConfiguration or AWS::APIPortal::OrganizationDiscoveryConfiguration while it exists, and stops discovery when it is deleted.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApiportalDiscoveryEnablementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalDiscoveryEnablement), &result)
	return &result
}
