package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccInternetmonitorMonitor = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)",
        "description_kind": "plain",
        "type": "string"
      },
      "health_events_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_local_health_events_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "health_score_threshold": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_traffic_impact": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "availability_score_threshold": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "performance_local_health_events_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "health_score_threshold": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_traffic_impact": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "performance_score_threshold": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "internet_measurements_log_delivery": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bucket_prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_delivery_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "max_city_networks_to_monitor": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "modified_at": {
        "computed": true,
        "description": "The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "processing_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "processing_status_info": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resources": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "resources_to_add": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "resources_to_remove": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      "traffic_percentage_to_monitor": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Represents a monitor, which defines the monitoring boundaries for measurements that Internet Monitor publishes information about for an application",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccInternetmonitorMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInternetmonitorMonitor), &result)
	return &result
}
