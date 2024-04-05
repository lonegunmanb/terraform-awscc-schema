package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecuritylakeAwsLogSource = `{
  "block": {
    "attributes": {
      "accounts": {
        "computed": true,
        "description": "AWS account where you want to collect logs from.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "data_lake_arn": {
        "computed": true,
        "description": "The ARN for the data lake.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_name": {
        "computed": true,
        "description": "The name for a AWS source. This must be a Regionally unique value.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_version": {
        "computed": true,
        "description": "The version for a AWS source. This must be a Regionally unique value.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityLake::AwsLogSource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecuritylakeAwsLogSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecuritylakeAwsLogSource), &result)
	return &result
}
