package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchDataSource = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_type": {
        "computed": true,
        "description": "The type of data source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_glue_data_catalog": {
              "computed": true,
              "description": "Configuration for an S3 Glue Data Catalog data source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "role_arn": {
                    "computed": true,
                    "description": "The ARN of the IAM role that grants OpenSearch Service permission to access the Glue Data Catalog.",
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
      "description": {
        "computed": true,
        "description": "A description of the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The name of the OpenSearch Service domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the data source.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::OpenSearch::DataSource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpensearchDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchDataSource), &result)
	return &result
}
