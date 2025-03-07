package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrReplicationConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "registry_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "replication_configuration": {
        "computed": true,
        "description": "The replication configuration for a registry.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rules": {
              "computed": true,
              "description": "An array of objects representing the replication destinations and repository filters for a replication configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destinations": {
                    "computed": true,
                    "description": "An array of objects representing the destination for a replication rule.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "region": {
                          "computed": true,
                          "description": "The Region to replicate to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "registry_id": {
                          "computed": true,
                          "description": "The AWS account ID of the Amazon ECR private registry to replicate to. When configuring cross-Region replication within your own registry, specify your own account ID.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "repository_filters": {
                    "computed": true,
                    "description": "An array of objects representing the filters for a replication rule. Specifying a repository filter for a replication rule provides a method for controlling which repositories in a private registry are replicated.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "filter": {
                          "computed": true,
                          "description": "The repository filter details. When the ` + "`" + `` + "`" + `PREFIX_MATCH` + "`" + `` + "`" + ` filter type is specified, this value is required and should be the repository name prefix to configure replication for.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "filter_type": {
                          "computed": true,
                          "description": "The repository filter type. The only supported value is ` + "`" + `` + "`" + `PREFIX_MATCH` + "`" + `` + "`" + `, which is a repository name prefix specified with the ` + "`" + `` + "`" + `filter` + "`" + `` + "`" + ` parameter.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::ECR::ReplicationConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcrReplicationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrReplicationConfiguration), &result)
	return &result
}
