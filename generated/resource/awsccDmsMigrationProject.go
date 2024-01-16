package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsMigrationProject = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The optional description of the migration project.",
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
      "instance_profile_arn": {
        "computed": true,
        "description": "The property describes an instance profile arn for the migration project. For read",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_profile_identifier": {
        "computed": true,
        "description": "The property describes an instance profile identifier for the migration project. For create",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_profile_name": {
        "computed": true,
        "description": "The property describes an instance profile name for the migration project. For read",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "migration_project_arn": {
        "computed": true,
        "description": "The property describes an ARN of the migration project.",
        "description_kind": "plain",
        "type": "string"
      },
      "migration_project_creation_time": {
        "computed": true,
        "description": "The property describes a creating time of the migration project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "migration_project_identifier": {
        "computed": true,
        "description": "The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "migration_project_name": {
        "computed": true,
        "description": "The property describes a name to identify the migration project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schema_conversion_application_attributes": {
        "computed": true,
        "description": "The property describes schema conversion application attributes for the migration project.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket_path": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_bucket_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "source_data_provider_descriptors": {
        "computed": true,
        "description": "The property describes source data provider descriptors for the migration project.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_provider_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_provider_identifier": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_provider_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, , and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, , and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target_data_provider_descriptors": {
        "computed": true,
        "description": "The property describes target data provider descriptors for the migration project.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_provider_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_provider_identifier": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_provider_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "transformation_rules": {
        "computed": true,
        "description": "The property describes transformation rules for the migration project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::DMS::MigrationProject",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDmsMigrationProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsMigrationProject), &result)
	return &result
}
