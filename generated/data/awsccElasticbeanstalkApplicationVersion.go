package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticbeanstalkApplicationVersion = `{
  "block": {
    "attributes": {
      "application_name": {
        "computed": true,
        "description": "The name of the Elastic Beanstalk application that is associated with this application version. ",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of this application version.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_bundle": {
        "computed": true,
        "description": "The Amazon S3 bucket and key that identify the location of the source bundle for this version. ",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket": {
              "computed": true,
              "description": "The Amazon S3 bucket where the data is located.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description": "The Amazon S3 key where the data is located.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::ElasticBeanstalk::ApplicationVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticbeanstalkApplicationVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticbeanstalkApplicationVersion), &result)
	return &result
}
