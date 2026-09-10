package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2ApplicationStatusCheck = `{
  "block": {
    "attributes": {
      "aggregation": {
        "computed": true,
        "description": "Whether this check is included in the rolled-up application status.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_status_check_id": {
        "computed": true,
        "description": "The unique identifier of the application status check.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the application status check.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "When the application status check was created (ISO 8601).",
        "description_kind": "plain",
        "type": "string"
      },
      "device_index": {
        "computed": true,
        "description": "The network interface device index used for the health check.",
        "description_kind": "plain",
        "type": "number"
      },
      "failure_threshold": {
        "computed": true,
        "description": "The number of consecutive failed probes required to mark the instance unhealthy.",
        "description_kind": "plain",
        "type": "number"
      },
      "health_check_paths": {
        "computed": true,
        "description": "The source/destination network paths used for the health check.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destinations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subnet_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "source": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subnet_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "initialization_grace_period_seconds": {
        "computed": true,
        "description": "Seconds to wait after instance launch before beginning health checks.",
        "description_kind": "plain",
        "type": "number"
      },
      "interval": {
        "computed": true,
        "description": "The interval, in seconds, between health check probes.",
        "description_kind": "plain",
        "type": "number"
      },
      "ip_scope": {
        "computed": true,
        "description": "The IP scope used for the health check.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_version": {
        "computed": true,
        "description": "The IP version used for the health check.",
        "description_kind": "plain",
        "type": "string"
      },
      "path": {
        "computed": true,
        "description": "The HTTP path used for the health check.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port used for the health check.",
        "description_kind": "plain",
        "type": "number"
      },
      "protocol": {
        "computed": true,
        "description": "The network protocol used for the health check.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_code_matcher": {
        "computed": true,
        "description": "The HTTP status codes considered successful (e.g., \"200-299\").",
        "description_kind": "plain",
        "type": "string"
      },
      "success_threshold": {
        "computed": true,
        "description": "The number of consecutive successful probes required to mark the instance healthy.",
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "Tags to apply to the application status check.",
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
          "nesting_mode": "set"
        }
      },
      "timeout": {
        "computed": true,
        "description": "The timeout, in seconds, for each health check probe.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::EC2::ApplicationStatusCheck",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2ApplicationStatusCheckSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2ApplicationStatusCheck), &result)
	return &result
}
