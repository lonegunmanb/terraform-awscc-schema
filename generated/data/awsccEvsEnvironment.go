package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEvsEnvironment = `{
  "block": {
    "attributes": {
      "checks": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "impaired_since": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "result": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
      "connectivity_info": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "private_route_server_peerings": {
              "computed": true,
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
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "credentials": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "secret_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "environment_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_name": {
        "computed": true,
        "description": "The name of an EVS environment",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hosts": {
        "computed": true,
        "description": "The initial hosts for environment only required upon creation. Modification after creation will have no effect",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dedicated_host_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "host_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "instance_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "placement_group_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "initial_vlans": {
        "computed": true,
        "description": "The initial Vlan configuration only required upon creation. Modification after creation will have no effect",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "edge_v_tep": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "expansion_vlan_1": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "expansion_vlan_2": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "hcx": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "hcx_network_acl_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "is_hcx_public": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "nsx_up_link": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "v_motion": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "v_san": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "v_tep": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vm_management": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vmk_management": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
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
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "license_info": {
        "computed": true,
        "description": "The license information for an EVS environment",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "solution_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vsan_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "modified_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_access_security_groups": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_groups": {
              "computed": true,
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
      "service_access_subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "site_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state_details": {
        "computed": true,
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "terms_accepted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "vcf_hostnames": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloud_builder": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "nsx": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "nsx_edge_1": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "nsx_edge_2": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "nsx_manager_1": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "nsx_manager_2": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "nsx_manager_3": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sddc_manager": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "v_center": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "vcf_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EVS::Environment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEvsEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEvsEnvironment), &result)
	return &result
}
