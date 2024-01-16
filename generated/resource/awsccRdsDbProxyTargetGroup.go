package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbProxyTargetGroup = `{
  "block": {
    "attributes": {
      "connection_pool_configuration_info": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connection_borrow_timeout": {
              "computed": true,
              "description": "The number of seconds for a proxy to wait for a connection to become available in the connection pool.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "init_query": {
              "computed": true,
              "description": "One or more SQL statements for the proxy to run when opening each new database connection.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_connections_percent": {
              "computed": true,
              "description": "The maximum size of the connection pool for each target in a target group.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_idle_connections_percent": {
              "computed": true,
              "description": "Controls how actively the proxy closes idle database connections in the connection pool.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "session_pinning_filters": {
              "computed": true,
              "description": "Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "db_cluster_identifiers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "db_instance_identifiers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "db_proxy_name": {
        "description": "The identifier for the proxy.",
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
      "target_group_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) representing the target group.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_group_name": {
        "description": "The identifier for the DBProxyTargetGroup",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::RDS::DBProxyTargetGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsDbProxyTargetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbProxyTargetGroup), &result)
	return &result
}
