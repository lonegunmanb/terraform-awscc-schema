package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityagentArtifact = `{
  "block": {
    "attributes": {
      "agent_space_id": {
        "computed": true,
        "description": "The unique identifier of the agent space to add the artifact to.",
        "description_kind": "plain",
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
        "type": "string"
      },
      "artifact_id": {
        "computed": true,
        "description": "The unique identifier assigned to the uploaded artifact.",
        "description_kind": "plain",
        "type": "string"
      },
      "artifact_type": {
        "computed": true,
        "description": "The file type of the artifact.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_name": {
        "computed": true,
        "description": "The file name of the artifact.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time the artifact was last updated, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityAgent::Artifact",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityagentArtifactSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityagentArtifact), &result)
	return &result
}
