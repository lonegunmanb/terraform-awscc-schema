package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspacesinstancesWorkspaceInstance = `{
  "block": {
    "attributes": {
      "ec2_managed_instance": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "instance_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_instance": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
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
                        "throughput": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
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
            "capacity_reservation_specification": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity_reservation_preference": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "capacity_reservation_target": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "capacity_reservation_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "capacity_reservation_resource_group_arn": {
                          "computed": true,
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
            "disable_api_stop": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "ebs_optimized": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "enable_primary_ipv_6": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
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
            "iam_instance_profile": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "image_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "instance_market_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "market_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "spot_options": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_interruption_behavior": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "max_price": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "spot_instance_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "valid_until_utc": {
                          "computed": true,
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
            "key_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
            "maintenance_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auto_recovery": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "metadata_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "http_endpoint": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "http_protocol_ipv_6": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "http_put_response_hop_limit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "http_tokens": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance_metadata_tags": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "monitoring": {
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
            "network_interfaces": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "device_index": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "groups": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
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
            "network_performance_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bandwidth_weighting": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "placement": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_zone": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "group_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "partition_number": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "tenancy": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
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
            "subnet_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tag_specifications": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "resource_type": {
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
                  }
                },
                "nesting_mode": "list"
              }
            },
            "user_data": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "provision_state": {
        "computed": true,
        "description": "The current state of the workspace instance",
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
      "workspace_instance_id": {
        "computed": true,
        "description": "Unique identifier for the workspace instance",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::WorkspacesInstances::WorkspaceInstance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWorkspacesinstancesWorkspaceInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspacesinstancesWorkspaceInstance), &result)
	return &result
}
