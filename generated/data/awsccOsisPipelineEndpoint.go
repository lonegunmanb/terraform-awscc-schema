package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOsisPipelineEndpoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the pipeline endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_id": {
        "computed": true,
        "description": "The unique identifier of the pipeline endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ingest_endpoint_url": {
        "computed": true,
        "description": "The URL used to ingest data to the pipeline through the VPC endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "pipeline_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the pipeline to create the endpoint for.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the pipeline endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the VPC where the pipeline endpoint was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_options": {
        "computed": true,
        "description": "Configuration settings for the VPC endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "A list of security group IDs that control network access to the pipeline endpoint.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "A list of subnet IDs where the pipeline endpoint network interfaces are created.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::OSIS::PipelineEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOsisPipelineEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOsisPipelineEndpoint), &result)
	return &result
}
