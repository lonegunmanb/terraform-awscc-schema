package resource

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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Represents a SCEP Challenge that is used for certificate enrollment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPcaconnectorscepChallengeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectorscepChallenge), &result)
	return &result
}
