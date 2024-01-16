package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMacieSession = `{
  "block": {
    "attributes": {
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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
        "optional": true,
        "type": "string"
      }
    },
    "description": "The AWS::Macie::Session resource specifies a new Amazon Macie session. A session is an object that represents the Amazon Macie service. A session is required for Amazon Macie to become operational.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMacieSessionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMacieSession), &result)
	return &result
}
