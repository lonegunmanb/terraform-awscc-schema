package resource

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
        "optional": true,
        "type": "string"
      },
      "domain_execution_role": {
        "description": "The domain execution role that is created when an Amazon DataZone domain is created. The domain execution role is created in the AWS account that houses the Amazon DataZone domain.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The id of the Amazon DataZone domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_version": {
        "computed": true,
        "description": "The version of the domain.",
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
      "kms_key_identifier": {
        "computed": true,
        "description": "The identifier of the AWS Key Management Service (KMS) key that is used to encrypt the Amazon DataZone domain, metadata, and reporting data.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The name of the Amazon DataZone domain.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "portal_url": {
        "computed": true,
        "description": "The URL of the data portal for this Amazon DataZone domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "root_domain_unit_id": {
        "computed": true,
        "description": "The ID of the root domain in Amazon Datazone.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_role": {
        "computed": true,
        "description": "The service role of the domain that is created.",
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "user_assignment": {
              "computed": true,
              "description": "The single sign-on user assignment in Amazon DataZone.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "A domain is an organizing entity for connecting together assets, users, and their projects",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazoneDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneDomain), &result)
	return &result
}
