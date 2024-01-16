package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSignerSigningProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified signing profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "platform_id": {
        "computed": true,
        "description": "The ID of the target signing platform.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_name": {
        "computed": true,
        "description": "A name for the signing profile. AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. ",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_version": {
        "computed": true,
        "description": "A version for the signing profile. AWS Signer generates a unique version for each profile of the same profile name.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_version_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified signing profile version.",
        "description_kind": "plain",
        "type": "string"
      },
      "signature_validity_period": {
        "computed": true,
        "description": "Signature validity period of the profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "A list of tags associated with the signing profile.",
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
    "description": "Data Source schema for AWS::Signer::SigningProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSignerSigningProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSignerSigningProfile), &result)
	return &result
}
