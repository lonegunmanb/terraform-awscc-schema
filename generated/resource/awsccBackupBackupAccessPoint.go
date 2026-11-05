package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupBackupAccessPoint = `{
  "block": {
    "attributes": {
      "access_point_in_time": {
        "computed": true,
        "description": "For continuous (point-in-time) recovery points, the timestamp in ISO 8601 format (for example 2021-11-27T03:30:27Z) that determines the point-in-time view of the backup data. The timestamp must be within the continuous backup's retention period at the time the access point is created. Do not specify this for snapshot recovery points. The service reports it back under AccessPointMetadata.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "access_point_metadata": {
        "computed": true,
        "description": "Metadata the service reports for the backup access point. Once the access point reaches the AVAILABLE status this carries the Amazon S3 access point ARN and alias; for continuous recovery points it also echoes the AccessPointInTime that was supplied at create time.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "access_point_in_time": {
              "computed": true,
              "description": "For a continuous recovery point, the point in time the backup access point exposes, echoed back by the service as supplied at create time. Absent for snapshot recovery points. A plain string rather than a date-time, because it is one entry of the service's string-to-string metadata map.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_access_point_alias": {
              "computed": true,
              "description": "The alias of the Amazon S3 access point that Backup provisioned. Use this with standard Amazon S3 read APIs to access the backup data. Populated once the backup access point reaches the AVAILABLE status.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_access_point_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the Amazon S3 access point that Backup provisioned. Use this with standard Amazon S3 read APIs to access the backup data. Populated once the backup access point reaches the AVAILABLE status.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "access_point_policy": {
        "computed": true,
        "description": "An optional resource-based policy document, in JSON, to apply to the underlying Amazon S3 access point. The policy controls how backup data can be accessed through the access point. If you do not specify a policy, access is governed by the caller's IAM permissions. The serialized policy may not exceed 200000 characters. This value is not returned by the service and cannot be updated in place.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that uniquely identifies the backup access point.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vault_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the backup vault that contains the recovery point.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vault_name": {
        "computed": true,
        "description": "The name of the backup vault that contains the recovery point.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The date and time that the backup access point was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the backup access point. This name is shared with the Amazon S3 access point namespace. It must be unique within your account and Region and cannot conflict with an existing Amazon S3 access point. If omitted, a name is generated.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recovery_point_arn": {
        "description": "The Amazon Resource Name (ARN) of the recovery point for which to create the backup access point. The recovery point must be an Amazon S3 recovery point in the AVAILABLE, STOPPED, or COMPLETED state.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resource that was backed up, such as an Amazon S3 bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The type of Amazon Web Services resource associated with the recovery point. For example, S3 for Amazon Simple Storage Service.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the backup access point.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to assign to the backup access point.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Creates a backup access point for an Amazon S3 recovery point. A backup access point provides on-demand, read-only access to the backup data in a recovery point through an Amazon S3 access point, without initiating a restore. While a backup access point is active for a recovery point, Backup pauses lifecycle transitions and blocks deletion of that recovery point.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBackupBackupAccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupBackupAccessPoint), &result)
	return &result
}
