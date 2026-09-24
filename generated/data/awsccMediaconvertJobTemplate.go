package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconvertJobTemplate = `{
  "block": {
    "attributes": {
      "acceleration_settings": {
        "computed": true,
        "description": "Accelerated transcoding can significantly speed up jobs with long, visually complex content. Outputs that use this feature incur pro-tier pricing.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mode": {
              "computed": true,
              "description": "Specify the conditions when the service will run your job with accelerated transcoding.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the job template, such as arn:aws:mediaconvert:us-west-2:123456789012:jobTemplates/my-job-template.",
        "description_kind": "plain",
        "type": "string"
      },
      "category": {
        "computed": true,
        "description": "Optional. A category for the job template you are creating.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Optional. A description of the job template you are creating.",
        "description_kind": "plain",
        "type": "string"
      },
      "hop_destinations": {
        "computed": true,
        "description": "Optional. Configuration for a destination queue to which the job can hop once a customer-defined minimum wait time has passed.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "priority": {
              "computed": true,
              "description": "Optional. A different relative priority for the job in the destination queue.",
              "description_kind": "plain",
              "type": "number"
            },
            "queue": {
              "computed": true,
              "description": "Optional. The destination queue for queue hopping.",
              "description_kind": "plain",
              "type": "string"
            },
            "wait_minutes": {
              "computed": true,
              "description": "Required for queue hopping. Minimum wait time in minutes until the job can hop to the destination queue.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the job template you are creating.",
        "description_kind": "plain",
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description": "Specify the relative priority for this job. In any given queue, the service begins processing the job with the highest value first. When more than one job has the same priority, the service begins processing the job that you submitted first.",
        "description_kind": "plain",
        "type": "number"
      },
      "queue": {
        "computed": true,
        "description": "Optional. The queue that jobs created from this template are assigned to. Specify the Amazon Resource Name (ARN) of the queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "settings_json": {
        "computed": true,
        "description": "Specify, in JSON format, the transcoding job settings for this job template. This specification must conform to the AWS Elemental MediaConvert job validation.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_update_interval": {
        "computed": true,
        "description": "Specify how often MediaConvert sends STATUS_UPDATE events to Amazon CloudWatch Events. Set the interval, in seconds, between status updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::MediaConvert::JobTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediaconvertJobTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconvertJobTemplate), &result)
	return &result
}
