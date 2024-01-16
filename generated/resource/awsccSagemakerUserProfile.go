package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerUserProfile = `{
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
      "single_sign_on_user_identifier": {
        "computed": true,
        "description": "A specifier for the type of value specified in SingleSignOnUserValue. Currently, the only supported value is \"UserName\". If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "single_sign_on_user_value": {
        "computed": true,
        "description": "The username of the associated AWS Single Sign-On User for this UserProfile. If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the user profile.",
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
      "user_profile_arn": {
        "computed": true,
        "description": "The user profile Amazon Resource Name (ARN).",
        "description_kind": "plain",
        "type": "string"
      },
      "user_profile_name": {
        "description": "A name for the UserProfile.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_settings": {
        "computed": true,
        "description": "A collection of settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code_editor_app_settings": {
              "computed": true,
              "description": "The CodeEditor app settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "default_resource_spec": {
                    "computed": true,
                    "description": "The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the CodeEditor app.",
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
                    "description": "A list of LifecycleConfigArns available for use with CodeEditor apps.",
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
            "custom_file_system_configs": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "efs_file_system_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "file_system_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "file_system_path": {
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
            "custom_posix_user_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "gid": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "uid": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "default_landing_uri": {
              "computed": true,
              "description": "Defines which Amazon SageMaker application users are directed to by default.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "execution_role": {
              "computed": true,
              "description": "The user profile Amazon Resource Name (ARN).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
                  "custom_images": {
                    "computed": true,
                    "description": "A list of custom images available for use for JupyterLab apps",
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
                    "description": "The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the JupyterLab app.",
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
                    "description": "A list of LifecycleConfigArns available for use with JupyterLab apps.",
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
            "r_studio_server_pro_app_settings": {
              "computed": true,
              "description": "A collection of settings that configure user interaction with the RStudioServerPro app.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_status": {
                    "computed": true,
                    "description": "Indicates whether the current user has access to the RStudioServerPro app.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_group": {
                    "computed": true,
                    "description": "The level of permissions that the user has within the RStudioServerPro app. This value defaults to User. The Admin value allows the user access to the RStudio Administrative Dashboard.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "security_groups": {
              "computed": true,
              "description": "The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "sharing_settings": {
              "computed": true,
              "description": "The sharing settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "notebook_output_option": {
                    "computed": true,
                    "description": "Whether to include the notebook cell output when sharing the notebook. The default is Disabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_kms_key_id": {
                    "computed": true,
                    "description": "When NotebookOutputOption is Allowed, the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_output_path": {
                    "computed": true,
                    "description": "When NotebookOutputOption is Allowed, the Amazon S3 bucket used to store the shared notebook snapshots.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
                  "default_ebs_storage_settings": {
                    "computed": true,
                    "description": "Properties related to the Amazon Elastic Block Store volume.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "default_ebs_volume_size_in_gb": {
                          "description": "Default size of the Amazon EBS volume in Gb",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "maximum_ebs_volume_size_in_gb": {
                          "description": "Maximum size of the Amazon EBS volume in Gb. Must be greater than or equal to the DefaultEbsVolumeSizeInGb.",
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
            },
            "studio_web_portal": {
              "computed": true,
              "description": "Indicates whether the Studio experience is available to users. If not, users cannot access Studio.",
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
    "description": "Resource Type definition for AWS::SageMaker::UserProfile",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerUserProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerUserProfile), &result)
	return &result
}
