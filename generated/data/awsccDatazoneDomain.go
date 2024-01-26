package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneDomain = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the Amazon DataZone domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the Amazon DataZone domain was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the Amazon DataZone domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_execution_role": {
        "computed": true,
        "description": "The domain execution role that is created when an Amazon DataZone domain is created. The domain execution role is created in the AWS account that houses the Amazon DataZone domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_identifier": {
        "computed": true,
        "description": "The identifier of the AWS Key Management Service (KMS) key that is used to encrypt the Amazon DataZone domain, metadata, and reporting data.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The timestamp of when the Amazon DataZone domain was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "managed_account_id": {
        "computed": true,
        "description": "The identifier of the AWS account that manages the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the Amazon DataZone domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_url": {
        "computed": true,
        "description": "The URL of the data portal for this Amazon DataZone domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "single_sign_on": {
        "computed": true,
        "description": "The single-sign on configuration of the Amazon DataZone domain.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description": "The type of single sign-on in Amazon DataZone.",
              "description_kind": "plain",
              "type": "string"
            },
            "user_assignment": {
              "computed": true,
              "description": "The single sign-on user assignment in Amazon DataZone.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "The status of the Amazon DataZone domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags specified for the Amazon DataZone domain.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::DataZone::Domain",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneDomain), &result)
	return &result
}
