package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccStoragegatewayNfsFileShare = `{
  "block": {
    "attributes": {
      "audit_destination_arn": {
        "computed": true,
        "description": "The ARN of the CloudWatch log group that receives the share's audit logs. Removing this property from a template detaches audit logging: the update sends the empty string, which the service treats as the reset.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_region": {
        "computed": true,
        "description": "The Region of the S3 bucket, required when the bucket is reached through a VPC endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_attributes": {
        "computed": true,
        "description": "Refresh cache information for the file share. Removing this property from a template leaves the share's current value: the service documents no value that resets it.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cache_stale_timeout_in_seconds": {
              "computed": true,
              "description": "How long the directory contents cached by the gateway are considered valid.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "client_list": {
        "computed": true,
        "description": "The IPv4 or IPv6 addresses and CIDR blocks permitted to mount the NFS share. Removing this property from a template leaves the share's current value: the service documents no value that resets it.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "default_storage_class": {
        "computed": true,
        "description": "The S3 storage class for new objects. The service accepts S3_STANDARD, S3_INTELLIGENT_TIERING, S3_STANDARD_IA and S3_ONEZONE_IA; it is a length-bounded string in the API model rather than an enumeration.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_type": {
        "computed": true,
        "description": "The server-side encryption the gateway applies to objects it writes.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_share_arn": {
        "computed": true,
        "description": "The ARN of the file share.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_share_id": {
        "computed": true,
        "description": "The ID of the file share.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_share_name": {
        "computed": true,
        "description": "The name of the file share, which defaults to the bucket name when omitted. Removing this property from a template leaves the share's current value: the service documents no value that resets it.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_share_status": {
        "computed": true,
        "description": "The status of the file share. The service reports CREATING, UPDATING, AVAILABLE, DELETING and FORCE_DELETING; it is a length-bounded string in the API model rather than an enumeration.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_arn": {
        "computed": true,
        "description": "The ARN of the S3 File Gateway that hosts the share.",
        "description_kind": "plain",
        "type": "string"
      },
      "guess_mime_type_enabled": {
        "computed": true,
        "description": "Whether the gateway infers the MIME type of uploaded objects from their extension.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_encrypted": {
        "computed": true,
        "description": "Deprecated by the service in favour of EncryptionType; true selects KMS encryption. Retained because it remains readable and settable.",
        "description_kind": "plain",
        "type": "bool"
      },
      "kms_key": {
        "computed": true,
        "description": "The KMS key used for encryption, when EncryptionType selects KMS. Removing this property from a template leaves the share's current value: the service documents no value that resets it.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_arn": {
        "computed": true,
        "description": "The ARN of the S3 bucket, or bucket prefix, the share exposes.",
        "description_kind": "plain",
        "type": "string"
      },
      "nfs_file_share_defaults": {
        "computed": true,
        "description": "Default values for the file and directory permissions of an NFS share. Removing this property from a template leaves the live values unchanged rather than clearing them, because the service has no reset for it: a create that omits the object leaves it absent, but the update path merges member by member keyed on the same internal sentinel that reports absence, so no request can restore that state. Set the members you want rather than relying on removal.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "directory_mode": {
              "computed": true,
              "description": "The Unix directory mode in the form nnnn.",
              "description_kind": "plain",
              "type": "string"
            },
            "file_mode": {
              "computed": true,
              "description": "The Unix file mode in the form nnnn.",
              "description_kind": "plain",
              "type": "string"
            },
            "group_id": {
              "computed": true,
              "description": "The default group ID for files in the file share.",
              "description_kind": "plain",
              "type": "number"
            },
            "owner_id": {
              "computed": true,
              "description": "The default owner ID for files in the file share.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "notification_policy": {
        "computed": true,
        "description": "The upload-notification settling period, as a JSON document.",
        "description_kind": "plain",
        "type": "string"
      },
      "object_acl": {
        "computed": true,
        "description": "The canned S3 ACL applied to objects the gateway writes.",
        "description_kind": "plain",
        "type": "string"
      },
      "path": {
        "computed": true,
        "description": "The export path the share is mounted or mapped at.",
        "description_kind": "plain",
        "type": "string"
      },
      "read_only": {
        "computed": true,
        "description": "Whether clients may write to the share.",
        "description_kind": "plain",
        "type": "bool"
      },
      "requester_pays": {
        "computed": true,
        "description": "Whether the bucket is a Requester Pays bucket.",
        "description_kind": "plain",
        "type": "bool"
      },
      "role": {
        "computed": true,
        "description": "The ARN of the IAM role the gateway assumes when it accesses the S3 location.",
        "description_kind": "plain",
        "type": "string"
      },
      "squash": {
        "computed": true,
        "description": "How the NFS share maps client user IDs. The service accepts RootSquash, NoSquash and AllSquash; it is a length-bounded string in the API model rather than an enumeration.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags assigned to the file share.",
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
          "nesting_mode": "set"
        }
      },
      "vpc_endpoint_dns_name": {
        "computed": true,
        "description": "The DNS name of the S3 VPC endpoint the gateway uses to reach the bucket.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::StorageGateway::NFSFileShare",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccStoragegatewayNfsFileShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccStoragegatewayNfsFileShare), &result)
	return &result
}
