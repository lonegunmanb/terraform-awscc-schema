package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMgnNetworkMigrationDefinition = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the network migration definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the network migration definition was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the network migration definition.",
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
      "name": {
        "description": "The name of the network migration definition.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_migration_definition_id": {
        "computed": true,
        "description": "The unique identifier of the network migration definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "scope_tags": {
        "computed": true,
        "description": "Scope tags for the network migration definition.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "source_configurations": {
        "description": "A list of source configurations for the network migration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_environment": {
              "description": "The source environment type.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_s3_configuration": {
              "description": "S3 configuration for source network data.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_bucket": {
                    "description": "The name of the S3 bucket containing source data.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "s3_bucket_owner": {
                    "description": "The AWS account ID of the S3 bucket owner.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "s3_key": {
                    "description": "The S3 key (path) for the source data.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the network migration definition.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target_deployment": {
        "computed": true,
        "description": "The target deployment configuration for the migrated network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_network": {
        "description": "The target network configuration including topology and CIDR ranges.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "inbound_cidr": {
              "computed": true,
              "description": "The CIDR block for inbound traffic in the target network.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "inspection_cidr": {
              "computed": true,
              "description": "The CIDR block for inspection traffic in the target network.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "outbound_cidr": {
              "computed": true,
              "description": "The CIDR block for outbound traffic in the target network.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topology": {
              "description": "The network topology type for the target environment.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "target_s3_configuration": {
        "description": "The S3 configuration for storing the target network artifacts.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket": {
              "description": "The name of the S3 bucket for target artifacts.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_bucket_owner": {
              "description": "The AWS account ID of the S3 bucket owner.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the network migration definition was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::MGN::NetworkMigrationDefinition",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMgnNetworkMigrationDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMgnNetworkMigrationDefinition), &result)
	return &result
}
