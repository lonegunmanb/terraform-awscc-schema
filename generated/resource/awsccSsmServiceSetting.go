package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmServiceSetting = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the service setting.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_date": {
        "computed": true,
        "description": "The last time the service setting was modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_user": {
        "computed": true,
        "description": "The ARN of the last modified user.",
        "description_kind": "plain",
        "type": "string"
      },
      "setting_id": {
        "description": "The ID of the service setting, such as /ssm/parameter-store/high-throughput-enabled.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "setting_value": {
        "description": "The value of the service setting.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the service setting. The value can be Default, Customized or PendingUpdate.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SSM::ServiceSetting. ServiceSetting is an account-level setting for an AWS service that defines how a user interacts with or uses a service or feature.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsmServiceSettingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmServiceSetting), &result)
	return &result
}
