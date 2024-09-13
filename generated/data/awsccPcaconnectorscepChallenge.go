package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectorscepChallenge = `{
  "block": {
    "attributes": {
      "challenge_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connector_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::PCAConnectorSCEP::Challenge",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPcaconnectorscepChallengeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectorscepChallenge), &result)
	return &result
}
