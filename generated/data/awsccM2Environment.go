package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccM2Environment = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_type": {
        "computed": true,
        "description": "The target platform for the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version of the runtime engine for the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the runtime environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "The unique identifier of the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "high_availability_config": {
        "computed": true,
        "description": "Defines the details of a high availability configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "desired_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description": "The type of instance underlying the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.",
        "description_kind": "plain",
        "type": "string"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "Specifies whether the environment is publicly accessible.",
        "description_kind": "plain",
        "type": "bool"
      },
      "security_group_ids": {
        "computed": true,
        "description": "The list of security groups for the VPC associated with this environment.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "storage_configurations": {
        "computed": true,
        "description": "The storage configurations defined for the runtime environment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "efs": {
              "computed": true,
              "description": "Defines the storage configuration for an Amazon EFS file system.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "file_system_id": {
                    "computed": true,
                    "description": "The file system identifier.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mount_point": {
                    "computed": true,
                    "description": "The mount point for the file system.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "fsx": {
              "computed": true,
              "description": "Defines the storage configuration for an Amazon FSx file system.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "file_system_id": {
                    "computed": true,
                    "description": "The file system identifier.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mount_point": {
                    "computed": true,
                    "description": "The mount point for the file system.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "subnet_ids": {
        "computed": true,
        "description": "The unique identifiers of the subnets assigned to this runtime environment.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "Tags associated to this environment.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::M2::Environment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccM2EnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccM2Environment), &result)
	return &result
}
