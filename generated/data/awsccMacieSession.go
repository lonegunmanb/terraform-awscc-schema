package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMacieSession = `{
  "block": {
    "attributes": {
      "automated_discovery_status": {
        "computed": true,
        "description": "The status of automated sensitive data discovery for the Macie session.",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description": "AWS account ID of customer",
        "description_kind": "plain",
        "type": "string"
      },
      "finding_publishing_frequency": {
        "computed": true,
        "description": "A enumeration value that specifies how frequently finding updates are published.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_role": {
        "computed": true,
        "description": "Service role used by Macie",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "A enumeration value that specifies the status of the Macie Session.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Macie::Session",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMacieSessionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMacieSession), &result)
	return &result
}
