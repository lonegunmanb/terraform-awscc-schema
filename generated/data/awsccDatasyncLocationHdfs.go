package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationHdfs = `{
  "block": {
    "attributes": {
      "agent_arns": {
        "computed": true,
        "description": "ARN(s) of the agent(s) to use for an HDFS location.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "authentication_type": {
        "computed": true,
        "description": "The authentication mode used to determine identity of user.",
        "description_kind": "plain",
        "type": "string"
      },
      "block_size": {
        "computed": true,
        "description": "Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.",
        "description_kind": "plain",
        "type": "number"
      },
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
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kerberos_keytab": {
        "computed": true,
        "description": "The Base64 string representation of the Keytab file.",
        "description_kind": "plain",
        "type": "string"
      },
      "kerberos_krb_5_conf": {
        "computed": true,
        "description": "The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "kerberos_principal": {
        "computed": true,
        "description": "The unique identity, or principal, to which Kerberos can assign tickets.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_provider_uri": {
        "computed": true,
        "description": "The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the HDFS location.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the HDFS location that was described.",
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
      "name_nodes": {
        "computed": true,
        "description": "An array of Name Node(s) of the HDFS location.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "hostname": {
              "computed": true,
              "description": "The DNS name or IP address of the Name Node in the customer's on premises HDFS cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port on which the Name Node is listening on for client requests.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "qop_configuration": {
        "computed": true,
        "description": "Configuration information for RPC Protection and Data Transfer Protection. These parameters can be set to AUTHENTICATION, INTEGRITY, or PRIVACY. The default value is PRIVACY.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_transfer_protection": {
              "computed": true,
              "description": "Configuration for Data Transfer Protection.",
              "description_kind": "plain",
              "type": "string"
            },
            "rpc_protection": {
              "computed": true,
              "description": "Configuration for RPC Protection.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "replication_factor": {
        "computed": true,
        "description": "Number of copies of each block that exists inside the HDFS cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "simple_user": {
        "computed": true,
        "description": "The user name that has read and write permissions on the specified HDFS cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "subdirectory": {
        "computed": true,
        "description": "The subdirectory in HDFS that is used to read data from the HDFS source location or write data to the HDFS destination.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::DataSync::LocationHDFS",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatasyncLocationHdfsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationHdfs), &result)
	return &result
}
