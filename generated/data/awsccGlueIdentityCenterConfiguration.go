package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueIdentityCenterConfiguration = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The identifier for the specified AWS account.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_arn": {
        "computed": true,
        "description": "The Glue IAM identity center application arn",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The IAM identity center instance arn",
        "description_kind": "plain",
        "type": "string"
      },
      "scopes": {
        "computed": true,
        "description": "The downstream scopes that Glue identity center configuration can access",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "user_background_sessions_enabled": {
        "computed": true,
        "description": "Enable or disable user background sessions for Glue Identity Center",
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::Glue::IdentityCenterConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueIdentityCenterConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueIdentityCenterConfiguration), &result)
	return &result
}
