package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationFsxOntap = `{
  "block": {
    "attributes": {
      "fsx_filesystem_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the FSx ONAP file system.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the FSx ONTAP file system that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description": "Configuration settings for NFS or SMB protocol.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "nfs": {
              "computed": true,
              "description": "NFS protocol configuration for FSx ONTAP file system.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "mount_options": {
                    "computed": true,
                    "description": "The NFS mount options that DataSync can use to mount your NFS share.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "version": {
                          "computed": true,
                          "description": "The specific NFS version that you want DataSync to use to mount your NFS share.",
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
            "smb": {
              "computed": true,
              "description": "SMB protocol configuration for FSx ONTAP file system.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cmk_secret_config": {
                    "computed": true,
                    "description": "Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "kms_key_arn": {
                          "computed": true,
                          "description": "Specifies the ARN for the customer-managed AWS KMS key used to encrypt the secret specified for SecretArn. DataSync provides this key to AWS Secrets Manager.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "secret_arn": {
                          "computed": true,
                          "description": "Specifies the ARN for an AWS Secrets Manager secret, managed by DataSync.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "custom_secret_config": {
                    "computed": true,
                    "description": "Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "secret_access_role_arn": {
                          "computed": true,
                          "description": "Specifies the ARN for the AWS Identity and Access Management role that DataSync uses to access the secret specified for SecretArn.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "secret_arn": {
                          "computed": true,
                          "description": "Specifies the ARN for a customer created AWS Secrets Manager secret.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "domain": {
                    "computed": true,
                    "description": "The name of the Windows domain that the SMB server belongs to.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "managed_secret_config": {
                    "computed": true,
                    "description": "Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location. DataSync uses the default AWS-managed KMS key to encrypt this secret in AWS Secrets Manager.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "secret_arn": {
                          "computed": true,
                          "description": "Specifies the ARN for an AWS Secrets Manager secret.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "mount_options": {
                    "computed": true,
                    "description": "The mount options used by DataSync to access the SMB server.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "version": {
                          "computed": true,
                          "description": "The specific SMB version that you want DataSync to use to mount your SMB share.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "password": {
                    "computed": true,
                    "description": "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "user": {
                    "computed": true,
                    "description": "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
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
      "security_group_arns": {
        "computed": true,
        "description": "The ARNs of the security groups that are to use to configure the FSx ONTAP file system.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "storage_virtual_machine_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the FSx ONTAP SVM.",
        "description_kind": "plain",
        "type": "string"
      },
      "subdirectory": {
        "computed": true,
        "description": "A subdirectory in the location's path.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::DataSync::LocationFSxONTAP",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatasyncLocationFsxOntapSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationFsxOntap), &result)
	return &result
}
