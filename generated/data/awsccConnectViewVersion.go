package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectViewVersion = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version of the view.",
        "description_kind": "plain",
        "type": "number"
      },
      "version_description": {
        "computed": true,
        "description": "The description for the view version.",
        "description_kind": "plain",
        "type": "string"
      },
      "view_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the view for which a version is being created.",
        "description_kind": "plain",
        "type": "string"
      },
      "view_content_sha_256": {
        "computed": true,
        "description": "The view content hash to be checked.",
        "description_kind": "plain",
        "type": "string"
      },
      "view_version_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the created view version.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::ViewVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectViewVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectViewVersion), &result)
	return &result
}
