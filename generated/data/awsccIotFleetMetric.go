package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotFleetMetric = `{
  "block": {
    "attributes": {
      "aggregation_field": {
        "computed": true,
        "description": "The aggregation field to perform aggregation and metric emission",
        "description_kind": "plain",
        "type": "string"
      },
      "aggregation_type": {
        "computed": true,
        "description": "Aggregation types supported by Fleet Indexing",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "Fleet Indexing aggregation type names such as Statistics, Percentiles and Cardinality",
              "description_kind": "plain",
              "type": "string"
            },
            "values": {
              "computed": true,
              "description": "Fleet Indexing aggregation type values",
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
      "creation_date": {
        "computed": true,
        "description": "The creation date of a fleet metric",
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description": "The description of a fleet metric",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "index_name": {
        "computed": true,
        "description": "The index name of a fleet metric",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_date": {
        "computed": true,
        "description": "The last modified date of a fleet metric",
        "description_kind": "plain",
        "type": "number"
      },
      "metric_arn": {
        "computed": true,
        "description": "The Amazon Resource Number (ARN) of a fleet metric metric",
        "description_kind": "plain",
        "type": "string"
      },
      "metric_name": {
        "computed": true,
        "description": "The name of the fleet metric",
        "description_kind": "plain",
        "type": "string"
      },
      "period": {
        "computed": true,
        "description": "The period of metric emission in seconds",
        "description_kind": "plain",
        "type": "number"
      },
      "query_string": {
        "computed": true,
        "description": "The Fleet Indexing query used by a fleet metric",
        "description_kind": "plain",
        "type": "string"
      },
      "query_version": {
        "computed": true,
        "description": "The version of a Fleet Indexing query used by a fleet metric",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "unit": {
        "computed": true,
        "description": "The unit of data points emitted by a fleet metric",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version of a fleet metric",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::IoT::FleetMetric",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotFleetMetricSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotFleetMetric), &result)
	return &result
}
