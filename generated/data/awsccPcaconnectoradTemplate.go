package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectoradTemplate = `{
  "block": {
    "attributes": {
      "connector_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "definition": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "template_v2": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_validity": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "renewal_period": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "period_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "validity_period": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "period_type": {
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
                  "enrollment_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enable_key_reuse_on_nt_token_keyset_storage_full": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "include_symmetric_algorithms": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "no_security_extension": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "remove_invalid_certificate_from_personal_store": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "user_interaction_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "extensions": {
                    "computed": true,
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
                                "type": "bool"
                              },
                              "policies": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "policy_object_identifier": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "policy_type": {
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
                        "key_usage": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "critical": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "usage_flags": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "data_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "digital_signature": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "key_agreement": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "key_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "non_repudiation": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
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
                  "general_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auto_enrollment": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "machine_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "private_key_attributes": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "crypto_providers": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "key_spec": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "minimal_key_length": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "private_key_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_version": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "exportable_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "strong_key_protection_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "subject_name_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "require_common_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_directory_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_dns_as_cn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_directory_guid": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_domain_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_spn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_upn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "superseded_templates": {
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
            "template_v3": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_validity": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "renewal_period": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "period_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "validity_period": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "period_type": {
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
                  "enrollment_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enable_key_reuse_on_nt_token_keyset_storage_full": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "include_symmetric_algorithms": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "no_security_extension": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "remove_invalid_certificate_from_personal_store": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "user_interaction_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "extensions": {
                    "computed": true,
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
                                "type": "bool"
                              },
                              "policies": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "policy_object_identifier": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "policy_type": {
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
                        "key_usage": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "critical": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "usage_flags": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "data_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "digital_signature": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "key_agreement": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "key_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "non_repudiation": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
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
                  "general_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auto_enrollment": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "machine_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "hash_algorithm": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "private_key_attributes": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "algorithm": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "crypto_providers": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "key_spec": {
                          "computed": true,
                          "description_kind": "plain",
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
                                      "type": "bool"
                                    },
                                    "key_agreement": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "sign": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "property_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "minimal_key_length": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "private_key_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_version": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "exportable_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_alternate_signature_algorithm": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "strong_key_protection_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "subject_name_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "require_common_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_directory_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_dns_as_cn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_directory_guid": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_domain_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_spn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_upn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "superseded_templates": {
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
            "template_v4": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_validity": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "renewal_period": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "period_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "validity_period": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "period": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "period_type": {
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
                  "enrollment_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enable_key_reuse_on_nt_token_keyset_storage_full": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "include_symmetric_algorithms": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "no_security_extension": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "remove_invalid_certificate_from_personal_store": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "user_interaction_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "extensions": {
                    "computed": true,
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
                                "type": "bool"
                              },
                              "policies": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "policy_object_identifier": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "policy_type": {
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
                        "key_usage": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "critical": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "usage_flags": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "data_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "digital_signature": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "key_agreement": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "key_encipherment": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "non_repudiation": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
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
                  "general_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auto_enrollment": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "machine_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "hash_algorithm": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "private_key_attributes": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "algorithm": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "crypto_providers": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "key_spec": {
                          "computed": true,
                          "description_kind": "plain",
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
                                      "type": "bool"
                                    },
                                    "key_agreement": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "sign": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "property_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "minimal_key_length": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "private_key_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_version": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "exportable_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_alternate_signature_algorithm": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_same_key_renewal": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "strong_key_protection_required": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "use_legacy_provider": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "subject_name_flags": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "require_common_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_directory_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_dns_as_cn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_directory_guid": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_domain_dns": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_email": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_spn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "san_require_upn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "superseded_templates": {
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
        "description_kind": "plain",
        "type": "string"
      },
      "reenroll_all_certificate_holders": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
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
    "description": "Data Source schema for AWS::PCAConnectorAD::Template",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPcaconnectoradTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectoradTemplate), &result)
	return &result
}
