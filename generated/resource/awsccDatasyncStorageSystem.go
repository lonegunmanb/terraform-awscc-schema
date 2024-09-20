package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncStorageSystem = `{
  "block": {
    "attributes": {
      "agent_arns": {
        "description": "The ARN of the DataSync agent that connects to and reads from the on-premises storage system's management interface.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "cloudwatch_log_group_arn": {
        "computed": true,
        "description": "The ARN of the Amazon CloudWatch log group used to monitor and log discovery job events.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connectivity_status": {
        "computed": true,
        "description": "Indicates whether the DataSync agent can access the on-premises storage system.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A familiar name for the on-premises storage system.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secrets_manager_arn": {
        "computed": true,
        "description": "The ARN of a secret stored by AWS Secrets Manager.",
        "description_kind": "plain",
        "type": "string"
      },
      "server_configuration": {
        "description": "The server name and network port required to connect with the management interface of the on-premises storage system.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "server_hostname": {
              "description": "The domain name or IP address of the storage system's management interface.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "server_port": {
              "computed": true,
              "description": "The network port needed to access the system's management interface",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "server_credentials": {
        "computed": true,
        "description": "The username and password for accessing your on-premises storage system's management interface.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "password": {
              "computed": true,
              "description": "The password for your storage system's management interface",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "username": {
              "computed": true,
              "description": "The username for your storage system's management interface.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "storage_system_arn": {
        "computed": true,
        "description": "The ARN of the on-premises storage system added to DataSync Discovery.",
        "description_kind": "plain",
        "type": "string"
      },
      "system_type": {
        "description": "The type of on-premises storage system that DataSync Discovery will analyze.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key for an AWS resource tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for an AWS resource tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::DataSync::StorageSystem.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatasyncStorageSystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncStorageSystem), &result)
	return &result
}
