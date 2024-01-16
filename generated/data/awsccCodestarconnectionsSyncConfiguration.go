package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodestarconnectionsSyncConfiguration = `{
  "block": {
    "attributes": {
      "branch": {
        "computed": true,
        "description": "The name of the branch of the repository from which resources are to be synchronized,",
        "description_kind": "plain",
        "type": "string"
      },
      "config_file": {
        "computed": true,
        "description": "The source provider repository path of the sync configuration file of the respective SyncType.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "the ID of the entity that owns the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "provider_type": {
        "computed": true,
        "description": "The name of the external provider where your third-party code repository is configured.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_link_id": {
        "computed": true,
        "description": "A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_name": {
        "computed": true,
        "description": "The name of the repository that is being synced to.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_name": {
        "computed": true,
        "description": "The name of the resource that is being synchronized to the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "sync_type": {
        "computed": true,
        "description": "The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CodeStarConnections::SyncConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodestarconnectionsSyncConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodestarconnectionsSyncConfiguration), &result)
	return &result
}
