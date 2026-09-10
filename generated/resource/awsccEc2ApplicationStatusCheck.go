package resource

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
        "optional": true,
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
        "optional": true,
        "type": "number"
      },
      "failure_threshold": {
        "computed": true,
        "description": "The number of consecutive failed probes required to mark the instance unhealthy.",
        "description_kind": "plain",
        "optional": true,
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
                    "optional": true,
                    "type": "string"
                  },
                  "subnet_id": {
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
            "source": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnet_id": {
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "initialization_grace_period_seconds": {
        "computed": true,
        "description": "Seconds to wait after instance launch before beginning health checks.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "interval": {
        "computed": true,
        "description": "The interval, in seconds, between health check probes.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ip_scope": {
        "computed": true,
        "description": "The IP scope used for the health check.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_version": {
        "computed": true,
        "description": "The IP version used for the health check.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "path": {
        "computed": true,
        "description": "The HTTP path used for the health check.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "description": "The port used for the health check.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "protocol": {
        "description": "The network protocol used for the health check.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status_code_matcher": {
        "computed": true,
        "description": "The HTTP status codes considered successful (e.g., \"200-299\").",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "success_threshold": {
        "computed": true,
        "description": "The number of consecutive successful probes required to mark the instance healthy.",
        "description_kind": "plain",
        "optional": true,
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
          "nesting_mode": "set"
        },
        "optional": true
      },
      "timeout": {
        "computed": true,
        "description": "The timeout, in seconds, for each health check probe.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "An application status check monitors an HTTP or HTTPS endpoint on Amazon EC2 instances and reports application-layer health. Configure a health check with a Protocol, Port, and Path, and then associate it with EC2 instances via AWS::EC2::ApplicationStatusCheckInstanceAssociation or AWS::EC2::ApplicationStatusCheckTagAssociation.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2ApplicationStatusCheckSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2ApplicationStatusCheck), &result)
	return &result
}
