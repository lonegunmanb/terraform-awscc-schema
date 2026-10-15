package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccStoragegatewaySmbFileShare = `{
  "block": {
    "attributes": {
      "access_based_enumeration": {
        "computed": true,
        "description": "Whether clients see only the files and folders they may access.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "admin_user_list": {
        "computed": true,
        "description": "Users and groups granted administrative privileges on the SMB share.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "audit_destination_arn": {
        "computed": true,
        "description": "The ARN of the CloudWatch log group that receives the share's audit logs. Removing this property from a template detaches audit logging: the update sends the empty string, which the service treats as the reset.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authentication": {
        "computed": true,
        "description": "The authentication the SMB share uses. The service accepts ActiveDirectory and GuestAccess; it is a length-bounded string in the API model rather than an enumeration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bucket_region": {
        "computed": true,
        "description": "The Region of the S3 bucket, required when the bucket is reached through a VPC endpoint.",
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "case_sensitivity": {
        "computed": true,
        "description": "Whether the SMB share is case sensitive or defers to the client.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_storage_class": {
        "computed": true,
        "description": "The S3 storage class for new objects. The service accepts S3_STANDARD, S3_INTELLIGENT_TIERING, S3_STANDARD_IA and S3_ONEZONE_IA; it is a length-bounded string in the API model rather than an enumeration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_type": {
        "computed": true,
        "description": "The server-side encryption the gateway applies to objects it writes.",
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "file_share_status": {
        "computed": true,
        "description": "The status of the file share. The service reports CREATING, UPDATING, AVAILABLE, DELETING and FORCE_DELETING; it is a length-bounded string in the API model rather than an enumeration.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_arn": {
        "description": "The ARN of the S3 File Gateway that hosts the share.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "guess_mime_type_enabled": {
        "computed": true,
        "description": "Whether the gateway infers the MIME type of uploaded objects from their extension.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "invalid_user_list": {
        "computed": true,
        "description": "Users and groups denied access to the SMB share.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "kms_encrypted": {
        "computed": true,
        "description": "Deprecated by the service in favour of EncryptionType; true selects KMS encryption. Retained because it remains readable and settable.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kms_key": {
        "computed": true,
        "description": "The KMS key used for encryption, when EncryptionType selects KMS. Removing this property from a template leaves the share's current value: the service documents no value that resets it.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location_arn": {
        "description": "The ARN of the S3 bucket, or bucket prefix, the share exposes.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_policy": {
        "computed": true,
        "description": "The upload-notification settling period, as a JSON document.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "object_acl": {
        "computed": true,
        "description": "The canned S3 ACL applied to objects the gateway writes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "oplocks_enabled": {
        "computed": true,
        "description": "Whether the SMB share grants opportunistic locks. Removing this property from a template leaves the live value unchanged rather than restoring the created default, because that default depends on the gateway's software version -- opportunistic locking is enabled by default only on appliances at 1.19.0.0 or above -- and the version is not part of this resource. Set the value explicitly rather than relying on removal.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
        "optional": true,
        "type": "bool"
      },
      "requester_pays": {
        "computed": true,
        "description": "Whether the bucket is a Requester Pays bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "role": {
        "description": "The ARN of the IAM role the gateway assumes when it accesses the S3 location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "smbacl_enabled": {
        "computed": true,
        "description": "Whether the SMB share uses Windows ACLs rather than POSIX permissions.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "valid_user_list": {
        "computed": true,
        "description": "Users and groups permitted to access the SMB share.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "vpc_endpoint_dns_name": {
        "computed": true,
        "description": "The DNS name of the S3 VPC endpoint the gateway uses to reach the bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::StorageGateway::SMBFileShare. An SMB file share exposes an Amazon S3 bucket through a Storage Gateway file gateway, so clients read and write objects as files over SMB. The share's variant is fixed by its resource type rather than by a property: the service constructs the two variants with different APIs, and a property carrying the choice would belong to neither API's input.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccStoragegatewaySmbFileShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccStoragegatewaySmbFileShare), &result)
	return &result
}
