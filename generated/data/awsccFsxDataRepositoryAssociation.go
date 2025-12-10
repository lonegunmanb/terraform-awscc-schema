package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFsxDataRepositoryAssociation = `{
  "block": {
    "attributes": {
      "association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "batch_import_meta_data_on_create": {
        "computed": true,
        "description": "A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "bool"
      },
      "data_repository_path": {
        "computed": true,
        "description": "The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format ` + "`" + `` + "`" + `s3://myBucket/myPrefix/` + "`" + `` + "`" + `. This path specifies where in the S3 data repository files will be imported from or exported to.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_id": {
        "computed": true,
        "description": "The ID of the file system on which the data repository association is configured.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_path": {
        "computed": true,
        "description": "A path on the Amazon FSx for Lustre file system that points to a high-level directory (such as ` + "`" + `` + "`" + `/ns1/` + "`" + `` + "`" + `) or subdirectory (such as ` + "`" + `` + "`" + `/ns1/subdir/` + "`" + `` + "`" + `) that will be mapped 1-1 with ` + "`" + `` + "`" + `DataRepositoryPath` + "`" + `` + "`" + `. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path ` + "`" + `` + "`" + `/ns1/` + "`" + `` + "`" + `, then you cannot link another data repository with file system path ` + "`" + `` + "`" + `/ns1/ns2` + "`" + `` + "`" + `.\n This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.\n  If you specify only a forward slash (` + "`" + `` + "`" + `/` + "`" + `` + "`" + `) as the file system path, you can link only one data repository to the file system. You can only specify \"/\" as the file system path for the first data repository associated with a file system.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "imported_file_chunk_size": {
        "computed": true,
        "description": "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system or cache.\n The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.",
        "description_kind": "plain",
        "type": "number"
      },
      "resource_arn": {
        "computed": true,
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
              "description": "Describes a data repository association's automatic export policy. The ` + "`" + `` + "`" + `AutoExportPolicy` + "`" + `` + "`" + ` defines the types of updated objects on the file system that will be automatically exported to the data repository. As you create, modify, or delete files, Amazon FSx for Lustre automatically exports the defined changes asynchronously once your application finishes modifying the file.\n The ` + "`" + `` + "`" + `AutoExportPolicy` + "`" + `` + "`" + ` is only supported on Amazon FSx for Lustre file systems with a data repository association.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "events": {
                    "computed": true,
                    "description": "The ` + "`" + `` + "`" + `AutoExportPolicy` + "`" + `` + "`" + ` can have the following event values:\n  +   ` + "`" + `` + "`" + `NEW` + "`" + `` + "`" + ` - New files and directories are automatically exported to the data repository as they are added to the file system.\n  +   ` + "`" + `` + "`" + `CHANGED` + "`" + `` + "`" + ` - Changes to files and directories on the file system are automatically exported to the data repository.\n  +   ` + "`" + `` + "`" + `DELETED` + "`" + `` + "`" + ` - Files and directories are automatically deleted on the data repository when they are deleted on the file system.\n  \n You can define any combination of event types for your ` + "`" + `` + "`" + `AutoExportPolicy` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "auto_import_policy": {
              "computed": true,
              "description": "Describes the data repository association's automatic import policy. The AutoImportPolicy defines how Amazon FSx keeps your file metadata and directory listings up to date by importing changes to your Amazon FSx for Lustre file system as you modify objects in a linked S3 bucket.\n The ` + "`" + `` + "`" + `AutoImportPolicy` + "`" + `` + "`" + ` is only supported on Amazon FSx for Lustre file systems with a data repository association.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "events": {
                    "computed": true,
                    "description": "The ` + "`" + `` + "`" + `AutoImportPolicy` + "`" + `` + "`" + ` can have the following event values:\n  +   ` + "`" + `` + "`" + `NEW` + "`" + `` + "`" + ` - Amazon FSx automatically imports metadata of files added to the linked S3 bucket that do not currently exist in the FSx file system.\n  +   ` + "`" + `` + "`" + `CHANGED` + "`" + `` + "`" + ` - Amazon FSx automatically updates file metadata and invalidates existing file content on the file system as files change in the data repository.\n  +   ` + "`" + `` + "`" + `DELETED` + "`" + `` + "`" + ` - Amazon FSx automatically deletes files on the file system as corresponding files are deleted in the data repository.\n  \n You can define any combination of event types for your ` + "`" + `` + "`" + `AutoImportPolicy` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.\n For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A value that specifies the ` + "`" + `` + "`" + `TagKey` + "`" + `` + "`" + `, the name of the tag. Tag keys must be unique for the resource to which they are attached.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A value that specifies the ` + "`" + `` + "`" + `TagValue` + "`" + `` + "`" + `, the value assigned to the corresponding tag key. Tag values can be null and don't have to be unique in a tag set. For example, you can have a key-value pair in a tag set of ` + "`" + `` + "`" + `finances : April` + "`" + `` + "`" + ` and also of ` + "`" + `` + "`" + `payroll : April` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::FSx::DataRepositoryAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFsxDataRepositoryAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFsxDataRepositoryAssociation), &result)
	return &result
}
