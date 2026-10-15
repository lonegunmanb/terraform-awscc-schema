package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDeployment = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Blue/Green Deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "blue_green_deployment_identifier": {
        "computed": true,
        "description": "The system-generated identifier of the Blue/Green Deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "blue_green_deployment_name": {
        "computed": true,
        "description": "The name of the Blue/Green Deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the Blue/Green Deployment was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source": {
        "computed": true,
        "description": "The ARN of the source production database.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the Blue/Green Deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "switchover_details": {
        "computed": true,
        "description": "The source and target member mappings in the Blue/Green Deployment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_member": {
              "computed": true,
              "description": "The source database member in the Blue/Green Deployment.",
              "description_kind": "plain",
              "type": "string"
            },
            "switchover_detail_status": {
              "computed": true,
              "description": "The status of this source and target member pair.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_member": {
              "computed": true,
              "description": "The target database member in the Blue/Green Deployment.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the Blue/Green Deployment.",
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
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target": {
        "computed": true,
        "description": "The ARN of the target database in the green environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_allocated_storage": {
        "computed": true,
        "description": "The allocated storage size in GiB for the target DB instance.",
        "description_kind": "plain",
        "type": "number"
      },
      "target_db_cluster_parameter_group_name": {
        "computed": true,
        "description": "The DB cluster parameter group for the target Aurora DB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_db_instance_class": {
        "computed": true,
        "description": "The DB instance class for target RDS DB instances.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_db_parameter_group_name": {
        "computed": true,
        "description": "The DB parameter group for the target DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_engine_version": {
        "computed": true,
        "description": "The DB engine version for the target database.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_iops": {
        "computed": true,
        "description": "The Provisioned IOPS value for the target DB instance.",
        "description_kind": "plain",
        "type": "number"
      },
      "target_storage_throughput": {
        "computed": true,
        "description": "The storage throughput value for the target DB instance.",
        "description_kind": "plain",
        "type": "number"
      },
      "target_storage_type": {
        "computed": true,
        "description": "The storage type for the target DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "tasks": {
        "computed": true,
        "description": "The provisioning tasks for the Blue/Green Deployment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "task_name": {
              "computed": true,
              "description": "The name of a Blue/Green Deployment provisioning task.",
              "description_kind": "plain",
              "type": "string"
            },
            "task_status": {
              "computed": true,
              "description": "The status of a Blue/Green Deployment provisioning task.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "upgrade_target_storage_config": {
        "computed": true,
        "description": "Whether to upgrade the storage file system configuration on the target database.",
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::RDS::Deployment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRdsDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDeployment), &result)
	return &result
}
