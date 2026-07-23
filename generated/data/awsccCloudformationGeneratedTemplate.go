package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationGeneratedTemplate = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time the generated template was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "generated_template_id": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the generated template.",
        "description_kind": "plain",
        "type": "string"
      },
      "generated_template_name": {
        "computed": true,
        "description": "The name assigned to the generated template.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The time the generated template was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "progress": {
        "computed": true,
        "description": "A summary of the progress of the template generation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resources_failed": {
              "computed": true,
              "description": "The number of resources that failed the template generation.",
              "description_kind": "plain",
              "type": "number"
            },
            "resources_pending": {
              "computed": true,
              "description": "The number of resources that are still pending the template generation.",
              "description_kind": "plain",
              "type": "number"
            },
            "resources_processing": {
              "computed": true,
              "description": "The number of resources that are in-process for the template generation.",
              "description_kind": "plain",
              "type": "number"
            },
            "resources_succeeded": {
              "computed": true,
              "description": "The number of resources that succeeded the template generation.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "The status of the template generation.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_configuration": {
        "computed": true,
        "description": "The configuration details of the generated template.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "deletion_policy": {
              "computed": true,
              "description": "The DeletionPolicy assigned to resources in the generated template.",
              "description_kind": "plain",
              "type": "string"
            },
            "update_replace_policy": {
              "computed": true,
              "description": "The UpdateReplacePolicy assigned to resources in the generated template.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "total_warnings": {
        "computed": true,
        "description": "The number of warnings generated for this template.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::CloudFormation::GeneratedTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudformationGeneratedTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationGeneratedTemplate), &result)
	return &result
}
