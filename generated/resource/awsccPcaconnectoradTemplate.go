package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectoradTemplate = `{
  "block": {
    "attributes": {
      "connector_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "definition": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "template_v2": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_validity": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "renewal_period": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "period_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        },
                        "validity_period": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "period_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "enrollment_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enable_key_reuse_on_nt_token_keyset_storage_full": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "include_symmetric_algorithms": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "no_security_extension": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "remove_invalid_certificate_from_personal_store": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "user_interaction_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "extensions": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "application_policies": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "critical": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "policies": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "policy_object_identifier": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "policy_type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
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
                        "key_usage": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "critical": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "usage_flags": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "data_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "digital_signature": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "key_agreement": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "key_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "non_repudiation": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "required": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "general_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auto_enrollment": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "machine_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "private_key_attributes": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "crypto_providers": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "key_spec": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "minimal_key_length": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "private_key_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_version": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "exportable_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "strong_key_protection_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "subject_name_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "require_common_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_directory_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_dns_as_cn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_directory_guid": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_domain_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_spn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_upn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "superseded_templates": {
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
            "template_v3": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_validity": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "renewal_period": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "period_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        },
                        "validity_period": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "period_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "enrollment_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enable_key_reuse_on_nt_token_keyset_storage_full": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "include_symmetric_algorithms": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "no_security_extension": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "remove_invalid_certificate_from_personal_store": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "user_interaction_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "extensions": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "application_policies": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "critical": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "policies": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "policy_object_identifier": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "policy_type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
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
                        "key_usage": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "critical": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "usage_flags": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "data_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "digital_signature": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "key_agreement": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "key_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "non_repudiation": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "required": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "general_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auto_enrollment": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "machine_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "hash_algorithm": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "private_key_attributes": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "algorithm": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "crypto_providers": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "key_spec": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "key_usage_property": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "property_flags": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "decrypt": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "key_agreement": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "sign": {
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
                              "property_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        },
                        "minimal_key_length": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "private_key_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_version": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "exportable_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_alternate_signature_algorithm": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "strong_key_protection_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "subject_name_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "require_common_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_directory_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_dns_as_cn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_directory_guid": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_domain_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_spn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_upn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "superseded_templates": {
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
            "template_v4": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_validity": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "renewal_period": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "period_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        },
                        "validity_period": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "period_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "enrollment_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enable_key_reuse_on_nt_token_keyset_storage_full": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "include_symmetric_algorithms": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "no_security_extension": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "remove_invalid_certificate_from_personal_store": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "user_interaction_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "extensions": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "application_policies": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "critical": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "policies": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "policy_object_identifier": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "policy_type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
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
                        "key_usage": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "critical": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "usage_flags": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "data_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "digital_signature": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "key_agreement": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "key_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "non_repudiation": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "required": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "general_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auto_enrollment": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "machine_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "hash_algorithm": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "private_key_attributes": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "algorithm": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "crypto_providers": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "key_spec": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "key_usage_property": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "property_flags": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "decrypt": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "key_agreement": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "sign": {
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
                              "property_type": {
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
                        "minimal_key_length": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "private_key_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_version": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "exportable_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_alternate_signature_algorithm": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_same_key_renewal": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "strong_key_protection_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "use_legacy_provider": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "subject_name_flags": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "require_common_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_directory_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_dns_as_cn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_directory_guid": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_domain_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_spn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "san_require_upn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "superseded_templates": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reenroll_all_certificate_holders": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "template_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Represents a template that defines certificate configurations, both for issuance and client handling",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPcaconnectoradTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectoradTemplate), &result)
	return &result
}
