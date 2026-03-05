package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectContactFlowModuleVersion = `{
  "block": {
    "attributes": {
      "contact_flow_module_id": {
        "computed": true,
        "description": "The identifier of the contact flow module (ARN) this version is tied to.",
        "description_kind": "plain",
        "type": "string"
      },
      "contact_flow_module_version_arn": {
        "computed": true,
        "description": "The identifier of the contact flow module version (ARN).",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the version.",
        "description_kind": "plain",
        "type": "string"
      },
      "flow_module_content_sha_256": {
        "computed": true,
        "description": "Indicates the checksum value of the latest published flow module content",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version number of this revision",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::Connect::ContactFlowModuleVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectContactFlowModuleVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectContactFlowModuleVersion), &result)
	return &result
}
