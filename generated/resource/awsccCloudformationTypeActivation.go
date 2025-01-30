package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationTypeActivation = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the extension.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_update": {
        "computed": true,
        "description": "Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging_config": {
        "computed": true,
        "description": "Specifies logging configuration information for a type.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_group_name": {
              "computed": true,
              "description": "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "log_role_arn": {
              "computed": true,
              "description": "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "major_version": {
        "computed": true,
        "description": "The Major Version of the type you want to enable",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_type_arn": {
        "computed": true,
        "description": "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publisher_id": {
        "computed": true,
        "description": "The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The kind of extension",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type_name": {
        "computed": true,
        "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type_name_alias": {
        "computed": true,
        "description": "An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_bump": {
        "computed": true,
        "description": "Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Enable a resource that has been published in the CloudFormation Registry.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudformationTypeActivationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationTypeActivation), &result)
	return &result
}
