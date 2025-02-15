package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectContactFlowVersion = `{
  "block": {
    "attributes": {
      "contact_flow_id": {
        "description": "The ARN of the contact flow this version is tied to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "contact_flow_version_arn": {
        "computed": true,
        "description": "The identifier of the contact flow version (ARN).",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the version.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flow_content_sha_256": {
        "computed": true,
        "description": "Indicates the checksum value of the latest published flow content",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version number of this revision",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Resource Type Definition for ContactFlowVersion",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectContactFlowVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectContactFlowVersion), &result)
	return &result
}
