package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubHubV2 = `{
  "block": {
    "attributes": {
      "hub_v2_arn": {
        "computed": true,
        "description": "The Amazon Resource Name of the Security Hub V2 resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscribed_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with the Security Hub V2 resource. You can specify a key that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "The AWS::SecurityHub::HubV2 resource represents the implementation of the AWS Security Hub V2 service in your account. Only one hubv2 resource can created in each region in which you enable Security Hub V2.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubHubV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubHubV2), &result)
	return &result
}
