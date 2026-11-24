package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerExtension = `{
  "block": {
    "attributes": {
      "cluster_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon SageMaker HyperPod cluster to install the extension on.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_values": {
        "computed": true,
        "description": "The configuration settings for the extension, as a JSON string. The settings must match the configuration schema for the extension version. For each setting that you don't specify, SageMaker HyperPod uses its default value. When you update the stack, SageMaker HyperPod applies the settings in the string to the current configuration. Settings that you leave out keep their current values. To restore the default value of a setting, set it to null. To write the settings as YAML or JSON in your template, use the Fn::ToJsonString intrinsic function with the AWS::LanguageExtensions transform. The JSON string can be up to 10,240 characters long.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The date and time that you created the extension.",
        "description_kind": "plain",
        "type": "string"
      },
      "extension_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the extension.",
        "description_kind": "plain",
        "type": "string"
      },
      "extension_name": {
        "computed": true,
        "description": "The name of the extension to install, such as sagemaker-observability. You can't install the same extension twice on a cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "extension_version": {
        "computed": true,
        "description": "The version of the extension to install, such as 1.0.0. When you create the extension without a version, SageMaker HyperPod installs the latest generally available version. If the extension doesn't have a generally available version, you must specify one. If you remove this property from your template during a stack update, the installed version doesn't change. You can't downgrade an extension. To install an earlier version, delete the extension and create it again.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The date and time of the most recent change to the extension.",
        "description_kind": "plain",
        "type": "string"
      },
      "node_summary": {
        "computed": true,
        "description": "The number of cluster nodes that the extension applies to, and how many of them run the requested extension version.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "active_count": {
              "computed": true,
              "description": "The number of cluster nodes that run the requested extension version.",
              "description_kind": "plain",
              "type": "number"
            },
            "total_count": {
              "computed": true,
              "description": "The number of cluster nodes that the extension applies to.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "The current status of the extension. CloudFormation doesn't wait for the extension to finish installing or updating on the cluster nodes. To check progress, use Status and NodeSummary.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to add to the extension. You can add up to 50 tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key. Each key must be unique for the extension.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::Extension",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerExtensionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerExtension), &result)
	return &result
}
