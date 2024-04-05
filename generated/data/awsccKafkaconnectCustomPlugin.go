package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKafkaconnectCustomPlugin = `{
  "block": {
    "attributes": {
      "content_type": {
        "computed": true,
        "description": "The type of the plugin file.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_plugin_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the custom plugin to use.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A summary description of the custom plugin.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_description": {
        "computed": true,
        "description": "Details about the custom plugin file.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "file_md_5": {
              "computed": true,
              "description": "The hex-encoded MD5 checksum of the custom plugin file. You can use it to validate the file.",
              "description_kind": "plain",
              "type": "string"
            },
            "file_size": {
              "computed": true,
              "description": "The size in bytes of the custom plugin file. You can use it to validate the file.",
              "description_kind": "plain",
              "type": "number"
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
      "location": {
        "computed": true,
        "description": "Information about the location of a custom plugin.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_location": {
              "computed": true,
              "description": "The S3 bucket Amazon Resource Name (ARN), file key, and object version of the plugin file stored in Amazon S3.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of an S3 bucket.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "file_key": {
                    "computed": true,
                    "description": "The file key for an object in an S3 bucket.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "object_version": {
                    "computed": true,
                    "description": "The version of an object in an S3 bucket.",
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
      "name": {
        "computed": true,
        "description": "The name of the custom plugin.",
        "description_kind": "plain",
        "type": "string"
      },
      "revision": {
        "computed": true,
        "description": "The revision of the custom plugin.",
        "description_kind": "plain",
        "type": "number"
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::KafkaConnect::CustomPlugin",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKafkaconnectCustomPluginSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKafkaconnectCustomPlugin), &result)
	return &result
}
