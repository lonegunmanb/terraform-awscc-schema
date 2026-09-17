package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMacieMember = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The AWS account ID for the account to associate with the Amazon Macie administrator account.",
        "description_kind": "plain",
        "type": "string"
      },
      "administrator_account_id": {
        "computed": true,
        "description": "The AWS account ID for the Amazon Macie administrator account.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the association between the member account and the Amazon Macie administrator account.",
        "description_kind": "plain",
        "type": "string"
      },
      "email": {
        "computed": true,
        "description": "The email address for the account to associate with the Amazon Macie administrator account. Required by the Amazon Macie CreateMember API at creation time; it is write-only because the service does not return it (it is null when the account is associated through AWS Organizations).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "relationship_status": {
        "computed": true,
        "description": "The current status of the relationship between the account and the Amazon Macie administrator account.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to associate with the member account in Amazon Macie.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag. The maximum length of a tag key is 128 characters.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag. The maximum length of a tag value is 256 characters.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time, in UTC and extended ISO 8601 format, of the most recent change to the status of the relationship between the account and the Amazon Macie administrator account.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Macie::Member",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMacieMemberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMacieMember), &result)
	return &result
}
