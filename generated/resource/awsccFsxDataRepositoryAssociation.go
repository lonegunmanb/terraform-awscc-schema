package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFsxDataRepositoryAssociation = `{
  "block": {
    "attributes": {
      "association_id": {
        "computed": true,
        "description": "The system-generated, unique ID of the data repository association.",
        "description_kind": "plain",
        "type": "string"
      },
      "batch_import_meta_data_on_create": {
        "computed": true,
        "description": "A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "data_repository_path": {
        "description": "The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format s3://myBucket/myPrefix/ . This path specifies where in the S3 data repository files will be imported from or exported to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "file_system_id": {
        "description": "The globally unique ID of the file system, assigned by Amazon FSx.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "file_system_path": {
        "description": "This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.",
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
      "imported_file_chunk_size": {
        "computed": true,
        "description": "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for a given resource. ARNs uniquely identify Amazon Web Services resources. We require an ARN when you need to specify a resource unambiguously across all of Amazon Web Services. For more information, see Amazon Resource Names (ARNs) in the Amazon Web Services General Reference.",
        "description_kind": "plain",
        "type": "string"
      },
      "s3": {
        "computed": true,
        "description": "The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_export_policy": {
              "computed": true,
              "description": "Specifies the type of updated objects (new, changed, deleted) that will be automatically exported from your file system to the linked S3 bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "events": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "auto_import_policy": {
              "computed": true,
              "description": "Specifies the type of updated objects (new, changed, deleted) that will be automatically imported from the linked S3 bucket to your file system.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "events": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "A list of Tag values, with a maximum of 50 elements.",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::FSx::DataRepositoryAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccFsxDataRepositoryAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFsxDataRepositoryAssociation), &result)
	return &result
}
