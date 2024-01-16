package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFinspaceEnvironment = `{
  "block": {
    "attributes": {
      "aws_account_id": {
        "computed": true,
        "description": "AWS account ID associated with the Environment",
        "description_kind": "plain",
        "type": "string"
      },
      "data_bundles": {
        "computed": true,
        "description": "ARNs of FinSpace Data Bundles to install",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "dedicated_service_account_id": {
        "computed": true,
        "description": "ID for FinSpace created account used to store Environment artifacts",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the Environment",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_arn": {
        "computed": true,
        "description": "ARN of the Environment",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "Unique identifier for representing FinSpace Environment",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_url": {
        "computed": true,
        "description": "URL used to login to the Environment",
        "description_kind": "plain",
        "type": "string"
      },
      "federation_mode": {
        "computed": true,
        "description": "Federation mode used with the Environment",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "federation_parameters": {
        "computed": true,
        "description": "Additional parameters to identify Federation mode",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_call_back_url": {
              "computed": true,
              "description": "SAML metadata URL to link with the Environment",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "attribute_map": {
              "computed": true,
              "description": "Attribute map for SAML configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "federation_provider_name": {
              "computed": true,
              "description": "Federation provider name to link with the Environment",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "federation_urn": {
              "computed": true,
              "description": "SAML metadata URL to link with the Environment",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "saml_metadata_document": {
              "computed": true,
              "description": "SAML metadata document to link the federation provider to the Environment",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "saml_metadata_url": {
              "computed": true,
              "description": "SAML metadata URL to link with the Environment",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
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
        "description": "KMS key used to encrypt customer data within FinSpace Environment infrastructure",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the Environment",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sage_maker_studio_domain_url": {
        "computed": true,
        "description": "SageMaker Studio Domain URL associated with the Environment",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "State of the Environment",
        "description_kind": "plain",
        "type": "string"
      },
      "superuser_parameters": {
        "computed": true,
        "description": "Parameters of the first Superuser for the FinSpace Environment",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "email_address": {
              "computed": true,
              "description": "Email address",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "first_name": {
              "computed": true,
              "description": "First name",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_name": {
              "computed": true,
              "description": "Last name",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "An example resource schema demonstrating some basic constructs and validation rules.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccFinspaceEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFinspaceEnvironment), &result)
	return &result
}
