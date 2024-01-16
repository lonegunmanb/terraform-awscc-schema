package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerPipeline = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "parallelism_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_parallel_execution_steps": {
              "description": "Maximum parallel execution steps",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "pipeline_definition": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "pipeline_definition_body": {
              "computed": true,
              "description": "A specification that defines the pipeline in JSON format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pipeline_definition_s3_location": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "description": "The name of the S3 bucket where the PipelineDefinition file is stored.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "e_tag": {
                    "computed": true,
                    "description": "The Amazon S3 ETag (a file checksum) of the PipelineDefinition file. If you don't specify a value, SageMaker skips ETag validation of your PipelineDefinition file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "The file name of the PipelineDefinition file (Amazon S3 object name).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "version": {
                    "computed": true,
                    "description": "For versioning-enabled buckets, a specific version of the PipelineDefinition file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "pipeline_description": {
        "computed": true,
        "description": "The description of the Pipeline.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pipeline_display_name": {
        "computed": true,
        "description": "The display name of the Pipeline.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pipeline_name": {
        "description": "The name of the Pipeline.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "Role Arn",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::Pipeline",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerPipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerPipeline), &result)
	return &result
}
