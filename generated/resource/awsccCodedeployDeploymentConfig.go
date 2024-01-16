package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodedeployDeploymentConfig = `{
  "block": {
    "attributes": {
      "compute_platform": {
        "computed": true,
        "description": "The destination platform type for the deployment (Lambda, Server, or ECS).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_config_name": {
        "computed": true,
        "description": "A name for the deployment configuration. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see Name Type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "minimum_healthy_hosts": {
        "computed": true,
        "description": "The minimum number of healthy instances that should be available at any time during the deployment. There are two parameters expected in the input: type and value.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "traffic_routing_config": {
        "computed": true,
        "description": "The configuration that specifies how the deployment traffic is routed.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "time_based_canary": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "canary_interval": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "canary_percentage": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "time_based_linear": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "linear_interval": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "linear_percentage": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "zonal_config": {
        "computed": true,
        "description": "The zonal deployment config that specifies how the zonal deployment behaves",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "first_zone_monitor_duration_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "minimum_healthy_hosts_per_zone": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "monitor_duration_in_seconds": {
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
    "description": "Resource Type definition for AWS::CodeDeploy::DeploymentConfig",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCodedeployDeploymentConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodedeployDeploymentConfig), &result)
	return &result
}
