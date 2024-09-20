package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailContainer = `{
  "block": {
    "attributes": {
      "container_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container_service_deployment": {
        "computed": true,
        "description": "Describes a container deployment configuration of an Amazon Lightsail container service.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "containers": {
              "computed": true,
              "description": "An object that describes the configuration for the containers of the deployment.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "command": {
                    "computed": true,
                    "description": "The launch command for the container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "container_name": {
                    "computed": true,
                    "description": "The name of the container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "environment": {
                    "computed": true,
                    "description": "The environment variables of the container.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "variable": {
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
                  "image": {
                    "computed": true,
                    "description": "The name of the image used for the container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ports": {
                    "computed": true,
                    "description": "The open firewall ports of the container.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "port": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "protocol": {
                          "computed": true,
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
                "nesting_mode": "set"
              },
              "optional": true
            },
            "public_endpoint": {
              "computed": true,
              "description": "An object that describes the endpoint of the deployment.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_name": {
                    "computed": true,
                    "description": "The name of the container for the endpoint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "container_port": {
                    "computed": true,
                    "description": "The port of the container to which traffic is forwarded to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "health_check_config": {
                    "computed": true,
                    "description": "An object that describes the health check configuration of the container.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "healthy_threshold": {
                          "computed": true,
                          "description": "The number of consecutive health checks successes required before moving the container to the Healthy state. The default value is 2.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_seconds": {
                          "computed": true,
                          "description": "The approximate interval, in seconds, between health checks of an individual container. You can specify between 5 and 300 seconds. The default value is 5.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "path": {
                          "computed": true,
                          "description": "The path on the container on which to perform the health check. The default value is /.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "success_codes": {
                          "computed": true,
                          "description": "The HTTP codes to use when checking for a successful response from a container. You can specify values between 200 and 499. You can specify multiple values (for example, 200,202) or a range of values (for example, 200-299).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "timeout_seconds": {
                          "computed": true,
                          "description": "The amount of time, in seconds, during which no response means a failed health check. You can specify between 2 and 60 seconds. The default value is 2.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "unhealthy_threshold": {
                          "computed": true,
                          "description": "The number of consecutive health check failures required before moving the container to the Unhealthy state. The default value is 2.",
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
      "is_disabled": {
        "computed": true,
        "description": "A Boolean value to indicate whether the container service is disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "power": {
        "description": "The power specification for the container service.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_arn": {
        "computed": true,
        "description": "The principal ARN of the container service.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_registry_access": {
        "computed": true,
        "description": "A Boolean value to indicate whether the container service has access to private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ecr_image_puller_role": {
              "computed": true,
              "description": "An object to describe a request to activate or deactivate the role that you can use to grant an Amazon Lightsail container service access to Amazon Elastic Container Registry (Amazon ECR) private repositories.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "is_active": {
                    "computed": true,
                    "description": "A Boolean value that indicates whether to activate the role.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "principal_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the role, if it is activated.",
                    "description_kind": "plain",
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
      "public_domain_names": {
        "computed": true,
        "description": "The public domain names to use with the container service, such as example.com and www.example.com.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "certificate_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "domain_names": {
              "computed": true,
              "description": "An object that describes the configuration for the containers of the deployment.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "scale": {
        "description": "The scale specification for the container service.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "service_name": {
        "description": "The name for the container service.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "url": {
        "computed": true,
        "description": "The publicly accessible URL of the container service.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lightsail::Container",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLightsailContainerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailContainer), &result)
	return &result
}
