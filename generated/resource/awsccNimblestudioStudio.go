package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNimblestudioStudio = `{
  "block": {
    "attributes": {
      "admin_role_arn": {
        "description": "\u003cp\u003eThe IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.\u003c/p\u003e",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "\u003cp\u003eA friendly name for the studio.\u003c/p\u003e",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "home_region": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Web Services Region where the studio resource is located.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "sso_client_id": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Web Services SSO application client ID used to integrate with Amazon Web Services SSO to enable Amazon Web Services SSO users to log in to Nimble Studio portal.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "studio_encryption_configuration": {
        "computed": true,
        "description": "\u003cp\u003eConfiguration of the encryption method that is used for the studio.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key_arn": {
              "computed": true,
              "description": "\u003cp\u003eThe ARN for a KMS key that is used to encrypt studio data.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_type": {
              "computed": true,
              "description": "\u003cp\u003eThe type of KMS key that is used to encrypt studio data.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "studio_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "studio_name": {
        "description": "\u003cp\u003eThe studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.\u003c/p\u003e",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "studio_url": {
        "computed": true,
        "description": "\u003cp\u003eThe address of the web page for the studio.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "user_role_arn": {
        "description": "\u003cp\u003eThe IAM role that Studio Users will assume when logging in to the Nimble Studio portal.\u003c/p\u003e",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Represents a studio that contains other Nimble Studio resources",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNimblestudioStudioSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNimblestudioStudio), &result)
	return &result
}
