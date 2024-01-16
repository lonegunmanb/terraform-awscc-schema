package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCustomerprofilesDomain = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time of this integration got created",
        "description_kind": "plain",
        "type": "string"
      },
      "dead_letter_queue_url": {
        "computed": true,
        "description": "The URL of the SQS dead letter queue",
        "description_kind": "plain",
        "type": "string"
      },
      "default_encryption_key": {
        "computed": true,
        "description": "The default encryption key",
        "description_kind": "plain",
        "type": "string"
      },
      "default_expiration_days": {
        "computed": true,
        "description": "The default number of days until the data within the domain expires.",
        "description_kind": "plain",
        "type": "number"
      },
      "domain_name": {
        "computed": true,
        "description": "The unique name of the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The time of this integration got last updated at",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags (keys and values) associated with the domain",
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
    "description": "Data Source schema for AWS::CustomerProfiles::Domain",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCustomerprofilesDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCustomerprofilesDomain), &result)
	return &result
}
