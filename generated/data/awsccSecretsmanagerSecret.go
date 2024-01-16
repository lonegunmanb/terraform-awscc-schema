package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecretsmanagerSecret = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "(Optional) Specifies a user-provided description of the secret.",
        "description_kind": "plain",
        "type": "string"
      },
      "generate_secret_string": {
        "computed": true,
        "description": "(Optional) Specifies text data that you want to encrypt and store in this new version of the secret.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "exclude_characters": {
              "computed": true,
              "description": "A string that excludes characters in the generated password. By default, all characters from the included sets can be used. The string can be a minimum length of 0 characters and a maximum length of 7168 characters. ",
              "description_kind": "plain",
              "type": "string"
            },
            "exclude_lowercase": {
              "computed": true,
              "description": "Specifies the generated password should not include lowercase letters. By default, ecrets Manager disables this parameter, and the generated password can include lowercase False, and the generated password can include lowercase letters.",
              "description_kind": "plain",
              "type": "bool"
            },
            "exclude_numbers": {
              "computed": true,
              "description": "Specifies that the generated password should exclude digits. By default, Secrets Manager does not enable the parameter, False, and the generated password can include digits.",
              "description_kind": "plain",
              "type": "bool"
            },
            "exclude_punctuation": {
              "computed": true,
              "description": "Specifies that the generated password should not include punctuation characters. The default if you do not include this switch parameter is that punctuation characters can be included. ",
              "description_kind": "plain",
              "type": "bool"
            },
            "exclude_uppercase": {
              "computed": true,
              "description": "Specifies that the generated password should not include uppercase letters. The default behavior is False, and the generated password can include uppercase letters. ",
              "description_kind": "plain",
              "type": "bool"
            },
            "generate_string_key": {
              "computed": true,
              "description": "The JSON key name used to add the generated password to the JSON structure specified by the SecretStringTemplate parameter. If you specify this parameter, then you must also specify SecretStringTemplate. ",
              "description_kind": "plain",
              "type": "string"
            },
            "include_space": {
              "computed": true,
              "description": "Specifies that the generated password can include the space character. By default, Secrets Manager disables this parameter, and the generated password doesn't include space",
              "description_kind": "plain",
              "type": "bool"
            },
            "password_length": {
              "computed": true,
              "description": "The desired length of the generated password. The default value if you do not include this parameter is 32 characters. ",
              "description_kind": "plain",
              "type": "number"
            },
            "require_each_included_type": {
              "computed": true,
              "description": "Specifies whether the generated password must include at least one of every allowed character type. By default, Secrets Manager enables this parameter, and the generated password includes at least one of every character type.",
              "description_kind": "plain",
              "type": "bool"
            },
            "secret_string_template": {
              "computed": true,
              "description": "A properly structured JSON string that the generated password can be added to. If you specify this parameter, then you must also specify GenerateStringKey.",
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
      "kms_key_id": {
        "computed": true,
        "description": "(Optional) Specifies the ARN, Key ID, or alias of the AWS KMS customer master key (CMK) used to encrypt the SecretString.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The friendly name of the secret. You can use forward slashes in the name to represent a path hierarchy.",
        "description_kind": "plain",
        "type": "string"
      },
      "replica_regions": {
        "computed": true,
        "description": "(Optional) A list of ReplicaRegion objects. The ReplicaRegion type consists of a Region (required) and the KmsKeyId which can be an ARN, Key ID, or Alias.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description": "The ARN, key ID, or alias of the KMS key to encrypt the secret. If you don't include this field, Secrets Manager uses aws/secretsmanager.",
              "description_kind": "plain",
              "type": "string"
            },
            "region": {
              "computed": true,
              "description": "(Optional) A string that represents a Region, for example \"us-east-1\".",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "secret_string": {
        "computed": true,
        "description": "(Optional) Specifies text data that you want to encrypt and store in this new version of the secret.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The list of user-defined tags associated with the secret. Use tags to manage your AWS resources. For additional information about tags, see TagResource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that's 1 to 256 characters in length.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SecretsManager::Secret",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecretsmanagerSecretSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecretsmanagerSecret), &result)
	return &result
}
