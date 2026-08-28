package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the network migration definition.",
        "description_kind": "plain",
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
        "type": [
          "map",
          "string"
        ]
      },
      "source_configurations": {
        "computed": true,
        "description": "A list of source configurations for the network migration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_environment": {
              "computed": true,
              "description": "The source environment type.",
              "description_kind": "plain",
              "type": "string"
            },
            "source_s3_configuration": {
              "computed": true,
              "description": "S3 configuration for source network data.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_bucket": {
                    "computed": true,
                    "description": "The name of the S3 bucket containing source data.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_bucket_owner": {
                    "computed": true,
                    "description": "The AWS account ID of the S3 bucket owner.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_key": {
                    "computed": true,
                    "description": "The S3 key (path) for the source data.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "target_deployment": {
        "computed": true,
        "description": "The target deployment configuration for the migrated network.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_network": {
        "computed": true,
        "description": "The target network configuration including topology and CIDR ranges.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "inbound_cidr": {
              "computed": true,
              "description": "The CIDR block for inbound traffic in the target network.",
              "description_kind": "plain",
              "type": "string"
            },
            "inspection_cidr": {
              "computed": true,
              "description": "The CIDR block for inspection traffic in the target network.",
              "description_kind": "plain",
              "type": "string"
            },
            "outbound_cidr": {
              "computed": true,
              "description": "The CIDR block for outbound traffic in the target network.",
              "description_kind": "plain",
              "type": "string"
            },
            "topology": {
              "computed": true,
              "description": "The network topology type for the target environment.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "target_s3_configuration": {
        "computed": true,
        "description": "The S3 configuration for storing the target network artifacts.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket": {
              "computed": true,
              "description": "The name of the S3 bucket for target artifacts.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_bucket_owner": {
              "computed": true,
              "description": "The AWS account ID of the S3 bucket owner.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the network migration definition was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::MGN::NetworkMigrationDefinition",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMgnNetworkMigrationDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMgnNetworkMigrationDefinition), &result)
	return &result
}
