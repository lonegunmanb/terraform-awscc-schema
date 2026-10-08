package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbClusterEndpoint = `{
  "block": {
    "attributes": {
      "custom_endpoint_type": {
        "description": "The type of the custom endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_cluster_endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the DB cluster endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_endpoint_identifier": {
        "description": "The identifier to use for the new endpoint. This parameter is stored as a lowercase string.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_cluster_endpoint_resource_identifier": {
        "computed": true,
        "description": "A unique system-generated identifier for an endpoint. It remains the same for the whole life of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_identifier": {
        "description": "The DB cluster identifier of the DB cluster associated with the endpoint. This parameter is stored as a lowercase string.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "The DNS address of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_type": {
        "computed": true,
        "description": "The endpoint classification returned by Amazon RDS.",
        "description_kind": "plain",
        "type": "string"
      },
      "excluded_members": {
        "computed": true,
        "description": "List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Once either member list is set, it can be changed or replaced by the other list, but both lists cannot be removed in place; removing them requires replacing the endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "static_members": {
        "computed": true,
        "description": "List of DB instance identifiers that are part of the custom endpoint group. Once either member list is set, it can be changed or replaced by the other list, but both lists cannot be removed in place; removing them requires replacing the endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description": "The current status of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to be assigned to the DB cluster endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
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
    "description": "Resource Type definition for AWS::RDS::DBClusterEndpoint. Creates a new custom endpoint and associates it with an Amazon Aurora DB cluster. This resource applies only to Aurora DB clusters.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsDbClusterEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbClusterEndpoint), &result)
	return &result
}
