package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNovaactWorkflowDefinition = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the workflow definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the workflow definition was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description of the workflow definition's purpose and functionality.",
        "description_kind": "plain",
        "type": "string"
      },
      "export_config": {
        "computed": true,
        "description": "Configuration settings for exporting workflow execution data and logs to Amazon S3.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket_name": {
              "computed": true,
              "description": "The name of the Amazon S3 bucket for exporting workflow data.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_key_prefix": {
              "computed": true,
              "description": "An optional prefix for Amazon S3 object keys to organize exported data.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the workflow definition. Must be unique within your account and region.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the workflow definition.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NovaAct::WorkflowDefinition",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNovaactWorkflowDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNovaactWorkflowDefinition), &result)
	return &result
}
