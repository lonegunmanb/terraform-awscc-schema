package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolReplica = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_pool_tags_at_create": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Cognito::UserPoolReplica",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoUserPoolReplicaSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolReplica), &result)
	return &result
}
