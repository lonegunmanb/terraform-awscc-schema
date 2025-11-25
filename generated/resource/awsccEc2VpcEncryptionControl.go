package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcEncryptionControl = `{
  "block": {
    "attributes": {
      "egress_only_internet_gateway_exclusion_input": {
        "computed": true,
        "description": "Used to enable or disable EIGW exclusion",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "elastic_file_system_exclusion_input": {
        "computed": true,
        "description": "Used to enable or disable EFS exclusion",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "internet_gateway_exclusion_input": {
        "computed": true,
        "description": "Used to enable or disable IGW exclusion",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lambda_exclusion_input": {
        "computed": true,
        "description": "Used to enable or disable Lambda exclusion",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mode": {
        "computed": true,
        "description": "The VPC encryption control mode, either monitor or enforce.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_gateway_exclusion_input": {
        "computed": true,
        "description": "Used to enable or disable Nat gateway exclusion",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_exclusions": {
        "computed": true,
        "description": "Enumerates the states of all the VPC encryption control resource exclusions",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "egress_only_internet_gateway": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "state": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "state_message": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "elastic_file_system": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "state": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "state_message": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "internet_gateway": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "state": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "state_message": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "lambda": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "state": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "state_message": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "nat_gateway": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "state": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "state_message": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "virtual_private_gateway": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "state": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "state_message": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vpc_lattice": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "state": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "state_message": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vpc_peering": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "state": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "state_message": {
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
      "state": {
        "computed": true,
        "description": "The current state of the VPC encryption control.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_message": {
        "computed": true,
        "description": "Provides additional context on the state of the VPC encryption control.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to assign to the VPC encryption control.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "virtual_private_gateway_exclusion_input": {
        "computed": true,
        "description": "Used to enable or disable VGW exclusion",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_encryption_control_id": {
        "computed": true,
        "description": "The VPC encryption control resource id.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The VPC on which this VPC encryption control is applied.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_lattice_exclusion_input": {
        "computed": true,
        "description": "Used to enable or disable Vpc Lattice exclusion",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_peering_exclusion_input": {
        "computed": true,
        "description": "Used to enable or disable VPC peering exclusion",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::VPCEncryptionControl",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcEncryptionControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcEncryptionControl), &result)
	return &result
}
