package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFsxS3AccessPointAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The Name of the S3AccessPointAttachment",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "open_zfs_configuration": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "file_system_identity": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "posix_user": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "gid": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "secondary_gids": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "gid": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "uid": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "volume_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "s3_access_point": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alias": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "policy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "vpc_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::FSx::S3AccessPointAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccFsxS3AccessPointAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFsxS3AccessPointAttachment), &result)
	return &result
}
