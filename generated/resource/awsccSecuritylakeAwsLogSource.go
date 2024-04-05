package resource

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
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "data_lake_arn": {
        "description": "The ARN for the data lake.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_name": {
        "description": "The name for a AWS source. This must be a Regionally unique value.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_version": {
        "description": "The version for a AWS source. This must be a Regionally unique value.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SecurityLake::AwsLogSource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecuritylakeAwsLogSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecuritylakeAwsLogSource), &result)
	return &result
}
