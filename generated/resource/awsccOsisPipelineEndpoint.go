package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ingest_endpoint_url": {
        "computed": true,
        "description": "The URL used to ingest data to the pipeline through the VPC endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "pipeline_arn": {
        "description": "The Amazon Resource Name (ARN) of the pipeline to create the endpoint for.",
        "description_kind": "plain",
        "required": true,
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
        "description": "Configuration settings for the VPC endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "A list of security group IDs that control network access to the pipeline endpoint.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "A list of subnet IDs where the pipeline endpoint network interfaces are created.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Creates a VPC endpoint for an OpenSearch Ingestion pipeline, enabling private connectivity between your VPC and the pipeline.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOsisPipelineEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOsisPipelineEndpoint), &result)
	return &result
}
