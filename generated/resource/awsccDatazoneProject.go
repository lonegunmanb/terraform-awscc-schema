package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneProject = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the project was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The Amazon DataZone user who created the project.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the Amazon DataZone project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone domain in which the project was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "description": "The ID of the Amazon DataZone domain in which this project is created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "glossary_terms": {
        "computed": true,
        "description": "The glossary terms that can be used in this Amazon DataZone project.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone project.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The timestamp of when the project was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the Amazon DataZone project.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Amazon DataZone projects are business use case?based groupings of people, assets (data), and tools used to simplify access to the AWS analytics.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazoneProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneProject), &result)
	return &result
}
