package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneFormType = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp of when this Amazon DataZone metadata form type was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The user who created this Amazon DataZone metadata form type.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of this Amazon DataZone metadata form type.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone domain in which this metadata form type is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description": "The ID of the Amazon DataZone domain in which this metadata form type is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "form_type_identifier": {
        "computed": true,
        "description": "The ID of this Amazon DataZone metadata form type.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model": {
        "computed": true,
        "description": "The model of this Amazon DataZone metadata form type.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "smithy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "The name of this Amazon DataZone metadata form type.",
        "description_kind": "plain",
        "type": "string"
      },
      "owning_project_id": {
        "computed": true,
        "description": "The ID of the project that owns this Amazon DataZone metadata form type.",
        "description_kind": "plain",
        "type": "string"
      },
      "owning_project_identifier": {
        "computed": true,
        "description": "The ID of the Amazon DataZone project that owns this metadata form type.",
        "description_kind": "plain",
        "type": "string"
      },
      "revision": {
        "computed": true,
        "description": "The revision of this Amazon DataZone metadata form type.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of this Amazon DataZone metadata form type.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataZone::FormType",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneFormTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneFormType), &result)
	return &result
}
