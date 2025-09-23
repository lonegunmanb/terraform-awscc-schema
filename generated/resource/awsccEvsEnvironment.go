package resource

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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "private_route_server_peerings": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "host_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "placement_group_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "expansion_vlan_1": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "expansion_vlan_2": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "hcx": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "hcx_network_acl_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "is_hcx_public": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "v_motion": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "v_san": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "v_tep": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "vm_management": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "vmk_management": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
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
          "nesting_mode": "single"
        },
        "optional": true
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "license_info": {
        "description": "The license information for an EVS environment",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "solution_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vsan_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
      "service_access_subnet_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
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
      "terms_accepted": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "vcf_hostnames": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloud_builder": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "nsx": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "nsx_edge_1": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "nsx_edge_2": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "nsx_manager_1": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "nsx_manager_2": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "nsx_manager_3": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sddc_manager": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "v_center": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "vcf_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "An environment created within the EVS service",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEvsEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEvsEnvironment), &result)
	return &result
}
