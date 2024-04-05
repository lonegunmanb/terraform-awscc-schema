package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodestarconnectionsSyncConfiguration = `{
  "block": {
    "attributes": {
      "branch": {
        "description": "The name of the branch of the repository from which resources are to be synchronized,",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "config_file": {
        "description": "The source provider repository path of the sync configuration file of the respective SyncType.",
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
      "publish_deployment_status": {
        "computed": true,
        "description": "Whether to enable or disable publishing of deployment status to source providers.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_link_id": {
        "description": "A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "repository_name": {
        "computed": true,
        "description": "The name of the repository that is being synced to.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_name": {
        "description": "The name of the resource that is being synchronized to the repository.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sync_type": {
        "description": "The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "trigger_resource_update_on": {
        "computed": true,
        "description": "When to trigger Git sync to begin the stack update.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Schema for AWS::CodeStarConnections::SyncConfiguration resource which is used to enables an AWS resource to be synchronized from a source-provider.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCodestarconnectionsSyncConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodestarconnectionsSyncConfiguration), &result)
	return &result
}
