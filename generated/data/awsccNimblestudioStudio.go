package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNimblestudioStudio = `{
  "block": {
    "attributes": {
      "admin_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "home_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sso_client_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "studio_encryption_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "studio_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "studio_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "studio_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "user_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NimbleStudio::Studio",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNimblestudioStudioSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNimblestudioStudio), &result)
	return &result
}
