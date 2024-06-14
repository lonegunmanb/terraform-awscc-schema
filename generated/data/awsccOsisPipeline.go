package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOsisPipeline = `{
  "block": {
    "attributes": {
      "buffer_options": {
        "computed": true,
        "description": "Key-value pairs to configure buffering.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "persistent_buffer_enabled": {
              "computed": true,
              "description": "Whether persistent buffering should be enabled.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "encryption_at_rest_options": {
        "computed": true,
        "description": "Key-value pairs to configure encryption at rest.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "The KMS key to use for encrypting data. By default an AWS owned key is used",
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
      "ingest_endpoint_urls": {
        "computed": true,
        "description": "A list of endpoints that can be used for ingesting data into a pipeline",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "log_publishing_options": {
        "computed": true,
        "description": "Key-value pairs to configure log publishing.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_log_destination": {
              "computed": true,
              "description": "The destination for OpenSearch Ingestion Service logs sent to Amazon CloudWatch.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_group": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "is_logging_enabled": {
              "computed": true,
              "description": "Whether logs should be published.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "max_units": {
        "computed": true,
        "description": "The maximum pipeline capacity, in Ingestion OpenSearch Compute Units (OCUs).",
        "description_kind": "plain",
        "type": "number"
      },
      "min_units": {
        "computed": true,
        "description": "The minimum pipeline capacity, in Ingestion OpenSearch Compute Units (OCUs).",
        "description_kind": "plain",
        "type": "number"
      },
      "pipeline_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "pipeline_configuration_body": {
        "computed": true,
        "description": "The Data Prepper pipeline configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "pipeline_name": {
        "computed": true,
        "description": "Name of the OpenSearch Ingestion Service pipeline to create. Pipeline names are unique across the pipelines owned by an account within an AWS Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "vpc_endpoint_service": {
        "computed": true,
        "description": "The VPC endpoint service name for the pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_endpoints": {
        "computed": true,
        "description": "The VPC interface endpoints that have access to the pipeline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vpc_endpoint_id": {
              "computed": true,
              "description": "The unique identifier of the endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_id": {
              "computed": true,
              "description": "The ID for your VPC. AWS Privatelink generates this value when you create a VPC.",
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_options": {
              "computed": true,
              "description": "Container for the values required to configure VPC access for the pipeline. If you don't specify these values, OpenSearch Ingestion Service creates the pipeline with a public endpoint.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_ids": {
                    "computed": true,
                    "description": "A list of security groups associated with the VPC endpoint.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_ids": {
                    "computed": true,
                    "description": "A list of subnet IDs associated with the VPC endpoint.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "vpc_endpoint_management": {
                    "computed": true,
                    "description": "Defines whether you or Amazon OpenSearch Ingestion service create and manage the VPC endpoint configured for the pipeline.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "vpc_options": {
        "computed": true,
        "description": "Container for the values required to configure VPC access for the pipeline. If you don't specify these values, OpenSearch Ingestion Service creates the pipeline with a public endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "A list of security groups associated with the VPC endpoint.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "A list of subnet IDs associated with the VPC endpoint.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "vpc_endpoint_management": {
              "computed": true,
              "description": "Defines whether you or Amazon OpenSearch Ingestion service create and manage the VPC endpoint configured for the pipeline.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::OSIS::Pipeline",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOsisPipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOsisPipeline), &result)
	return &result
}
