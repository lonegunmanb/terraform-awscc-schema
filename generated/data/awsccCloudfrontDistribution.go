package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontDistribution = `{
  "block": {
    "attributes": {
      "distribution_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aliases": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "cache_behaviors": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_methods": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cache_policy_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "cached_methods": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "compress": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "default_ttl": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "field_level_encryption_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "forwarded_values": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cookies": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "forward": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "whitelisted_names": {
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
                        "headers": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "query_string": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "query_string_cache_keys": {
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
                  "function_associations": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "function_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "lambda_function_associations": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "include_body": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "lambda_function_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "max_ttl": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min_ttl": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "origin_request_policy_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "path_pattern": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "realtime_log_config_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "response_headers_policy_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "smooth_streaming": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "target_origin_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "trusted_key_groups": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "trusted_signers": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "viewer_protocol_policy": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "cnames": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "comment": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "continuous_deployment_policy_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "custom_error_responses": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "error_caching_min_ttl": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "error_code": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "response_code": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "response_page_path": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "custom_origin": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dns_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "http_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "https_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "origin_protocol_policy": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "origin_ssl_protocols": {
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
            "default_cache_behavior": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_methods": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cache_policy_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "cached_methods": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "compress": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "default_ttl": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "field_level_encryption_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "forwarded_values": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cookies": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "forward": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "whitelisted_names": {
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
                        "headers": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "query_string": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "query_string_cache_keys": {
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
                  "function_associations": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "function_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "lambda_function_associations": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "include_body": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "lambda_function_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "max_ttl": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min_ttl": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "origin_request_policy_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "realtime_log_config_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "response_headers_policy_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "smooth_streaming": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "target_origin_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "trusted_key_groups": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "trusted_signers": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "viewer_protocol_policy": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "default_root_object": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "http_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ipv6_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "logging": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "include_cookies": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "origin_groups": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "items": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "failover_criteria": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "status_codes": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "items": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "number"
                                      ]
                                    },
                                    "quantity": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "members": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "items": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "origin_id": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "quantity": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "quantity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "origins": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connection_attempts": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "connection_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "custom_origin_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "http_port": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "https_port": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "origin_keepalive_timeout": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "origin_protocol_policy": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "origin_read_timeout": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "origin_ssl_protocols": {
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
                  "domain_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "origin_access_control_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "origin_custom_headers": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "header_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "header_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "origin_path": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "origin_shield": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "origin_shield_region": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "s3_origin_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "origin_access_identity": {
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
            },
            "price_class": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "restrictions": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "geo_restriction": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "locations": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "restriction_type": {
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
            "s3_origin": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dns_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "origin_access_identity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "staging": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "viewer_certificate": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "acm_certificate_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "cloudfront_default_certificate": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "iam_certificate_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "minimum_protocol_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ssl_support_method": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "web_acl_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
    "description": "Data Source schema for AWS::CloudFront::Distribution",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontDistributionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontDistribution), &result)
	return &result
}
