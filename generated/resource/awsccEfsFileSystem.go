package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEfsFileSystem = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_policy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "bypass_policy_lockout_safety_check": {
        "computed": true,
        "description": "Whether to bypass the FileSystemPolicy lockout safety check. The policy lockout safety check determines whether the policy in the request will prevent the principal making the request to be locked out from making future PutFileSystemPolicy requests on the file system. Set BypassPolicyLockoutSafetyCheck to True only when you intend to prevent the principal that is making the request from making a subsequent PutFileSystemPolicy request on the file system. Defaults to false",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "encrypted": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "file_system_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "file_system_protection": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "replication_overwrite_protection": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "file_system_tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lifecycle_policies": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "transition_to_archive": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "transition_to_ia": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "transition_to_primary_storage_class": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "performance_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioned_throughput_in_mibps": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "replication_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destinations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_zone_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "file_system_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
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
          "nesting_mode": "single"
        },
        "optional": true
      },
      "throughput_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EFS::FileSystem",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEfsFileSystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEfsFileSystem), &result)
	return &result
}
