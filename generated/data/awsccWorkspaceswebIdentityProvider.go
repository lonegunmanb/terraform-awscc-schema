package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspaceswebIdentityProvider = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_provider_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_provider_details": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "identity_provider_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_provider_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "portal_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::WorkSpacesWeb::IdentityProvider",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWorkspaceswebIdentityProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspaceswebIdentityProvider), &result)
	return &result
}
