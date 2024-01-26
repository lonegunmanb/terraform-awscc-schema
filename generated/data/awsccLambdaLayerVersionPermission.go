package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaLayerVersionPermission = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description": "The API action that grants access to the layer.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "layer_version_arn": {
        "computed": true,
        "description": "The name or Amazon Resource Name (ARN) of the layer.",
        "description_kind": "plain",
        "type": "string"
      },
      "organization_id": {
        "computed": true,
        "description": "With the principal set to *, grant permission to all accounts in the specified organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal": {
        "computed": true,
        "description": "An account ID, or * to grant layer usage permission to all accounts in an organization, or all AWS accounts (if organizationId is not specified).",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Lambda::LayerVersionPermission",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLambdaLayerVersionPermissionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaLayerVersionPermission), &result)
	return &result
}