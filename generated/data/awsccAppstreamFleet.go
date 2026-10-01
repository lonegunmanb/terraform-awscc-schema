package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamFleet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "attributes_to_delete": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "compute_capacity": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "desired_instances": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "desired_sessions": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disable_imdsv1": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "disconnect_timeout_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_join_info": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "directory_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "organizational_unit_distinguished_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "enable_default_internet_access": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "fleet_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "iam_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "idle_disconnect_timeout_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "image_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "max_concurrent_sessions": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_sessions_per_instance": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_user_duration_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "platform": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "root_volume_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "volume_size_in_gb": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "session_script_s3_location": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "stream_view": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      },
      "usb_device_filter_strings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::AppStream::Fleet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppstreamFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamFleet), &result)
	return &result
}
