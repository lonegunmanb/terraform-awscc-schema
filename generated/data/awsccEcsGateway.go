package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsGateway = `{
  "block": {
    "attributes": {
      "access_type": {
        "computed": true,
        "description": "The access type of the gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description": "The configuration for the gateway.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "host_shared_alb": {
              "computed": true,
              "description": "The host shared ALB configuration for the gateway.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_logs": {
                    "computed": true,
                    "description": "The access log configuration for the gateway ALB. Access logging is enabled by the presence of a destination.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "destination": {
                          "computed": true,
                          "description": "The destination for gateway ALB access logs. Exactly one destination type must be specified.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "s3": {
                                "computed": true,
                                "description": "The S3 destination for gateway ALB access logs.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "bucket": {
                                      "computed": true,
                                      "description": "The S3 bucket for access logs.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "prefix": {
                                      "computed": true,
                                      "description": "The S3 prefix for access logs.",
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
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "current_alb_count": {
                    "computed": true,
                    "description": "The current number of ALBs provisioned for the gateway.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "http_desync": {
                    "computed": true,
                    "description": "The HTTP desync mitigation mode for the ALB.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "max_alb_count": {
                    "computed": true,
                    "description": "The maximum number of ALBs to provision for the gateway.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "network_configuration": {
                    "computed": true,
                    "description": "The network configuration for the gateway ALB.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "inbound_security_group": {
                          "computed": true,
                          "description": "The inbound security group for the ALB.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "ip_address_type": {
                          "computed": true,
                          "description": "The IP address type for the ALB.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "subnets": {
                          "computed": true,
                          "description": "The subnets for the ALB.",
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
                  "port": {
                    "computed": true,
                    "description": "The port on which the ALB listens.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "protocol": {
                    "computed": true,
                    "description": "The protocol for the ALB listener.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "scheme": {
                    "computed": true,
                    "description": "The scheme for the ALB.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "test_traffic_rules": {
                    "computed": true,
                    "description": "The test traffic rules for the gateway ALB.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "header": {
                          "computed": true,
                          "description": "The header configuration for test traffic routing.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "computed": true,
                                "description": "The name of the header.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "values": {
                                "computed": true,
                                "description": "The values for the header.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "web_acl_arn": {
                    "computed": true,
                    "description": "The ARN of the WAFv2 web ACL to associate with the ALB.",
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
      "created_at": {
        "computed": true,
        "description": "The timestamp when the gateway was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Whether deletion protection is enabled for the gateway.",
        "description_kind": "plain",
        "type": "bool"
      },
      "gateway_name": {
        "computed": true,
        "description": "The name of the gateway. Up to 255 letters (uppercase and lowercase), numbers, and hyphens are allowed.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "infrastructure_role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role that ECS uses to manage infrastructure for the gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "propagate_tags": {
        "computed": true,
        "description": "Specifies whether to propagate the tags from the gateway to other resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The metadata that you apply to the gateway to help you categorize and organize them.",
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
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the gateway was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ECS::Gateway",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcsGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsGateway), &result)
	return &result
}
