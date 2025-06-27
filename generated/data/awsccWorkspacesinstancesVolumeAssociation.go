package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspacesinstancesVolumeAssociation = `{
  "block": {
    "attributes": {
      "device": {
        "computed": true,
        "description": "The device name for the volume attachment",
        "description_kind": "plain",
        "type": "string"
      },
      "disassociate_mode": {
        "computed": true,
        "description": "Mode to use when disassociating the volume",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "volume_id": {
        "computed": true,
        "description": "ID of the volume to attach to the workspace instance",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_instance_id": {
        "computed": true,
        "description": "ID of the workspace instance to associate with the volume",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::WorkspacesInstances::VolumeAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWorkspacesinstancesVolumeAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspacesinstancesVolumeAssociation), &result)
	return &result
}
