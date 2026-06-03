package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccResiliencehubv2System = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp when the system was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the system.",
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
        "description": "The KMS key ID for encrypting system data.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the system.",
        "description_kind": "plain",
        "type": "string"
      },
      "system_arn": {
        "computed": true,
        "description": "The ARN of the system.",
        "description_kind": "plain",
        "type": "string"
      },
      "system_id": {
        "computed": true,
        "description": "The system ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags assigned to the system.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the system was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ResilienceHubV2::System",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccResiliencehubv2SystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResiliencehubv2System), &result)
	return &result
}
