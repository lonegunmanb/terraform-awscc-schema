package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsoInstance = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_store_id": {
        "computed": true,
        "description": "The ID of the identity store associated with the created Identity Center (SSO) Instance",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The SSO Instance ARN that is returned upon creation of the Identity Center (SSO) Instance",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name you want to assign to this Identity Center (SSO) Instance",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description": "The AWS accountId of the owner of the Identity Center (SSO) Instance",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the Identity Center (SSO) Instance, create_in_progress/delete_in_progress/active",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SSO::Instance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsoInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsoInstance), &result)
	return &result
}
