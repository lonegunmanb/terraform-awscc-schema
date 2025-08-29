package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsCustomDbEngineVersion = `{
  "block": {
    "attributes": {
      "database_installation_files_s3_bucket_name": {
        "computed": true,
        "description": "The name of an Amazon S3 bucket that contains database installation files for your CEV. For example, a valid bucket name is ` + "`" + `` + "`" + `my-custom-installation-files` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_installation_files_s3_prefix": {
        "computed": true,
        "description": "The Amazon S3 directory that contains the database installation files for your CEV. For example, a valid bucket name is ` + "`" + `` + "`" + `123456789012/cev1` + "`" + `` + "`" + `. If this setting isn't specified, no prefix is assumed.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_engine_version_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description of your CEV.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine": {
        "computed": true,
        "description": "The database engine to use for your custom engine version (CEV).\n Valid values:\n  +   ` + "`" + `` + "`" + `custom-oracle-ee` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `custom-oracle-ee-cdb` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The name of your CEV. The name format is ` + "`" + `` + "`" + `major version.customized_string` + "`" + `` + "`" + `. For example, a valid CEV name is ` + "`" + `` + "`" + `19.my_cev1` + "`" + `` + "`" + `. This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of ` + "`" + `` + "`" + `Engine` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `EngineVersion` + "`" + `` + "`" + ` is unique per customer per Region.\n *Constraints:* Minimum length is 1. Maximum length is 60.\n *Pattern:*` + "`" + `` + "`" + `^[a-z0-9_.-]{1,60$` + "`" + `` + "`" + `}",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_id": {
        "computed": true,
        "description": "A value that indicates the ID of the AMI.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The AWS KMS key identifier for an encrypted CEV. A symmetric encryption KMS key is required for RDS Custom, but optional for Amazon RDS.\n If you have an existing symmetric encryption KMS key in your account, you can use it with RDS Custom. No further action is necessary. If you don't already have a symmetric encryption KMS key in your account, follow the instructions in [Creating a symmetric encryption KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html#create-symmetric-cmk) in the *Key Management Service Developer Guide*.\n You can choose the same symmetric encryption key when you create a CEV and a DB instance, or choose different keys.",
        "description_kind": "plain",
        "type": "string"
      },
      "manifest": {
        "computed": true,
        "description": "The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3. Specify the name/value pairs in a file or a quoted string. RDS Custom applies the patches in the order in which they are listed.\n The following JSON fields are valid:\n  + MediaImportTemplateVersion Version of the CEV manifest. The date is in the format YYYY-MM-DD. + databaseInstallationFileNames Ordered list of installation files for the CEV. + opatchFileNames Ordered list of OPatch installers used for the Oracle DB engine. + psuRuPatchFileNames The PSU and RU patches for this CEV. + OtherPatchFileNames The patches that are not in the list of PSU and RU patches. Amazon RDS applies these patches after applying the PSU and RU patches. \n For more information, see [Creating the CEV manifest](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.preparing.manifest) in the *Amazon RDS User Guide*.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_custom_db_engine_version_identifier": {
        "computed": true,
        "description": "The ARN of a CEV to use as a source for creating a new CEV. You can specify a different Amazon Machine Imagine (AMI) by using either ` + "`" + `` + "`" + `Source` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `UseAwsProvidedLatestImage` + "`" + `` + "`" + `. You can't specify a different JSON manifest when you specify ` + "`" + `` + "`" + `SourceCustomDbEngineVersionIdentifier` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "A value that indicates the status of a custom engine version (CEV).",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "use_aws_provided_latest_image": {
        "computed": true,
        "description": "Specifies whether to use the latest service-provided Amazon Machine Image (AMI) for the CEV. If you specify ` + "`" + `` + "`" + `UseAwsProvidedLatestImage` + "`" + `` + "`" + `, you can't also specify ` + "`" + `` + "`" + `ImageId` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::RDS::CustomDBEngineVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRdsCustomDbEngineVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsCustomDbEngineVersion), &result)
	return &result
}
