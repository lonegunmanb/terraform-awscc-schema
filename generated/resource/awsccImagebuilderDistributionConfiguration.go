package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccImagebuilderDistributionConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the distribution configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the distribution configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "distributions": {
        "description": "The distributions of the distribution configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ami_distribution_configuration": {
              "computed": true,
              "description": "The specific AMI settings (for example, launch permissions, AMI tags).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ami_tags": {
                    "computed": true,
                    "description": "The tags to apply to AMIs distributed to this Region.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "description": {
                    "computed": true,
                    "description": "The description of the AMI distribution configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description": "The KMS key identifier used to encrypt the distributed image.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "launch_permission_configuration": {
                    "computed": true,
                    "description": "Launch permissions can be used to configure which AWS accounts can use the AMI to launch instances.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "organization_arns": {
                          "computed": true,
                          "description": "The ARN for an Amazon Web Services Organization that you want to share your AMI with.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "organizational_unit_arns": {
                          "computed": true,
                          "description": "The ARN for an Organizations organizational unit (OU) that you want to share your AMI with.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "user_groups": {
                          "computed": true,
                          "description": "The name of the group.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "user_ids": {
                          "computed": true,
                          "description": "The AWS account ID.",
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
                  "name": {
                    "computed": true,
                    "description": "The name of the AMI distribution configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_account_ids": {
                    "computed": true,
                    "description": "The ID of accounts to which you want to distribute an image.",
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
            "container_distribution_configuration": {
              "computed": true,
              "description": "Container distribution settings for encryption, licensing, and sharing in a specific Region.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_tags": {
                    "computed": true,
                    "description": "Tags that are attached to the container distribution configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "description": {
                    "computed": true,
                    "description": "The description of the container distribution configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_repository": {
                    "computed": true,
                    "description": "The destination repository for the container distribution configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "repository_name": {
                          "computed": true,
                          "description": "The repository name of target container repository.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "service": {
                          "computed": true,
                          "description": "The service of target container repository.",
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
            "fast_launch_configurations": {
              "computed": true,
              "description": "The Windows faster-launching configurations to use for AMI distribution.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_id": {
                    "computed": true,
                    "description": "The owner account ID for the fast-launch enabled Windows AMI.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description": "A Boolean that represents the current state of faster launching for the Windows AMI. Set to true to start using Windows faster launching, or false to stop using it.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "launch_template": {
                    "computed": true,
                    "description": "The launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "launch_template_id": {
                          "computed": true,
                          "description": "The ID of the launch template to use for faster launching for a Windows AMI.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "launch_template_name": {
                          "computed": true,
                          "description": "The name of the launch template to use for faster launching for a Windows AMI.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "launch_template_version": {
                          "computed": true,
                          "description": "The version of the launch template to use for faster launching for a Windows AMI.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "max_parallel_launches": {
                    "computed": true,
                    "description": "The maximum number of parallel instances that are launched for creating resources.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "snapshot_configuration": {
                    "computed": true,
                    "description": "Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "target_resource_count": {
                          "computed": true,
                          "description": "The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.",
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
                "nesting_mode": "list"
              },
              "optional": true
            },
            "launch_template_configurations": {
              "computed": true,
              "description": "A group of launchTemplateConfiguration settings that apply to image distribution.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_id": {
                    "computed": true,
                    "description": "The account ID that this configuration applies to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "launch_template_id": {
                    "computed": true,
                    "description": "Identifies the EC2 launch template to use.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "set_default_version": {
                    "computed": true,
                    "description": "Set the specified EC2 launch template as the default launch template for the specified account.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "license_configuration_arns": {
              "computed": true,
              "description": "The License Manager Configuration to associate with the AMI in the specified Region.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "region": {
              "description": "region",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the distribution configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the component.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource schema for AWS::ImageBuilder::DistributionConfiguration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccImagebuilderDistributionConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccImagebuilderDistributionConfiguration), &result)
	return &result
}
