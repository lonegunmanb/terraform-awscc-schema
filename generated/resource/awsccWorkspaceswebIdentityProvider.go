package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspaceswebIdentityProvider = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_provider_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_provider_details": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      },
      "identity_provider_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_provider_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "portal_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::WorkSpacesWeb::IdentityProvider Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWorkspaceswebIdentityProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspaceswebIdentityProvider), &result)
	return &result
}
