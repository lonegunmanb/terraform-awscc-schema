package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWickrNetwork = `{
  "block": {
    "attributes": {
      "access_level": {
        "computed": true,
        "description": "The access level of the network, which determines available features and capabilities.",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description": "The AWS account ID that owns the network.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "migration_state": {
        "computed": true,
        "description": "The SSO redirect URI migration state. Values: 0 (not started), 1 (in progress), or 2 (completed).",
        "description_kind": "plain",
        "type": "number"
      },
      "network_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the network.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_id": {
        "computed": true,
        "description": "The unique identifier of the network.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_name": {
        "computed": true,
        "description": "The name of the network. Must be between 1 and 20 characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "standing": {
        "computed": true,
        "description": "The current standing or status of the network.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::Wickr::Network",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWickrNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWickrNetwork), &result)
	return &result
}
