package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubHub = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "An ARN is automatically created for the customer.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_enable_controls": {
        "computed": true,
        "description": "Whether to automatically enable new controls when they are added to standards that are enabled",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "control_finding_generator": {
        "computed": true,
        "description": "This field, used when enabling Security Hub, specifies whether the calling account has consolidated control findings turned on. If the value for this field is set to SECURITY_CONTROL, Security Hub generates a single finding for a control check even when the check applies to multiple enabled standards.  If the value for this field is set to STANDARD_CONTROL, Security Hub generates separate findings for a control check when the check applies to multiple enabled standards.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_default_standards": {
        "computed": true,
        "description": "Whether to enable the security standards that Security Hub has designated as automatically enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscribed_at": {
        "computed": true,
        "description": "The date and time when Security Hub was enabled in the account.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "The AWS::SecurityHub::Hub resource represents the implementation of the AWS Security Hub service in your account. One hub resource is created for each Region in which you enable Security Hub.\n\n",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubHubSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubHub), &result)
	return &result
}
