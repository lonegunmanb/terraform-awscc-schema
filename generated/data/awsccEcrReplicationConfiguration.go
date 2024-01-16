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
        "description": "The RegistryId associated with the aws account.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_configuration": {
        "computed": true,
        "description": "An object representing the replication configuration for a registry.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rules": {
              "computed": true,
              "description": "An array of objects representing the replication rules for a replication configuration. A replication configuration may contain a maximum of 10 rules.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destinations": {
                    "computed": true,
                    "description": "An array of objects representing the details of a replication destination.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "region": {
                          "computed": true,
                          "description": "A Region to replicate to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "registry_id": {
                          "computed": true,
                          "description": "The account ID of the destination registry to replicate to.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "repository_filters": {
                    "computed": true,
                    "description": "An array of objects representing the details of a repository filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "filter": {
                          "computed": true,
                          "description": "The repository filter to be applied for replication.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "filter_type": {
                          "computed": true,
                          "description": "Type of repository filter",
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
