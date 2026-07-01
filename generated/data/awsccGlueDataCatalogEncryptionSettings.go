package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueDataCatalogEncryptionSettings = `{
  "block": {
    "attributes": {
      "catalog_id": {
        "computed": true,
        "description": "The ID of the Data Catalog in which the settings are created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_catalog_encryption_settings": {
        "computed": true,
        "description": "Contains configuration information for maintaining Data Catalog security.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connection_password_encryption": {
              "computed": true,
              "description": "When connection password protection is enabled, the Data Catalog uses a customer-provided key to encrypt the password as part of CreateConnection or UpdateConnection and store it in the ENCRYPTED_PASSWORD field in the connection properties. You can enable catalog encryption or only password encryption.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kms_key_id": {
                    "computed": true,
                    "description": "An AWS KMS key that is used to encrypt the connection password.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "return_connection_password_encrypted": {
                    "computed": true,
                    "description": "When the ReturnConnectionPasswordEncrypted flag is set to 'true', passwords remain encrypted in the responses of GetConnection and GetConnections. This encryption takes effect independently from catalog encryption.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "encryption_at_rest": {
              "computed": true,
              "description": "Specifies the encryption-at-rest configuration for the Data Catalog.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_encryption_mode": {
                    "computed": true,
                    "description": "The encryption-at-rest mode for encrypting Data Catalog data.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "catalog_encryption_service_role": {
                    "computed": true,
                    "description": "The role that AWS Glue assumes to encrypt and decrypt the Data Catalog objects on the caller's behalf.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "sse_aws_kms_key_id": {
                    "computed": true,
                    "description": "The ID of the AWS KMS key to use for encryption at rest.",
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
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Glue::DataCatalogEncryptionSettings",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueDataCatalogEncryptionSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueDataCatalogEncryptionSettings), &result)
	return &result
}
