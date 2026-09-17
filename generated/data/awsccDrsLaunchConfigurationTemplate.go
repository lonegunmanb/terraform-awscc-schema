package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDrsLaunchConfigurationTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "ARN of the Launch Configuration Template.",
        "description_kind": "plain",
        "type": "string"
      },
      "copy_private_ip": {
        "computed": true,
        "description": "Copy private IP.",
        "description_kind": "plain",
        "type": "bool"
      },
      "copy_tags": {
        "computed": true,
        "description": "Copy tags.",
        "description_kind": "plain",
        "type": "bool"
      },
      "export_bucket_arn": {
        "computed": true,
        "description": "S3 bucket ARN to export Source Network templates.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "launch_configuration_template_id": {
        "computed": true,
        "description": "ID of the Launch Configuration Template.",
        "description_kind": "plain",
        "type": "string"
      },
      "launch_disposition": {
        "computed": true,
        "description": "Launch disposition.",
        "description_kind": "plain",
        "type": "string"
      },
      "launch_into_source_instance": {
        "computed": true,
        "description": "DRS will set the 'launch into instance ID' of any source server when performing a drill, recovery or failback to the previous region or availability zone, using the instance ID of the source instance.",
        "description_kind": "plain",
        "type": "bool"
      },
      "licensing": {
        "computed": true,
        "description": "Configuration of a machine's license.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "os_byol": {
              "computed": true,
              "description": "Whether to enable Bring your own license or not.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "post_launch_enabled": {
        "computed": true,
        "description": "Whether we want to activate post-launch actions.",
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "A set of tags associated with the Launch Configuration Template.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_instance_type_right_sizing_method": {
        "computed": true,
        "description": "Target instance type right-sizing method.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DRS::LaunchConfigurationTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDrsLaunchConfigurationTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDrsLaunchConfigurationTemplate), &result)
	return &result
}
