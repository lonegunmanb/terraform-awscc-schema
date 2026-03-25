package data

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
        "type": "string"
      },
      "account_id": {
        "computed": true,
        "description": "The AWS account ID that this customization belongs to. This is automatically determined from the caller's identity.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "visible_regions": {
        "computed": true,
        "description": "A list of AWS region identifiers visible to the account in the AWS Console.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "visible_services": {
        "computed": true,
        "description": "A list of AWS service identifiers visible to the account in the AWS Console.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::UXC::AccountCustomization",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccUxcAccountCustomizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccUxcAccountCustomization), &result)
	return &result
}
