package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2NetworkInsightsPath = `{
  "block": {
    "attributes": {
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination_ip": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "filter_at_destination": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_address": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_port_range": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "from_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "to_port": {
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
            "source_address": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_port_range": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "from_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "to_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      "filter_at_source": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_address": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_port_range": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "from_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "to_port": {
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
            "source_address": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_port_range": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "from_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "to_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_insights_path_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_insights_path_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_ip": {
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
      }
    },
    "description": "Resource schema for AWS::EC2::NetworkInsightsPath",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2NetworkInsightsPathSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2NetworkInsightsPath), &result)
	return &result
}
