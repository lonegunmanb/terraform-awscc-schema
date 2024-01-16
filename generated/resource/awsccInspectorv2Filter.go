package resource

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
        "optional": true,
        "type": "string"
      },
      "filter_action": {
        "description": "Findings filter action.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filter_criteria": {
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
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "component_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "component_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "ec_2_instance_image_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "ec_2_instance_subnet_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "ec_2_instance_vpc_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "ecr_image_architecture": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "ecr_image_hash": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "ecr_image_pushed_at": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "start_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "ecr_image_registry": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "ecr_image_repository_name": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "ecr_image_tags": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "finding_arn": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "finding_status": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "finding_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "first_observed_at": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "start_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "inspector_score": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lower_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "upper_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "last_observed_at": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "start_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "network_protocol": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "port_range": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "begin_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "end_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "related_vulnerabilities": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_tags": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
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
                "nesting_mode": "list"
              },
              "optional": true
            },
            "resource_type": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "severity": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "title": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "updated_at": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "start_inclusive": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "vendor_severity": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "vulnerability_id": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "vulnerability_source": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
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
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "epoch": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "lower_inclusive": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "upper_inclusive": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "release": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "source_layer_hash": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "version": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "comparison": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Findings filter name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Inspector Filter resource schema",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccInspectorv2FilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInspectorv2Filter), &result)
	return &result
}
