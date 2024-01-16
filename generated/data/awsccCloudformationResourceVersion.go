package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationResourceVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the type, here the ResourceVersion. This is used to uniquely identify a ResourceVersion resource",
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_default_version": {
        "computed": true,
        "description": "Indicates if this type version is the current default version",
        "description_kind": "plain",
        "type": "bool"
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
              "type": "string"
            },
            "log_role_arn": {
              "computed": true,
              "description": "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "provisioning_type": {
        "computed": true,
        "description": "The provisioning behavior of the type. AWS CloudFormation determines the provisioning type during registration, based on the types of handlers in the schema handler package submitted.",
        "description_kind": "plain",
        "type": "string"
      },
      "schema_handler_package": {
        "computed": true,
        "description": "A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.\n\nFor information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.",
        "description_kind": "plain",
        "type": "string"
      },
      "type_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the type without the versionID.",
        "description_kind": "plain",
        "type": "string"
      },
      "type_name": {
        "computed": true,
        "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_id": {
        "computed": true,
        "description": "The ID of the version of the type represented by this resource instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "visibility": {
        "computed": true,
        "description": "The scope at which the type is visible and usable in CloudFormation operations.\n\nValid values include:\n\nPRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.\n\nPUBLIC: The type is publically visible and usable within any Amazon account.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFormation::ResourceVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudformationResourceVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationResourceVersion), &result)
	return &result
}
