package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationModuleDefaultVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the module version to set as the default version.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "module_name": {
        "computed": true,
        "description": "The name of a module existing in the registry.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_id": {
        "computed": true,
        "description": "The ID of an existing version of the named module to set as the default.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFormation::ModuleDefaultVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudformationModuleDefaultVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationModuleDefaultVersion), &result)
	return &result
}
