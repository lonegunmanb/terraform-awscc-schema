package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2FcaPolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the FCA policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the FCA policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fca_policy_id": {
        "computed": true,
        "description": "The ID of the FCA policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_fca_policy_document": {
        "computed": true,
        "description": "The FCA policy document for EC2 instance cost allocation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr_range_specifications": {
              "computed": true,
              "description": "CIDR range specifications, used when the reporting granularity is 'cidr'.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "prefixes": {
                    "computed": true,
                    "description": "The CIDR range to label mappings.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cidr_label": {
                          "computed": true,
                          "description": "The label for the CIDR range.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cidr_range": {
                          "computed": true,
                          "description": "The CIDR range.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "prefixes_type": {
                    "computed": true,
                    "description": "The type of prefixes the mappings apply to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "dimensions": {
              "computed": true,
              "description": "The cost dimensions to allocate.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "reporting_granularity": {
              "computed": true,
              "description": "The granularity at which costs are reported. Contains exactly one value.",
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
      "latest_version": {
        "computed": true,
        "description": "The latest version number of the FCA policy.",
        "description_kind": "plain",
        "type": "number"
      },
      "nat_fca_policy_document": {
        "computed": true,
        "description": "The FCA policy document for NAT gateway cost allocation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr_range_specifications": {
              "computed": true,
              "description": "CIDR range specifications, used when the reporting granularity is 'cidr'.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "prefixes": {
                    "computed": true,
                    "description": "The CIDR range to label mappings.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cidr_label": {
                          "computed": true,
                          "description": "The label for the CIDR range.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cidr_range": {
                          "computed": true,
                          "description": "The CIDR range.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "prefixes_type": {
                    "computed": true,
                    "description": "The type of prefixes the mappings apply to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "dimensions": {
              "computed": true,
              "description": "The cost dimensions to allocate.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "reporting_granularity": {
              "computed": true,
              "description": "The granularity at which costs are reported. Contains exactly one value.",
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
      "network_resource_type": {
        "description": "The type of network resource the FCA policy applies to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_account": {
        "computed": true,
        "description": "The AWS account ID that owns the FCA policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "pcx_fca_policy_document": {
        "computed": true,
        "description": "The FCA policy document for VPC peering connection cost allocation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr_range_specifications": {
              "computed": true,
              "description": "CIDR range specifications, used when the reporting granularity is 'cidr'.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "prefixes": {
                    "computed": true,
                    "description": "The CIDR range to label mappings.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cidr_label": {
                          "computed": true,
                          "description": "The label for the CIDR range.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cidr_range": {
                          "computed": true,
                          "description": "The CIDR range.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "prefixes_type": {
                    "computed": true,
                    "description": "The type of prefixes the mappings apply to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "dimensions": {
              "computed": true,
              "description": "The cost dimensions to allocate.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "reporting_granularity": {
              "computed": true,
              "description": "The granularity at which costs are reported. Contains exactly one value.",
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
      "policy_created_at": {
        "computed": true,
        "description": "The timestamp when the FCA policy was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "show_back_enabled": {
        "computed": true,
        "description": "Indicates whether show-back reporting is enabled for the FCA policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the FCA policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "version_expected_effective_at": {
        "computed": true,
        "description": "The timestamp when the latest version of the FCA policy is expected to take effect.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "A Flex Cost Allocation (FCA) policy defines how shared network costs (data transfer, data processing) are allocated and reported for network resources such as NAT gateways, EC2 instances, and VPC peering connections.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2FcaPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2FcaPolicy), &result)
	return &result
}
