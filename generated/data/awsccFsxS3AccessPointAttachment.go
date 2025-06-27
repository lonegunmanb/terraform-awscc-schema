package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFsxS3AccessPointAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The Name of the S3AccessPointAttachment",
        "description_kind": "plain",
        "type": "string"
      },
      "open_zfs_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "file_system_identity": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "posix_user": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "gid": {
                          "computed": true,
                          "description_kind": "plain",
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
                                "type": "number"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "uid": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "volume_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::FSx::S3AccessPointAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFsxS3AccessPointAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFsxS3AccessPointAttachment), &result)
	return &result
}
