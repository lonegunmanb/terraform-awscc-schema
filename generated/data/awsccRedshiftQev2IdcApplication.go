package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftQev2IdcApplication = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "idc_display_name": {
        "computed": true,
        "description": "The display name for the Amazon Redshift Query Editor (QEV2) IAM Identity Center application. It appears in the console.",
        "description_kind": "plain",
        "type": "string"
      },
      "idc_instance_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM Identity Center instance used to create the Amazon Redshift Query Editor (QEV2) managed application.",
        "description_kind": "plain",
        "type": "string"
      },
      "idc_managed_application_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the Amazon Redshift Query Editor (QEV2) IAM Identity Center managed application.",
        "description_kind": "plain",
        "type": "string"
      },
      "idc_onboard_status": {
        "computed": true,
        "description": "The onboarding status for the Amazon Redshift Query Editor (QEV2) IAM Identity Center application.",
        "description_kind": "plain",
        "type": "string"
      },
      "qev_2_idc_application_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the Amazon Redshift Query Editor (QEV2) application that integrates with IAM Identity Center.",
        "description_kind": "plain",
        "type": "string"
      },
      "qev_2_idc_application_name": {
        "computed": true,
        "description": "The name of the Amazon Redshift Query Editor (QEV2) application in IAM Identity Center.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags associated with the application. Tags are key-value pairs that you can use to organize and identify your resources.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key, or name, for the resource tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the resource tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Redshift::QEV2IdcApplication",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRedshiftQev2IdcApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftQev2IdcApplication), &result)
	return &result
}
