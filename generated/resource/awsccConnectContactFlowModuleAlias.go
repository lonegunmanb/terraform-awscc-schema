package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectContactFlowModuleAlias = `{
  "block": {
    "attributes": {
      "alias_id": {
        "computed": true,
        "description": "The unique identifier of the alias.",
        "description_kind": "plain",
        "type": "string"
      },
      "contact_flow_module_alias_arn": {
        "computed": true,
        "description": "The identifier of the contact flow module alias (ARN). This is constructed from the ContactFlowModuleArn and AliasId.",
        "description_kind": "plain",
        "type": "string"
      },
      "contact_flow_module_id": {
        "description": "The identifier of the contact flow module (ARN) this alias is tied to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "contact_flow_module_version": {
        "description": "The version number of the contact flow module this alias points to.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "description": {
        "computed": true,
        "description": "The description of the alias.",
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
      "name": {
        "description": "The name of the alias.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for ContactFlowModuleAlias",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectContactFlowModuleAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectContactFlowModuleAlias), &result)
	return &result
}
