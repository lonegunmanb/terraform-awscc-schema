package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53DelegationSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the delegation set.",
        "description_kind": "plain",
        "type": "string"
      },
      "caller_reference": {
        "computed": true,
        "description": "A unique string that identifies the request and allows retrying failed CreateReusableDelegationSet requests without risk of executing the operation twice.",
        "description_kind": "plain",
        "type": "string"
      },
      "delegation_set_id": {
        "computed": true,
        "description": "The ID that Amazon Route 53 assigns to a reusable delegation set.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name_servers": {
        "computed": true,
        "description": "A list of the authoritative name servers for the delegation set.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Route53::DelegationSet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53DelegationSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53DelegationSet), &result)
	return &result
}
