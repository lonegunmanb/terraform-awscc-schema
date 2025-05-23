package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationObjectStorage = `{
  "block": {
    "attributes": {
      "access_key": {
        "computed": true,
        "description": "Optional. The access key is used if credentials are required to access the self-managed object storage server.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_arns": {
        "computed": true,
        "description": "Specifies the Amazon Resource Names (ARNs) of the DataSync agents that can connect with your object storage system. If you are setting up an agentless cross-cloud transfer, you do not need to specify a value for this parameter.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "bucket_name": {
        "computed": true,
        "description": "The name of the bucket on the self-managed object storage server.",
        "description_kind": "plain",
        "type": "string"
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
      "location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the object storage location that was described.",
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
      "secret_key": {
        "computed": true,
        "description": "Optional. The secret key is used if credentials are required to access the self-managed object storage server.",
        "description_kind": "plain",
        "type": "string"
      },
      "server_certificate": {
        "computed": true,
        "description": "X.509 PEM content containing a certificate authority or chain to trust.",
        "description_kind": "plain",
        "type": "string"
      },
      "server_hostname": {
        "computed": true,
        "description": "The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.",
        "description_kind": "plain",
        "type": "string"
      },
      "server_port": {
        "computed": true,
        "description": "The port that your self-managed server accepts inbound network traffic on.",
        "description_kind": "plain",
        "type": "number"
      },
      "server_protocol": {
        "computed": true,
        "description": "The protocol that the object storage server uses to communicate.",
        "description_kind": "plain",
        "type": "string"
      },
      "subdirectory": {
        "computed": true,
        "description": "The subdirectory in the self-managed object storage server that is used to read data from.",
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
    "description": "Data Source schema for AWS::DataSync::LocationObjectStorage",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatasyncLocationObjectStorageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationObjectStorage), &result)
	return &result
}
