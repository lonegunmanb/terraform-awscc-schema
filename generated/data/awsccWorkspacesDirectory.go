package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspacesDirectory = `{
  "block": {
    "attributes": {
      "active_directory_config": {
        "computed": true,
        "description": "Information about the Active Directory config.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "domain_name": {
              "computed": true,
              "description": "The name of the domain.",
              "description_kind": "plain",
              "type": "string"
            },
            "service_account_secret_arn": {
              "computed": true,
              "description": "Indicates the secret ARN on the service account.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "alias": {
        "computed": true,
        "description": "The directory alias.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the WorkSpace directory.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_based_auth_properties": {
        "computed": true,
        "description": "Describes the properties of the certificate-based authentication you want to use with your WorkSpaces.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "certificate_authority_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the Amazon Web Services Certificate Manager Private CA resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description": "The status of the certificate-based authentication properties.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "customer_user_name": {
        "computed": true,
        "description": "The user name for the service account.",
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
        "computed": true,
        "description": "The directory identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "directory_name": {
        "computed": true,
        "description": "The name of the directory.",
        "description_kind": "plain",
        "type": "string"
      },
      "directory_type": {
        "computed": true,
        "description": "The directory type.",
        "description_kind": "plain",
        "type": "string"
      },
      "dns_ip_addresses": {
        "computed": true,
        "description": "The IP addresses of the DNS servers for the directory.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "dns_ipv_6_addresses": {
        "computed": true,
        "description": "The IPv6 addresses of the DNS servers for the directory.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "enable_self_service": {
        "computed": true,
        "description": "Indicates whether self-service capabilities are enabled or disabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint_encryption_mode": {
        "computed": true,
        "description": "Endpoint encryption mode that allows you to configure the specified directory between Standard TLS and FIPS 140-2 validated mode.",
        "description_kind": "plain",
        "type": "string"
      },
      "iam_role_id": {
        "computed": true,
        "description": "The identifier of the IAM role.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "idc_config": {
        "computed": true,
        "description": "Specifies the configurations of the identity center.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the application.",
              "description_kind": "plain",
              "type": "string"
            },
            "instance_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the identity center instance.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "idc_instance_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the identity center instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_group_ids": {
        "computed": true,
        "description": "The identifiers of the IP access control groups associated with the directory.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "microsoft_entra_config": {
        "computed": true,
        "description": "Specifies the configurations of the Microsoft Entra.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_config_secret_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the application config.",
              "description_kind": "plain",
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description": "The identifier of the tenant.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "registration_code": {
        "computed": true,
        "description": "The registration code for the directory.",
        "description_kind": "plain",
        "type": "string"
      },
      "saml_properties": {
        "computed": true,
        "description": "Describes the enablement status, user access URL, and relay state parameter name that are used for configuring federation with an SAML 2.0 identity provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "relay_state_parameter_name": {
              "computed": true,
              "description": "The relay state parameter name supported by the SAML 2.0 identity provider (IdP).",
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description": "Indicates the status of SAML 2.0 authentication.",
              "description_kind": "plain",
              "type": "string"
            },
            "user_access_url": {
              "computed": true,
              "description": "The SAML 2.0 identity provider (IdP) user access URL.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "selfservice_permissions": {
        "computed": true,
        "description": "Describes the self-service permissions for a directory.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "change_compute_type": {
              "computed": true,
              "description": "Specifies whether users can change the compute type (bundle) for their WorkSpace.",
              "description_kind": "plain",
              "type": "string"
            },
            "increase_volume_size": {
              "computed": true,
              "description": "Specifies whether users can increase the volume size of the drives on their WorkSpace.",
              "description_kind": "plain",
              "type": "string"
            },
            "rebuild_workspace": {
              "computed": true,
              "description": "Specifies whether users can rebuild the operating system of a WorkSpace to its original state.",
              "description_kind": "plain",
              "type": "string"
            },
            "restart_workspace": {
              "computed": true,
              "description": "Specifies whether users can restart their WorkSpace.",
              "description_kind": "plain",
              "type": "string"
            },
            "switch_running_mode": {
              "computed": true,
              "description": "Specifies whether users can switch the running mode of their WorkSpace.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "state": {
        "computed": true,
        "description": "The state of the directory's registration with Amazon WorkSpaces.",
        "description_kind": "plain",
        "type": "string"
      },
      "streaming_properties": {
        "computed": true,
        "description": "Describes the streaming properties.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "global_accelerator": {
              "computed": true,
              "description": "Describes the Global Accelerator for directory.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "mode": {
                    "computed": true,
                    "description": "Indicates if Global Accelerator for directory is enabled or disabled.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "preferred_protocol": {
                    "computed": true,
                    "description": "Indicates the preferred protocol for Global Accelerator.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "storage_connectors": {
              "computed": true,
              "description": "Indicates the storage connector used.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connector_type": {
                    "computed": true,
                    "description": "The type of connector used to save user files.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description": "Indicates if the storage connector is enabled or disabled.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "streaming_experience_preferred_protocol": {
              "computed": true,
              "description": "Indicates the type of preferred protocol for the streaming experience.",
              "description_kind": "plain",
              "type": "string"
            },
            "user_settings": {
              "computed": true,
              "description": "Indicates the permission settings associated with the user.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description": "Indicates the type of action.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "maximum_length": {
                    "computed": true,
                    "description": "Indicates the maximum character length for the specified user setting.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "permission": {
                    "computed": true,
                    "description": "Indicates if the setting is enabled or disabled.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "subnet_ids": {
        "computed": true,
        "description": "The identifiers of the subnets for your virtual private cloud (VPC).",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the directory.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "tenancy": {
        "computed": true,
        "description": "Indicates whether your WorkSpace directory is dedicated or shared.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_identity_type": {
        "computed": true,
        "description": "The type of identity management the user is using.",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_access_properties": {
        "computed": true,
        "description": "The device types and operating systems that can be used to access a WorkSpace.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "access_endpoint_config": {
              "computed": true,
              "description": "Describes the access endpoint configuration for a WorkSpace.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_endpoints": {
                    "computed": true,
                    "description": "Indicates a list of access endpoints associated with this directory.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_endpoint_type": {
                          "computed": true,
                          "description": "Indicates the type of access endpoint.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "vpc_endpoint_id": {
                          "computed": true,
                          "description": "Indicates the VPC endpoint to use for access.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "internet_fallback_protocols": {
                    "computed": true,
                    "description": "Indicates a list of protocols that fallback to using the public Internet when streaming over a VPC endpoint is not available.",
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
            "device_type_android": {
              "computed": true,
              "description": "Indicates whether users can use Android and Android-compatible Chrome OS devices to access their WorkSpaces.",
              "description_kind": "plain",
              "type": "string"
            },
            "device_type_chrome_os": {
              "computed": true,
              "description": "Indicates whether users can use Chromebooks to access their WorkSpaces.",
              "description_kind": "plain",
              "type": "string"
            },
            "device_type_ios": {
              "computed": true,
              "description": "Indicates whether users can use iOS devices to access their WorkSpaces.",
              "description_kind": "plain",
              "type": "string"
            },
            "device_type_linux": {
              "computed": true,
              "description": "Indicates whether users can use Linux clients to access their WorkSpaces.",
              "description_kind": "plain",
              "type": "string"
            },
            "device_type_osx": {
              "computed": true,
              "description": "Indicates whether users can use macOS clients to access their WorkSpaces.",
              "description_kind": "plain",
              "type": "string"
            },
            "device_type_web": {
              "computed": true,
              "description": "Indicates whether users can access their WorkSpaces through a web browser.",
              "description_kind": "plain",
              "type": "string"
            },
            "device_type_windows": {
              "computed": true,
              "description": "Indicates whether users can use Windows clients to access their WorkSpaces.",
              "description_kind": "plain",
              "type": "string"
            },
            "device_type_work_spaces_thin_client": {
              "computed": true,
              "description": "Indicates whether users can access their WorkSpaces through a WorkSpaces Thin Client.",
              "description_kind": "plain",
              "type": "string"
            },
            "device_type_zero_client": {
              "computed": true,
              "description": "Indicates whether users can use zero client devices to access their WorkSpaces.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "workspace_creation_properties": {
        "computed": true,
        "description": "The default values that are used to create WorkSpaces.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_security_group_id": {
              "computed": true,
              "description": "The identifier of the default security group to apply to WorkSpaces when they are created.",
              "description_kind": "plain",
              "type": "string"
            },
            "default_ou": {
              "computed": true,
              "description": "The organizational unit (OU) in the directory for the WorkSpace machine accounts.",
              "description_kind": "plain",
              "type": "string"
            },
            "enable_internet_access": {
              "computed": true,
              "description": "Specifies whether to automatically assign an Elastic public IP address to WorkSpaces in this directory by default.",
              "description_kind": "plain",
              "type": "bool"
            },
            "enable_maintenance_mode": {
              "computed": true,
              "description": "Specifies whether maintenance mode is enabled for WorkSpaces.",
              "description_kind": "plain",
              "type": "bool"
            },
            "instance_iam_role_arn": {
              "computed": true,
              "description": "Indicates the IAM role ARN of the instance.",
              "description_kind": "plain",
              "type": "string"
            },
            "user_enabled_as_local_administrator": {
              "computed": true,
              "description": "Specifies whether WorkSpace users are local administrators on their WorkSpaces.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "workspace_directory_description": {
        "computed": true,
        "description": "Description of the directory to register.",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_directory_name": {
        "computed": true,
        "description": "The name of the directory to register.",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_security_group_id": {
        "computed": true,
        "description": "The identifier of the security group that is assigned to new WorkSpaces.",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_type": {
        "computed": true,
        "description": "Indicates whether the directory's WorkSpace type is personal or pools.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::WorkSpaces::Directory",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWorkspacesDirectorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspacesDirectory), &result)
	return &result
}
