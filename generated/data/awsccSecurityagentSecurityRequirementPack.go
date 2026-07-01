package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityagentSecurityRequirementPack = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "Description of the pack",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "KMS key for client-side encryption of pack contents",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the security requirement pack",
        "description_kind": "plain",
        "type": "string"
      },
      "pack_id": {
        "computed": true,
        "description": "Unique identifier of the security requirement pack",
        "description_kind": "plain",
        "type": "string"
      },
      "security_requirements": {
        "computed": true,
        "description": "Security requirements within this pack",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "Description of the security requirement",
              "description_kind": "plain",
              "type": "string"
            },
            "domain": {
              "computed": true,
              "description": "Security domain this requirement belongs to",
              "description_kind": "plain",
              "type": "string"
            },
            "evaluation": {
              "computed": true,
              "description": "How to evaluate compliance with this requirement",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "Name of the security requirement",
              "description_kind": "plain",
              "type": "string"
            },
            "remediation": {
              "computed": true,
              "description": "How to remediate non-compliance",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "status": {
        "computed": true,
        "description": "Whether the pack is enabled or disabled",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for the security requirement pack",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SecurityAgent::SecurityRequirementPack",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityagentSecurityRequirementPackSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityagentSecurityRequirementPack), &result)
	return &result
}
