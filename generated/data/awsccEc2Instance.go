package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Instance = `{
  "block": {
    "attributes": {
      "additional_info": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "affinity": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "block_device_mappings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "device_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ebs": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delete_on_termination": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "encrypted": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "iops": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "snapshot_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "volume_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "volume_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "no_device": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "virtual_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "cpu_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "core_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "threads_per_core": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "credit_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cpu_credits": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "disable_api_termination": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ebs_optimized": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "elastic_gpu_specifications": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "elastic_inference_accelerators": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "enclave_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "hibernation_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configured": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "host_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "host_resource_group_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "iam_instance_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_initiated_shutdown_behavior": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_6_address_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "ipv_6_addresses": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ipv_6_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "kernel_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "launch_template": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "launch_template_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "launch_template_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "license_specifications": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "license_configuration_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "monitoring": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "network_interfaces": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "associate_carrier_ip_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "associate_public_ip_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "delete_on_termination": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "device_index": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "group_set": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "ipv_6_address_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "ipv_6_addresses": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ipv_6_address": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "network_interface_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "private_ip_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "private_ip_addresses": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "primary": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "private_ip_address": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "secondary_private_ip_address_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
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
      "placement_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns_name_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_resource_name_dns_a_record": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "enable_resource_name_dns_aaaa_record": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "hostname_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "private_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "propagate_tags_to_volume_on_creation": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "public_dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ramdisk_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "security_groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "source_dest_check": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ssm_associations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "association_parameters": {
              "computed": true,
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
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "document_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      "tenancy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volumes": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "device": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "volume_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::EC2::Instance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2InstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Instance), &result)
	return &result
}
