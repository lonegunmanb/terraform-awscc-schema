package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsCustomDbEngineVersion = `{
  "block": {
    "attributes": {
      "database_installation_files_s3_bucket_name": {
        "computed": true,
        "description": "The name of an Amazon S3 bucket that contains database installation files for your CEV. For example, a valid bucket name is ` + "`" + `my-custom-installation-files` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_installation_files_s3_prefix": {
        "computed": true,
        "description": "The Amazon S3 directory that contains the database installation files for your CEV. For example, a valid bucket name is ` + "`" + `123456789012/cev1` + "`" + `. If this setting isn't specified, no prefix is assumed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_engine_version_arn": {
        "computed": true,
        "description": "The ARN of the custom engine version.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description of your CEV.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine": {
        "description": "The database engine to use for your custom engine version (CEV). The only supported value is ` + "`" + `custom-oracle-ee` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_version": {
        "description": "The name of your CEV. The name format is 19.customized_string . For example, a valid name is 19.my_cev1. This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of Engine and EngineVersion is unique per customer per Region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "image_id": {
        "computed": true,
        "description": "The identifier of Amazon Machine Image (AMI) used for CEV.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The AWS KMS key identifier for an encrypted CEV. A symmetric KMS key is required for RDS Custom, but optional for Amazon RDS.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "manifest": {
        "computed": true,
        "description": "The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3. Specify the name/value pairs in a file or a quoted string. RDS Custom applies the patches in the order in which they are listed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_custom_db_engine_version_identifier": {
        "computed": true,
        "description": "The identifier of the source custom engine version.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The availability status to be assigned to the CEV.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "use_aws_provided_latest_image": {
        "computed": true,
        "description": "A value that indicates whether AWS provided latest image is applied automatically to the Custom Engine Version. By default, AWS provided latest image is applied automatically. This value is only applied on create.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "The AWS::RDS::CustomDBEngineVersion resource creates an Amazon RDS custom DB engine version.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsCustomDbEngineVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsCustomDbEngineVersion), &result)
	return &result
}
