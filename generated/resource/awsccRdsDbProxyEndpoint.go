package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbProxyEndpoint = `{
  "block": {
    "attributes": {
      "db_proxy_endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the DB proxy endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_proxy_endpoint_name": {
        "description": "The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_proxy_name": {
        "description": "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_default": {
        "computed": true,
        "description": "A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.",
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "target_role": {
        "computed": true,
        "description": "A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "VPC ID to associate with the new DB proxy endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "VPC security group IDs to associate with the new DB proxy endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_subnet_ids": {
        "description": "VPC subnet IDs to associate with the new DB proxy endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Resource schema for AWS::RDS::DBProxyEndpoint.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsDbProxyEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbProxyEndpoint), &result)
	return &result
}
