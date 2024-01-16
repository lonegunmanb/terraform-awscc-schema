package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerDomain = `{
  "block": {
    "attributes": {
      "app_network_access_type": {
        "computed": true,
        "description": "Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.",
        "description_kind": "plain",
        "type": "string"
      },
      "app_security_group_management": {
        "computed": true,
        "description": "The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.",
        "description_kind": "plain",
        "type": "string"
      },
      "auth_mode": {
        "computed": true,
        "description": "The mode of authentication that members use to access the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_space_settings": {
        "computed": true,
        "description": "The default space settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "execution_role": {
              "computed": true,
              "description": "The execution role for the space.",
              "description_kind": "plain",
              "type": "string"
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
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the image version created on the instance.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
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
                          "computed": true,
                          "description": "The Name of the AppImageConfig.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_name": {
                          "computed": true,
                          "description": "The name of the CustomImage. Must be unique to your account.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_version_number": {
                          "computed": true,
                          "description": "The version number of the CustomImage.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
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
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the image version created on the instance.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "security_groups": {
              "computed": true,
              "description": "The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "default_user_settings": {
        "computed": true,
        "description": "The default user settings.",
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
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the image version created on the instance.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "lifecycle_config_arns": {
                    "computed": true,
                    "description": "A list of LifecycleConfigArns available for use with CodeEditor apps.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
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
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "file_system_path": {
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
            "custom_posix_user_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "gid": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "uid": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "default_landing_uri": {
              "computed": true,
              "description": "Defines which Amazon SageMaker application users are directed to by default.",
              "description_kind": "plain",
              "type": "string"
            },
            "execution_role": {
              "computed": true,
              "description": "The execution role for the user.",
              "description_kind": "plain",
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
                          "computed": true,
                          "description": "A CodeRepository (valid URL) to be used within Jupyter's Git extension.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "custom_images": {
                    "computed": true,
                    "description": "A list of custom images for use for JupyterLab apps.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "app_image_config_name": {
                          "computed": true,
                          "description": "The Name of the AppImageConfig.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_name": {
                          "computed": true,
                          "description": "The name of the CustomImage. Must be unique to your account.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_version_number": {
                          "computed": true,
                          "description": "The version number of the CustomImage.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
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
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the image version created on the instance.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "lifecycle_config_arns": {
                    "computed": true,
                    "description": "A list of LifecycleConfigArns available for use with JupyterLab apps.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
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
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the image version created on the instance.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
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
                          "computed": true,
                          "description": "The Name of the AppImageConfig.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_name": {
                          "computed": true,
                          "description": "The name of the CustomImage. Must be unique to your account.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_version_number": {
                          "computed": true,
                          "description": "The version number of the CustomImage.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
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
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the image version created on the instance.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "r_session_app_settings": {
              "computed": true,
              "description": "A collection of settings that apply to an RSessionGateway app.",
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
                          "computed": true,
                          "description": "The Name of the AppImageConfig.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_name": {
                          "computed": true,
                          "description": "The name of the CustomImage. Must be unique to your account.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_version_number": {
                          "computed": true,
                          "description": "The version number of the CustomImage.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
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
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the image version created on the instance.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
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
                    "type": "string"
                  },
                  "user_group": {
                    "computed": true,
                    "description": "The level of permissions that the user has within the RStudioServerPro app. This value defaults to User. The Admin value allows the user access to the RStudio Administrative Dashboard.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "security_groups": {
              "computed": true,
              "description": "The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
              "description_kind": "plain",
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
                    "type": "string"
                  },
                  "s3_kms_key_id": {
                    "computed": true,
                    "description": "When NotebookOutputOption is Allowed, the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_output_path": {
                    "computed": true,
                    "description": "When NotebookOutputOption is Allowed, the Amazon S3 bucket used to store the shared notebook snapshots.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "space_storage_settings": {
              "computed": true,
              "description": "Default storage settings for a space.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "default_ebs_storage_settings": {
                    "computed": true,
                    "description": "Properties related to the Amazon Elastic Block Store volume. Must be provided if storage type is Amazon EBS and must not be provided if storage type is not Amazon EBS",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "default_ebs_volume_size_in_gb": {
                          "computed": true,
                          "description": "Default size of the Amazon EBS volume in Gb",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "maximum_ebs_volume_size_in_gb": {
                          "computed": true,
                          "description": "Maximum size of the Amazon EBS volume in Gb. Must be greater than or equal to the DefaultEbsVolumeSizeInGb.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "studio_web_portal": {
              "computed": true,
              "description": "Indicates whether the Studio experience is available to users. If not, users cannot access Studio.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the created domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The domain name.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "A name for the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_settings": {
        "computed": true,
        "description": "A collection of Domain settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "r_studio_server_pro_domain_settings": {
              "computed": true,
              "description": "A collection of settings that update the current configuration for the RStudioServerPro Domain-level app.",
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
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the image version created on the instance.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "domain_execution_role_arn": {
                    "computed": true,
                    "description": "The ARN of the execution role for the RStudioServerPro Domain-level app.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "r_studio_connect_url": {
                    "computed": true,
                    "description": "A URL pointing to an RStudio Connect server.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "r_studio_package_manager_url": {
                    "computed": true,
                    "description": "A URL pointing to an RStudio Package Manager server.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "security_group_ids": {
              "computed": true,
              "description": "The security groups for the Amazon Virtual Private Cloud that the Domain uses for communication between Domain-level apps and user apps.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "home_efs_file_system_id": {
        "computed": true,
        "description": "The ID of the Amazon Elastic File System (EFS) managed by this Domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "SageMaker uses AWS KMS to encrypt the EFS volume attached to the domain with an AWS managed customer master key (CMK) by default.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_id_for_domain_boundary": {
        "computed": true,
        "description": "The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.",
        "description_kind": "plain",
        "type": "string"
      },
      "single_sign_on_application_arn": {
        "computed": true,
        "description": "The ARN of the application managed by SageMaker in IAM Identity Center. This value is only returned for domains created after October 1, 2023.",
        "description_kind": "plain",
        "type": "string"
      },
      "single_sign_on_managed_application_instance_id": {
        "computed": true,
        "description": "The SSO managed application instance ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "The VPC subnets that Studio uses for communication.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the user profile.",
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
          "nesting_mode": "list"
        }
      },
      "url": {
        "computed": true,
        "description": "The URL to the created domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SageMaker::Domain",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerDomain), &result)
	return &result
}
