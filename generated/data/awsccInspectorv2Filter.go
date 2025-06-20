package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccInspectorv2Filter = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Findings filter ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Findings filter description.",
        "description_kind": "plain",
        "type": "string"
      },
      "filter_action": {
        "computed": true,
        "description": "Findings filter action.",
        "description_kind": "plain",
        "type": "string"
      },
      "filter_criteria": {
        "computed": true,
        "description": "Findings filter criteria.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_account_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "code_vulnerability_detector_name": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "code_vulnerability_detector_tags": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "code_vulnerability_file_path": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "component_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "component_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "ec_2_instance_image_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "ec_2_instance_subnet_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "ec_2_instance_vpc_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "ecr_image_architecture": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "ecr_image_hash": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "ecr_image_pushed_at": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "start_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "ecr_image_registry": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "ecr_image_repository_name": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "ecr_image_tags": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "epss_score": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lower_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "upper_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "exploit_available": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "finding_arn": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "finding_status": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "finding_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "first_observed_at": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "start_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "fix_available": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "inspector_score": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lower_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "upper_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "lambda_function_execution_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "lambda_function_last_modified_at": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "start_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "lambda_function_layers": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "lambda_function_name": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "lambda_function_runtime": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "last_observed_at": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "start_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "network_protocol": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "port_range": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "begin_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "end_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "related_vulnerabilities": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "resource_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "resource_tags": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
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
            "resource_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "severity": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "title": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "start_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "vendor_severity": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "vulnerability_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "vulnerability_source": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
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
            "vulnerable_packages": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "architecture": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "epoch": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "lower_inclusive": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "upper_inclusive": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "file_path": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "release": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "source_lambda_layer_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "source_layer_hash": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "version": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison": {
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
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
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
      "name": {
        "computed": true,
        "description": "Findings filter name.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::InspectorV2::Filter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccInspectorv2FilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInspectorv2Filter), &result)
	return &result
}
