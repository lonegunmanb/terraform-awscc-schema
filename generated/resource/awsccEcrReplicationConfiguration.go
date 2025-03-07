package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrReplicationConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "registry_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "replication_configuration": {
        "description": "The replication configuration for a registry.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rules": {
              "description": "An array of objects representing the replication destinations and repository filters for a replication configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destinations": {
                    "description": "An array of objects representing the destination for a replication rule.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "region": {
                          "description": "The Region to replicate to.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "registry_id": {
                          "description": "The AWS account ID of the Amazon ECR private registry to replicate to. When configuring cross-Region replication within your own registry, specify your own account ID.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "required": true
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
                          "optional": true,
                          "type": "string"
                        },
                        "filter_type": {
                          "computed": true,
                          "description": "The repository filter type. The only supported value is ` + "`" + `` + "`" + `PREFIX_MATCH` + "`" + `` + "`" + `, which is a repository name prefix specified with the ` + "`" + `` + "`" + `filter` + "`" + `` + "`" + ` parameter.",
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
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ECR::ReplicationConfiguration` + "`" + `` + "`" + ` resource creates or updates the replication configuration for a private registry. The first time a replication configuration is applied to a private registry, a service-linked IAM role is created in your account for the replication process. For more information, see [Using Service-Linked Roles for Amazon ECR](https://docs.aws.amazon.com/AmazonECR/latest/userguide/using-service-linked-roles.html) in the *Amazon Elastic Container Registry User Guide*.\n  When configuring cross-account replication, the destination account must grant the source account permission to replicate. This permission is controlled using a private registry permissions policy. For more information, see ` + "`" + `` + "`" + `AWS::ECR::RegistryPolicy` + "`" + `` + "`" + `.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcrReplicationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrReplicationConfiguration), &result)
	return &result
}
