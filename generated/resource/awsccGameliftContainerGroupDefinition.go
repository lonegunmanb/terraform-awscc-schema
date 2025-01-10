package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGameliftContainerGroupDefinition = `{
  "block": {
    "attributes": {
      "container_group_definition_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift container group resource and uniquely identifies it across all AWS Regions.",
        "description_kind": "plain",
        "type": "string"
      },
      "container_group_type": {
        "computed": true,
        "description": "The scope of the container group",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example \"1469498468.057\").",
        "description_kind": "plain",
        "type": "string"
      },
      "game_server_container_definition": {
        "computed": true,
        "description": "Specifies the information required to run game servers with this container group",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_name": {
              "computed": true,
              "description": "A descriptive label for the container definition. Container definition names must be unique with a container group definition.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "depends_on": {
              "computed": true,
              "description": "A list of container dependencies that determines when this container starts up and shuts down. For container groups with multiple containers, dependencies let you define a startup/shutdown sequence across the containers.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "condition": {
                    "computed": true,
                    "description": "The type of dependency.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "container_name": {
                    "computed": true,
                    "description": "A descriptive label for the container definition. The container being defined depends on this container's condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "environment_override": {
              "computed": true,
              "description": "The environment variables to pass to a container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The environment variable name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The environment variable value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "image_uri": {
              "computed": true,
              "description": "Specifies the image URI of this container.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mount_points": {
              "computed": true,
              "description": "A list of mount point configurations to be used in a container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_level": {
                    "computed": true,
                    "description": "The access permissions for the mounted path.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "container_path": {
                    "computed": true,
                    "description": "The path inside the container where the mount is accessible.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_path": {
                    "computed": true,
                    "description": "The path on the host that will be mounted in the container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "port_configuration": {
              "computed": true,
              "description": "Defines the ports on the container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_port_ranges": {
                    "computed": true,
                    "description": "Specifies one or more ranges of ports on a container.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from_port": {
                          "computed": true,
                          "description": "A starting value for the range of allowed port numbers.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "protocol": {
                          "computed": true,
                          "description": "Defines the protocol of these ports.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "to_port": {
                          "computed": true,
                          "description": "An ending value for the range of allowed port numbers. Port numbers are end-inclusive. This value must be equal to or greater than FromPort.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "resolved_image_digest": {
              "computed": true,
              "description": "The digest of the container image.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server_sdk_version": {
              "computed": true,
              "description": "The version of the server SDK used in this container group",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
      "name": {
        "description": "A descriptive label for the container group definition.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "operating_system": {
        "description": "The operating system of the container group",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_version_number": {
        "computed": true,
        "description": "A specific ContainerGroupDefinition version to be updated",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "A string indicating ContainerGroupDefinition status.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "A string indicating the reason for ContainerGroupDefinition status.",
        "description_kind": "plain",
        "type": "string"
      },
      "support_container_definitions": {
        "computed": true,
        "description": "A collection of support container definitions that define the containers in this group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_name": {
              "computed": true,
              "description": "A descriptive label for the container definition.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "depends_on": {
              "computed": true,
              "description": "A list of container dependencies that determines when this container starts up and shuts down. For container groups with multiple containers, dependencies let you define a startup/shutdown sequence across the containers.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "condition": {
                    "computed": true,
                    "description": "The type of dependency.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "container_name": {
                    "computed": true,
                    "description": "A descriptive label for the container definition. The container being defined depends on this container's condition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "environment_override": {
              "computed": true,
              "description": "The environment variables to pass to a container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The environment variable name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The environment variable value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "essential": {
              "computed": true,
              "description": "Specifies if the container is essential. If an essential container fails a health check, then all containers in the container group will be restarted. You must specify exactly 1 essential container in a container group.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "health_check": {
              "computed": true,
              "description": "Specifies how the health of the containers will be checked.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "command": {
                    "computed": true,
                    "description": "A string array representing the command that the container runs to determine if it is healthy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "interval": {
                    "computed": true,
                    "description": "How often (in seconds) the health is checked.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "retries": {
                    "computed": true,
                    "description": "How many times the process manager will retry the command after a timeout. (The first run of the command does not count as a retry.)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "start_period": {
                    "computed": true,
                    "description": "The optional grace period (in seconds) to give a container time to boostrap before teh health check is declared failed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "timeout": {
                    "computed": true,
                    "description": "How many seconds the process manager allows the command to run before canceling it.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "image_uri": {
              "computed": true,
              "description": "Specifies the image URI of this container.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "memory_hard_limit_mebibytes": {
              "computed": true,
              "description": "The total memory limit of container groups following this definition in MiB",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "mount_points": {
              "computed": true,
              "description": "A list of mount point configurations to be used in a container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_level": {
                    "computed": true,
                    "description": "The access permissions for the mounted path.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "container_path": {
                    "computed": true,
                    "description": "The path inside the container where the mount is accessible.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_path": {
                    "computed": true,
                    "description": "The path on the host that will be mounted in the container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "port_configuration": {
              "computed": true,
              "description": "Defines the ports on the container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_port_ranges": {
                    "computed": true,
                    "description": "Specifies one or more ranges of ports on a container.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from_port": {
                          "computed": true,
                          "description": "A starting value for the range of allowed port numbers.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "protocol": {
                          "computed": true,
                          "description": "Defines the protocol of these ports.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "to_port": {
                          "computed": true,
                          "description": "An ending value for the range of allowed port numbers. Port numbers are end-inclusive. This value must be equal to or greater than FromPort.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "resolved_image_digest": {
              "computed": true,
              "description": "The digest of the container image.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vcpu": {
              "computed": true,
              "description": "The number of virtual CPUs to give to the support group",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "total_memory_limit_mebibytes": {
        "description": "The total memory limit of container groups following this definition in MiB",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "total_vcpu_limit": {
        "description": "The total amount of virtual CPUs on the container group definition",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "version_description": {
        "computed": true,
        "description": "The description of this version",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_number": {
        "computed": true,
        "description": "The version of this ContainerGroupDefinition",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "The AWS::GameLift::ContainerGroupDefinition resource creates an Amazon GameLift container group definition.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGameliftContainerGroupDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGameliftContainerGroupDefinition), &result)
	return &result
}
