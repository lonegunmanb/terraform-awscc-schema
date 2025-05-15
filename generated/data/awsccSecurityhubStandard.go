package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubStandard = `{
  "block": {
    "attributes": {
      "disabled_standards_controls": {
        "computed": true,
        "description": "Specifies which controls are to be disabled in a standard. \n *Maximum*: ` + "`" + `` + "`" + `100` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "reason": {
              "computed": true,
              "description": "A user-defined reason for changing a control's enablement status in a specified standard. If you are disabling a control, then this property is required.",
              "description_kind": "plain",
              "type": "string"
            },
            "standards_control_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the control.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "standards_arn": {
        "computed": true,
        "description": "The ARN of the standard that you want to enable. To view a list of available ASH standards and their ARNs, use the [DescribeStandards](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "standards_subscription_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityHub::Standard",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityhubStandardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubStandard), &result)
	return &result
}
