package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontResponseHeadersPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "response_headers_policy_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cors_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_control_allow_credentials": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "access_control_allow_headers": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "items": {
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
                  "access_control_allow_methods": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "items": {
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
                  "access_control_allow_origins": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "items": {
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
                  "access_control_expose_headers": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "items": {
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
                    "optional": true
                  },
                  "access_control_max_age_sec": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "origin_override": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "custom_headers_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "items": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "header": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "override": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "remove_headers_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "items": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "header": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "security_headers_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "content_security_policy": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_security_policy": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "override": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "content_type_options": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "override": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "frame_options": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "frame_option": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "override": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "referrer_policy": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "override": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "referrer_policy": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "strict_transport_security": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_control_max_age_sec": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "include_subdomains": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "override": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "preload": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "xss_protection": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "mode_block": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "override": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "protection": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "report_uri": {
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
            "server_timing_headers_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "sampling_rate": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Resource Type definition for AWS::CloudFront::ResponseHeadersPolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontResponseHeadersPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontResponseHeadersPolicy), &result)
	return &result
}
