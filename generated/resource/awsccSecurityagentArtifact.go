package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityagentArtifact = `{
  "block": {
    "attributes": {
      "agent_space_id": {
        "description": "The unique identifier of the agent space to add the artifact to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the artifact.",
        "description_kind": "plain",
        "type": "string"
      },
      "artifact_content": {
        "computed": true,
        "description": "The binary content of the artifact to upload, encoded as a Base64 string.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "artifact_id": {
        "computed": true,
        "description": "The unique identifier assigned to the uploaded artifact.",
        "description_kind": "plain",
        "type": "string"
      },
      "artifact_type": {
        "description": "The file type of the artifact.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "file_name": {
        "description": "The file name of the artifact.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time the artifact was last updated, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Uploads an artifact to an agent space. Artifacts provide additional context for security testing, such as architecture diagrams, API specifications, or configuration files.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityagentArtifactSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityagentArtifact), &result)
	return &result
}
