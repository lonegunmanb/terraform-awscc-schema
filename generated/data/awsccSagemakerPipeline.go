package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerPipeline = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parallelism_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_parallel_execution_steps": {
              "computed": true,
              "description": "Maximum parallel execution steps",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "pipeline_definition": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "pipeline_definition_body": {
              "computed": true,
              "description": "A specification that defines the pipeline in JSON format.",
              "description_kind": "plain",
              "type": "string"
            },
            "pipeline_definition_s3_location": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "computed": true,
                    "description": "The name of the S3 bucket where the PipelineDefinition file is stored.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "e_tag": {
                    "computed": true,
                    "description": "The Amazon S3 ETag (a file checksum) of the PipelineDefinition file. If you don't specify a value, SageMaker skips ETag validation of your PipelineDefinition file.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key": {
                    "computed": true,
                    "description": "The file name of the PipelineDefinition file (Amazon S3 object name).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "version": {
                    "computed": true,
                    "description": "For versioning-enabled buckets, a specific version of the PipelineDefinition file.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "pipeline_description": {
        "computed": true,
        "description": "The description of the Pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "pipeline_display_name": {
        "computed": true,
        "description": "The display name of the Pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "pipeline_name": {
        "computed": true,
        "description": "The name of the Pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "Role Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::Pipeline",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerPipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerPipeline), &result)
	return &result
}
