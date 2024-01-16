package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPersonalizeSchema = `{
  "block": {
    "attributes": {
      "domain": {
        "computed": true,
        "description": "The domain of a Domain dataset group.",
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
        "description": "Name for the schema.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema": {
        "description": "A schema in Avro JSON format.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema_arn": {
        "computed": true,
        "description": "Arn for the schema.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Personalize::Schema.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPersonalizeSchemaSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPersonalizeSchema), &result)
	return &result
}
