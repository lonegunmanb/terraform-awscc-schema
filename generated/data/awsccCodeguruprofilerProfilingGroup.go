package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodeguruprofilerProfilingGroup = `{
  "block": {
    "attributes": {
      "agent_permissions": {
        "computed": true,
        "description": "The agent permissions attached to this profiling group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "principals": {
              "computed": true,
              "description": "The principals for the agent permissions.",
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
      "anomaly_detection_notification_configuration": {
        "computed": true,
        "description": "Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "channel_id": {
              "computed": true,
              "description": "Unique identifier for each Channel in the notification configuration of a Profiling Group",
              "description_kind": "plain",
              "type": "string"
            },
            "channel_uri": {
              "computed": true,
              "description": "Unique arn of the resource to be used for notifications. We support a valid SNS topic arn as a channel uri.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified profiling group.",
        "description_kind": "plain",
        "type": "string"
      },
      "compute_platform": {
        "computed": true,
        "description": "The compute platform of the profiling group.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profiling_group_name": {
        "computed": true,
        "description": "The name of the profiling group.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with a profiling group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::CodeGuruProfiler::ProfilingGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodeguruprofilerProfilingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodeguruprofilerProfilingGroup), &result)
	return &result
}
