package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubStandard = `{
  "block": {
    "attributes": {
      "disabled_standards_controls": {
        "computed": true,
        "description": "StandardsControls to disable from this Standard.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "reason": {
              "computed": true,
              "description": "the reason the standard control is disabled",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "standards_control_arn": {
              "description": "the Arn for the standard control.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "standards_arn": {
        "description": "The ARN of the Standard being enabled",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "standards_subscription_arn": {
        "computed": true,
        "description": "The ARN of the StandardsSubscription for the account ID, region, and Standard.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The AWS::SecurityHub::Standard resource represents the implementation of an individual AWS Security Hub Standard in your account. It requires you have SecurityHub enabled before you can enable the Standard.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubStandardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubStandard), &result)
	return &result
}
