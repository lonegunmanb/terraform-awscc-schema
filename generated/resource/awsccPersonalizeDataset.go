package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPersonalizeDataset = `{
  "block": {
    "attributes": {
      "dataset_arn": {
        "computed": true,
        "description": "The ARN of the dataset",
        "description_kind": "plain",
        "type": "string"
      },
      "dataset_group_arn": {
        "description": "The Amazon Resource Name (ARN) of the dataset group to add the dataset to",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dataset_import_job": {
        "computed": true,
        "description": "Initial DatasetImportJob for the created dataset",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_source": {
              "computed": true,
              "description": "The Amazon S3 bucket that contains the training data to import.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_location": {
                    "computed": true,
                    "description": "The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "dataset_arn": {
              "computed": true,
              "description": "The ARN of the dataset that receives the imported data",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dataset_import_job_arn": {
              "computed": true,
              "description": "The ARN of the dataset import job",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "job_name": {
              "computed": true,
              "description": "The name for the dataset import job.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The ARN of the IAM role that has permissions to read from the Amazon S3 data source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "dataset_type": {
        "description": "The type of dataset",
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
      "name": {
        "description": "The name for the dataset",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema_arn": {
        "description": "The ARN of the schema to associate with the dataset. The schema defines the dataset fields.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Personalize::Dataset.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPersonalizeDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPersonalizeDataset), &result)
	return &result
}
