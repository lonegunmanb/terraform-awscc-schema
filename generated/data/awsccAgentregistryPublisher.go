package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAgentregistryPublisher = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp when the Publisher was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "publisher_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Publisher. The Publisher is a global resource, so the ARN has no Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The verification status of the Publisher.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the Publisher.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the Publisher was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "value": {
        "computed": true,
        "description": "The identifier that this Publisher proves ownership of.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "domain": {
              "computed": true,
              "description": "The fully qualified domain name (RFC 1035/1123) with at least two labels, in lowercase. The service stores domain names in lowercase, so specify the value in lowercase to avoid drift.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::AgentRegistry::Publisher",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAgentregistryPublisherSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAgentregistryPublisher), &result)
	return &result
}
