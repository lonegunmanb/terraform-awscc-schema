package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticbeanstalkApplicationVersion = `{
  "block": {
    "attributes": {
      "application_name": {
        "description": "The name of the Elastic Beanstalk application that is associated with this application version. ",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "application_version_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of this application version.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_bundle": {
        "description": "The Amazon S3 bucket and key that identify the location of the source bundle for this version. ",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket": {
              "description": "The Amazon S3 bucket where the data is located.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_key": {
              "description": "The Amazon S3 key where the data is located.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Resource Type definition for AWS::ElasticBeanstalk::ApplicationVersion",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccElasticbeanstalkApplicationVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticbeanstalkApplicationVersion), &result)
	return &result
}
