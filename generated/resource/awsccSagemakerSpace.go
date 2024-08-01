package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerSpace = `{
  "block": {
    "attributes": {
      "domain_id": {
        "description": "The ID of the associated Domain.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ownership_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "owner_user_profile_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "space_arn": {
        "computed": true,
        "description": "The space Amazon Resource Name (ARN).",
        "description_kind": "plain",
        "type": "string"
      },
      "space_display_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "space_name": {
        "description": "A name for the Space.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "space_settings": {
        "computed": true,
        "description": "A collection of settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "app_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "code_editor_app_settings": {
              "computed": true,
              "description": "The CodeEditor app settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "default_resource_spec": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_type": {
                          "computed": true,
                          "description": "The instance type that the image version runs on.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The ARN of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The ARN of the image version created on the instance.",
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
            "custom_file_systems": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "efs_file_system": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "file_system_id": {
                          "description_kind": "plain",
                          "required": true,
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
            "jupyter_lab_app_settings": {
              "computed": true,
              "description": "The JupyterLab app settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "code_repositories": {
                    "computed": true,
                    "description": "A list of CodeRepositories available for use with JupyterLab apps.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "repository_url": {
                          "description": "A CodeRepository (valid URL) to be used within Jupyter's Git extension.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "default_resource_spec": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_type": {
                          "computed": true,
                          "description": "The instance type that the image version runs on.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The ARN of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The ARN of the image version created on the instance.",
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
            "jupyter_server_app_settings": {
              "computed": true,
              "description": "The Jupyter server's app settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "default_resource_spec": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_type": {
                          "computed": true,
                          "description": "The instance type that the image version runs on.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The ARN of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The ARN of the image version created on the instance.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "lifecycle_config_arns": {
                    "computed": true,
                    "description": "A list of LifecycleConfigArns available for use with JupyterServer apps.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "kernel_gateway_app_settings": {
              "computed": true,
              "description": "The kernel gateway app settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "custom_images": {
                    "computed": true,
                    "description": "A list of custom SageMaker images that are configured to run as a KernelGateway app.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "app_image_config_name": {
                          "description": "The Name of the AppImageConfig.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "image_name": {
                          "description": "The name of the CustomImage. Must be unique to your account.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "image_version_number": {
                          "computed": true,
                          "description": "The version number of the CustomImage.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "default_resource_spec": {
                    "computed": true,
                    "description": "The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_type": {
                          "computed": true,
                          "description": "The instance type that the image version runs on.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The ARN of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The ARN of the image version created on the instance.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "lifecycle_config_arns": {
                    "computed": true,
                    "description": "A list of LifecycleConfigArns available for use with KernelGateway apps.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "space_storage_settings": {
              "computed": true,
              "description": "Default storage settings for a space.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ebs_storage_settings": {
                    "computed": true,
                    "description": "Properties related to the space's Amazon Elastic Block Store volume.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ebs_volume_size_in_gb": {
                          "description": "Size of the Amazon EBS volume in Gb",
                          "description_kind": "plain",
                          "required": true,
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
      "space_sharing_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sharing_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the space.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::Space",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerSpaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerSpace), &result)
	return &result
}
