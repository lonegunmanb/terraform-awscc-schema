package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApprunnerAutoScalingConfiguration = `{
  "block": {
    "attributes": {
      "auto_scaling_configuration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of this auto scaling configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_scaling_configuration_name": {
        "computed": true,
        "description": "The customer-provided auto scaling configuration name.  When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. The auto scaling configuration name can be used in multiple revisions of a configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_scaling_configuration_revision": {
        "computed": true,
        "description": "The revision of this auto scaling configuration. It's unique among all the active configurations (\"Status\": \"ACTIVE\") that share the same AutoScalingConfigurationName.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "latest": {
        "computed": true,
        "description": "It's set to true for the configuration with the highest Revision among all configurations that share the same AutoScalingConfigurationName. It's set to false otherwise. App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code.",
        "description_kind": "plain",
        "type": "bool"
      },
      "max_concurrency": {
        "computed": true,
        "description": "The maximum number of concurrent requests that an instance processes. If the number of concurrent requests exceeds this limit, App Runner scales the service up to use more instances to process the requests.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_size": {
        "computed": true,
        "description": "The maximum number of instances that an App Runner service scales up to. At most MaxSize instances actively serve traffic for your service.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_size": {
        "computed": true,
        "description": "The minimum number of instances that App Runner provisions for a service. The service always has at least MinSize provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.",
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::AppRunner::AutoScalingConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApprunnerAutoScalingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApprunnerAutoScalingConfiguration), &result)
	return &result
}
