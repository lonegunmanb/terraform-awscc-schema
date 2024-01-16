package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsReplicationConfig = `{
  "block": {
    "attributes": {
      "compute_config": {
        "computed": true,
        "description": "Configuration parameters for provisioning a AWS DMS Serverless replication",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "dns_name_servers": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "max_capacity_units": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "min_capacity_units": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "multi_az": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "preferred_maintenance_window": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "replication_subnet_group_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_config_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Replication Config",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_config_identifier": {
        "computed": true,
        "description": "A unique identifier of replication configuration",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_settings": {
        "computed": true,
        "description": "JSON settings for Servereless replications that are provisioned using this replication configuration",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_type": {
        "computed": true,
        "description": "The type of AWS DMS Serverless replication to provision using this replication configuration",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_identifier": {
        "computed": true,
        "description": "A unique value or name that you get set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource",
        "description_kind": "plain",
        "type": "string"
      },
      "source_endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration",
        "description_kind": "plain",
        "type": "string"
      },
      "supplemental_settings": {
        "computed": true,
        "description": "JSON settings for specifying supplemental data",
        "description_kind": "plain",
        "type": "string"
      },
      "table_mappings": {
        "computed": true,
        "description": "JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "\u003cp\u003eContains a map of the key-value pairs for the resource tag or tags assigned to the dataset.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "\u003cp\u003eTag key.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DMS::ReplicationConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDmsReplicationConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsReplicationConfig), &result)
	return &result
}
