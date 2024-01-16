package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotfleethubApplication = `{
  "block": {
    "attributes": {
      "application_arn": {
        "computed": true,
        "description": "The ARN of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_creation_date": {
        "computed": true,
        "description": "When the Application was created",
        "description_kind": "plain",
        "type": "number"
      },
      "application_description": {
        "computed": true,
        "description": "Application Description, should be between 1 and 2048 characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_id": {
        "computed": true,
        "description": "The ID of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_last_update_date": {
        "computed": true,
        "description": "When the Application was last updated",
        "description_kind": "plain",
        "type": "number"
      },
      "application_name": {
        "computed": true,
        "description": "Application Name, should be between 1 and 256 characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_state": {
        "computed": true,
        "description": "The current state of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_url": {
        "computed": true,
        "description": "The URL of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "error_message": {
        "computed": true,
        "description": "A message indicating why Create or Delete Application failed.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of the role that the web application assumes when it interacts with AWS IoT Core. For more info on configuring this attribute, see https://docs.aws.amazon.com/iot/latest/apireference/API_iotfleethub_CreateApplication.html#API_iotfleethub_CreateApplication_RequestSyntax",
        "description_kind": "plain",
        "type": "string"
      },
      "sso_client_id": {
        "computed": true,
        "description": "The AWS SSO application generated client ID (used with AWS SSO APIs).",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::IoTFleetHub::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotfleethubApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotfleethubApplication), &result)
	return &result
}
