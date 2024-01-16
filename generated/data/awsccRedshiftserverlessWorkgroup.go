package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftserverlessWorkgroup = `{
  "block": {
    "attributes": {
      "base_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "config_parameters": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "parameter_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "parameter_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "enhanced_vpc_routing": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "publicly_accessible": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
      "workgroup": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "config_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "parameter_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "parameter_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "creation_date": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "endpoint": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "address": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "vpc_endpoints": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "network_interfaces": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "availability_zone": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
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
                              "subnet_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "vpc_endpoint_id": {
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
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "enhanced_vpc_routing": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "namespace_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "publicly_accessible": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "workgroup_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "workgroup_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "workgroup_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "workgroup_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::RedshiftServerless::Workgroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRedshiftserverlessWorkgroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftserverlessWorkgroup), &result)
	return &result
}
