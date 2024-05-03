package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockDataSource = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time at which the data source was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_deletion_policy": {
        "computed": true,
        "description": "The deletion policy for the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_configuration": {
        "computed": true,
        "description": "Specifies a raw data source location to ingest.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_configuration": {
              "computed": true,
              "description": "Contains information about the S3 configuration of the data source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description": "The ARN of the bucket that contains the data source.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "bucket_owner_account_id": {
                    "computed": true,
                    "description": "The account ID for the owner of the S3 bucket.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "inclusion_prefixes": {
                    "computed": true,
                    "description": "A list of S3 prefixes that define the object containing the data sources.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "type": {
              "computed": true,
              "description": "The type of the data source location.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "data_source_id": {
        "computed": true,
        "description": "Identifier for a resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_status": {
        "computed": true,
        "description": "The status of a data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the Resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "failure_reasons": {
        "computed": true,
        "description": "The details of the failure reasons related to the data source.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "knowledge_base_id": {
        "computed": true,
        "description": "The unique identifier of the knowledge base to which to add the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "server_side_encryption_configuration": {
        "computed": true,
        "description": "Contains details about the server-side encryption for the data source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "The ARN of the AWS KMS key used to encrypt the resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The time at which the knowledge base was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "vector_ingestion_configuration": {
        "computed": true,
        "description": "Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "chunking_configuration": {
              "computed": true,
              "description": "Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "chunking_strategy": {
                    "computed": true,
                    "description": "Knowledge base can split your source data into chunks. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. You have the following options for chunking your data. If you opt for NONE, then you may want to pre-process your files by splitting them up such that each file corresponds to a chunk.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "fixed_size_chunking_configuration": {
                    "computed": true,
                    "description": "Configurations for when you choose fixed-size chunking. If you set the chunkingStrategy as NONE, exclude this field.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_tokens": {
                          "computed": true,
                          "description": "The maximum number of tokens to include in a chunk.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "overlap_percentage": {
                          "computed": true,
                          "description": "The percentage of overlap between adjacent chunks of a data source.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Bedrock::DataSource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockDataSource), &result)
	return &result
}
