package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppconfigConfigurationProfile = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "The application ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "configuration_profile_id": {
        "computed": true,
        "description": "The configuration profile ID",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the configuration profile.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The Amazon Resource Name of the AWS Key Management Service key to encrypt new configuration data versions in the AWS AppConfig hosted configuration store. This attribute is only used for hosted configuration types. To encrypt data managed in other configuration stores, see the documentation for how to specify an AWS KMS key for that particular service.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_identifier": {
        "computed": true,
        "description": "The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location_uri": {
        "description": "A URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "A name for the configuration profile.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "retrieval_role_arn": {
        "computed": true,
        "description": "The ARN of an IAM role with permission to access the configuration at the specified LocationUri.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Metadata to assign to the configuration profile. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key-value string map. The tag key can be up to 128 characters and must not start with aws:.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value can be up to 256 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "type": {
        "computed": true,
        "description": "The type of configurations contained in the profile. When calling this API, enter one of the following values for Type: AWS.AppConfig.FeatureFlags, AWS.Freeform",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "validators": {
        "computed": true,
        "description": "A list of methods for validating the configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content": {
              "computed": true,
              "description": "Either the JSON Schema content or the Amazon Resource Name (ARN) of an Lambda function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "AWS AppConfig supports validators of type JSON_SCHEMA and LAMBDA.",
              "description_kind": "plain",
              "optional": true,
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

func AwsccAppconfigConfigurationProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppconfigConfigurationProfile), &result)
	return &result
}
