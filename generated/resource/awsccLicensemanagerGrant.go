package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLicensemanagerGrant = `{
  "block": {
    "attributes": {
      "allowed_operations": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "grant_arn": {
        "computed": true,
        "description": "Arn of the grant.",
        "description_kind": "plain",
        "type": "string"
      },
      "grant_name": {
        "computed": true,
        "description": "Name for the created Grant.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "home_region": {
        "computed": true,
        "description": "Home region for the created grant.",
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
      "license_arn": {
        "computed": true,
        "description": "License Arn for the grant.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "principals": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version of the grant.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "An example resource schema demonstrating some basic constructs and validation rules.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLicensemanagerGrantSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLicensemanagerGrant), &result)
	return &result
}
