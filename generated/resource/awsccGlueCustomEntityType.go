package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueCustomEntityType = `{
  "block": {
    "attributes": {
      "context_words": {
        "computed": true,
        "description": "A list of context words.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the custom entity type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "regex_string": {
        "computed": true,
        "description": "A regular expression string that is used for detecting sensitive data in a custom pattern.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to associate with the custom entity type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Glue::CustomEntityType",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueCustomEntityTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueCustomEntityType), &result)
	return &result
}
