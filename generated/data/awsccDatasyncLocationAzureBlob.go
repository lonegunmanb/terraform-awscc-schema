package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationAzureBlob = `{
  "block": {
    "attributes": {
      "agent_arns": {
        "computed": true,
        "description": "Specifies the Amazon Resource Name (ARN) of the DataSync agent that can connect with your Azure Blob Storage container. If you are setting up an agentless cross-cloud transfer, you do not need to specify a value for this parameter.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "azure_access_tier": {
        "computed": true,
        "description": "Specifies an access tier for the objects you're transferring into your Azure Blob Storage container.",
        "description_kind": "plain",
        "type": "string"
      },
      "azure_blob_authentication_type": {
        "computed": true,
        "description": "The specific authentication type that you want DataSync to use to access your Azure Blob Container.",
        "description_kind": "plain",
        "type": "string"
      },
      "azure_blob_container_url": {
        "computed": true,
        "description": "The URL of the Azure Blob container that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "azure_blob_sas_configuration": {
        "computed": true,
        "description": "Specifies the shared access signature (SAS) that DataSync uses to access your Azure Blob Storage container.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "azure_blob_sas_token": {
              "computed": true,
              "description": "Specifies the shared access signature (SAS) token, which indicates the permissions DataSync needs to access your Azure Blob Storage container.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "azure_blob_type": {
        "computed": true,
        "description": "Specifies a blob type for the objects you're transferring into your Azure Blob Storage container.",
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
        "description": "The Amazon Resource Name (ARN) of the Azure Blob Location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the Azure Blob Location that was described.",
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
      "subdirectory": {
        "computed": true,
        "description": "The subdirectory in the Azure Blob Container that is used to read data from the Azure Blob Source Location.",
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
    "description": "Data Source schema for AWS::DataSync::LocationAzureBlob",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatasyncLocationAzureBlobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationAzureBlob), &result)
	return &result
}
