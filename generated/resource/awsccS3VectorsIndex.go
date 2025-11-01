package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3VectorsIndex = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "Date and time when the vector index was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_type": {
        "description": "The data type of the vectors to be inserted into the vector index.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dimension": {
        "description": "The dimensions of the vectors to be inserted into the vector index.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "distance_metric": {
        "description": "The distance metric to be used for similarity search.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "index_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the index",
        "description_kind": "plain",
        "type": "string"
      },
      "index_name": {
        "computed": true,
        "description": "The name of the vector index to create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metadata_configuration": {
        "computed": true,
        "description": "The metadata configuration for the vector index.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "non_filterable_metadata_keys": {
              "computed": true,
              "description": "Non-filterable metadata keys allow you to enrich vectors with additional context during storage and retrieval. Unlike default metadata keys, these keys cannot be used as query filters. Non-filterable metadata keys can be retrieved but cannot be searched, queried, or filtered. You can access non-filterable metadata keys of your vectors after finding the vectors.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "vector_bucket_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the vector bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vector_bucket_name": {
        "computed": true,
        "description": "The name of the vector bucket that contains the vector index.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::S3Vectors::Index",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3VectorsIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3VectorsIndex), &result)
	return &result
}
