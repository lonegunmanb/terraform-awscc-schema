package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbProxy = `{
  "block": {
    "attributes": {
      "auth": {
        "computed": true,
        "description": "The authorization mechanism that the proxy uses.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auth_scheme": {
              "computed": true,
              "description": "The type of authentication that the proxy uses for connections from the proxy to the underlying database. ",
              "description_kind": "plain",
              "type": "string"
            },
            "client_password_auth_type": {
              "computed": true,
              "description": "The type of authentication the proxy uses for connections from clients.",
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "A user-specified description about the authentication used by a proxy to log in as a specific database user. ",
              "description_kind": "plain",
              "type": "string"
            },
            "iam_auth": {
              "computed": true,
              "description": "Whether to require or disallow Amazon Web Services Identity and Access Management (IAM) authentication for connections to the proxy. The ENABLED value is valid only for proxies with RDS for Microsoft SQL Server.",
              "description_kind": "plain",
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "db_proxy_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the proxy.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_proxy_name": {
        "computed": true,
        "description": "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "debug_logging": {
        "computed": true,
        "description": "Whether the proxy includes detailed information about SQL statements in its logs.",
        "description_kind": "plain",
        "type": "bool"
      },
      "default_auth_scheme": {
        "computed": true,
        "description": "The default authentication scheme that the proxy uses for client connections to the proxy and connections from the proxy to the underlying database.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_network_type": {
        "computed": true,
        "description": "The network type of the DB proxy endpoint. The network type determines the IP version that the proxy endpoint supports.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_family": {
        "computed": true,
        "description": "The kinds of databases that the proxy can connect to.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "idle_client_timeout": {
        "computed": true,
        "description": "The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.",
        "description_kind": "plain",
        "type": "number"
      },
      "require_tls": {
        "computed": true,
        "description": "A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.",
        "description_kind": "plain",
        "type": "bool"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_connection_network_type": {
        "computed": true,
        "description": "The network type that the proxy uses to connect to the target database. The network type determines the IP version that the proxy uses for connections to the database.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "VPC ID to associate with the new DB proxy.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "VPC security group IDs to associate with the new proxy.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_subnet_ids": {
        "computed": true,
        "description": "VPC subnet IDs to associate with the new proxy.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::RDS::DBProxy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRdsDbProxySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbProxy), &result)
	return &result
}
