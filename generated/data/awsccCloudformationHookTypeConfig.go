package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationHookTypeConfig = `{
  "block": {
    "attributes": {
      "configuration": {
        "computed": true,
        "description": "The configuration data for the extension, in this account and region.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_alias": {
        "computed": true,
        "description": "An alias by which to refer to this extension configuration data.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the configuration data, in this account and region.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the type without version number.",
        "description_kind": "plain",
        "type": "string"
      },
      "type_name": {
        "computed": true,
        "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFormation::HookTypeConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudformationHookTypeConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationHookTypeConfig), &result)
	return &result
}
