package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecordSet = `{
  "block": {
    "attributes": {
      "alias_target": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dns_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "evaluate_target_health": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "hosted_zone_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "cidr_routing_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "collection_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "location_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "comment": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "failover": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "geo_location": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "continent_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "country_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subdivision_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "geo_proximity_location": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "bias": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "coordinates": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "latitude": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "longitude": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "local_zone_group": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "health_check_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "multi_value_answer": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "record_set_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_records": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "set_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ttl": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "weight": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::Route53::RecordSet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53RecordSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecordSet), &result)
	return &result
}
