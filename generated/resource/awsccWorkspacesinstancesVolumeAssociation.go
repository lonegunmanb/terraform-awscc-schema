package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspacesinstancesVolumeAssociation = `{
  "block": {
    "attributes": {
      "device": {
        "description": "The device name for the volume attachment",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disassociate_mode": {
        "computed": true,
        "description": "Mode to use when disassociating the volume",
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
      "volume_id": {
        "description": "ID of the volume to attach to the workspace instance",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "workspace_instance_id": {
        "description": "ID of the workspace instance to associate with the volume",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::WorkspacesInstances::VolumeAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWorkspacesinstancesVolumeAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspacesinstancesVolumeAssociation), &result)
	return &result
}
