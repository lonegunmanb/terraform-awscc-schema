package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationModuleVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the module.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the registered module.",
        "description_kind": "plain",
        "type": "string"
      },
      "documentation_url": {
        "computed": true,
        "description": "The URL of a page providing detailed documentation for this module.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_default_version": {
        "computed": true,
        "description": "Indicator of whether this module version is the current default version",
        "description_kind": "plain",
        "type": "bool"
      },
      "module_name": {
        "description": "The name of the module being registered.\n\nRecommended module naming pattern: company_or_organization::service::type::MODULE.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "module_package": {
        "description": "The url to the S3 bucket containing the schema and template fragment for the module you want to register.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema": {
        "computed": true,
        "description": "The schema defining input parameters to and resources generated by the module.",
        "description_kind": "plain",
        "type": "string"
      },
      "time_created": {
        "computed": true,
        "description": "The time that the specified module version was registered.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_id": {
        "computed": true,
        "description": "The version ID of the module represented by this module instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "visibility": {
        "computed": true,
        "description": "The scope at which the type is visible and usable in CloudFormation operations.\n\nThe only allowed value at present is:\n\nPRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "A module that has been registered in the CloudFormation registry.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudformationModuleVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationModuleVersion), &result)
	return &result
}
