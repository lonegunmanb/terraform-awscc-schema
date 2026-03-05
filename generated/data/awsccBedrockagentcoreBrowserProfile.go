package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreBrowserProfile = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "Timestamp when the browser profile was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the browser profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_saved_at": {
        "computed": true,
        "description": "Timestamp when the browser profile was last saved.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_saved_browser_id": {
        "computed": true,
        "description": "ID of the last saved browser.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_saved_browser_session_id": {
        "computed": true,
        "description": "ID of the last saved browser session.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "Timestamp when the browser profile was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the browser profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_arn": {
        "computed": true,
        "description": "The ARN of a BrowserProfile resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_id": {
        "computed": true,
        "description": "The id of the browser profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Status of browser profile",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of tag keys and values.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::BedrockAgentCore::BrowserProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcoreBrowserProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreBrowserProfile), &result)
	return &result
}
