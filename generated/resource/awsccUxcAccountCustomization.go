package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccUxcAccountCustomization = `{
  "block": {
    "attributes": {
      "account_color": {
        "computed": true,
        "description": "The color theme assigned to the account for visual identification in the AWS Console.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "account_id": {
        "computed": true,
        "description": "The AWS account ID that this customization belongs to. This is automatically determined from the caller's identity.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "visible_regions": {
        "computed": true,
        "description": "A list of AWS region identifiers visible to the account in the AWS Console.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "visible_services": {
        "computed": true,
        "description": "A list of AWS service identifiers visible to the account in the AWS Console.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Resource schema for managing AWS account-level UX customization settings, including account color, visible services, and visible regions.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccUxcAccountCustomizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccUxcAccountCustomization), &result)
	return &result
}
