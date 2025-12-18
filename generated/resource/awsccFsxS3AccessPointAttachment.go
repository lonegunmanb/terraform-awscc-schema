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
        "description": "The name of the S3 access point attachment; also used for the name of the S3 access point.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ontap_configuration": {
        "computed": true,
        "description": "The OntapConfiguration of the S3 access point attachment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "file_system_identity": {
              "computed": true,
              "description": "The file system identity used to authorize file access requests made using the S3 access point.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "computed": true,
                    "description": "Specifies the FSx for ONTAP user identity type, accepts either UNIX or WINDOWS.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "unix_user": {
                    "computed": true,
                    "description": "Specifies the properties of the file system UNIX user.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the UNIX user.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "windows_user": {
                    "computed": true,
                    "description": "Specifies the properties of the file system Windows user.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the Windows user.",
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
            "volume_id": {
              "computed": true,
              "description": "The ID of the FSx for ONTAP volume that the S3 access point is attached to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "open_zfs_configuration": {
        "computed": true,
        "description": "The OpenZFSConfiguration of the S3 access point attachment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "file_system_identity": {
              "computed": true,
              "description": "The file system identity used to authorize file access requests made using the S3 access point.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "posix_user": {
                    "computed": true,
                    "description": "Specifies the UID and GIDs of the file system POSIX user.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "gid": {
                          "computed": true,
                          "description": "The GID of the file system user.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "secondary_gids": {
                          "computed": true,
                          "description": "The list of secondary GIDs for the file system user.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "gid": {
                                "computed": true,
                                "description": "The GID of the file system user.",
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
                          "computed": true,
                          "description": "The UID of the file system user.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "type": {
                    "computed": true,
                    "description": "Specifies the FSx for OpenZFS user identity type, accepts only POSIX.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "volume_id": {
              "computed": true,
              "description": "The ID of the FSx for OpenZFS volume that the S3 access point is attached to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "s3_access_point": {
        "computed": true,
        "description": "The S3 access point configuration of the S3 access point attachment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alias": {
              "computed": true,
              "description": "The S3 access point's alias.",
              "description_kind": "plain",
              "type": "string"
            },
            "policy": {
              "computed": true,
              "description": "The S3 access point's policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_arn": {
              "computed": true,
              "description": "The S3 access point's ARN.",
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_configuration": {
              "computed": true,
              "description": "The S3 access point's virtual private cloud (VPC) configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "vpc_id": {
                    "computed": true,
                    "description": "Specifies the virtual private cloud (VPC) for the S3 access point VPC configuration, if one exists.",
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
        "description": "The type of Amazon FSx volume that the S3 access point is attached to.",
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
