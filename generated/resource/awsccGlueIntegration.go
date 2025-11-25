package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueIntegration = `{
  "block": {
    "attributes": {
      "additional_encryption_context": {
        "computed": true,
        "description": "An optional set of non-secret key value pairs that contains additional contextual information about the data.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The time (UTC) when the integration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_filter": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
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
      "integration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_config": {
        "computed": true,
        "description": "The configuration settings for the integration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "continuous_sync": {
              "computed": true,
              "description": "Enables continuous synchronization for on-demand data extractions.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "refresh_interval": {
              "computed": true,
              "description": "Specifies the frequency at which CDC (Change Data Capture) pulls or incremental loads should occur.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_properties": {
              "computed": true,
              "description": "A collection of key-value pairs that specify additional properties for the integration source.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "integration_name": {
        "description": "The name of the integration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "An KMS key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, the default AWS owned KMS key is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_arn": {
        "description": "The Amazon Resource Name (ARN) of the database to use as the source for replication",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the integration.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target_arn": {
        "description": "The Amazon Resource Name (ARN) of the Glue data warehouse to use as the target for replication",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Glue::Integration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueIntegration), &result)
	return &result
}
