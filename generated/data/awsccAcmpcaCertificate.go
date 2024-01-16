package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAcmpcaCertificate = `{
  "block": {
    "attributes": {
      "api_passthrough": {
        "computed": true,
        "description": "These are fields to be overridden in a certificate at the time of issuance. These requires an API_Passthrough template be used or they will be ignored.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "extensions": {
              "computed": true,
              "description": "Structure that contains X.500 extensions for a Certificate.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_policies": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cert_policy_id": {
                          "computed": true,
                          "description": "String that contains X.509 ObjectIdentifier information.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "policy_qualifiers": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "policy_qualifier_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "qualifier": {
                                "computed": true,
                                "description": "Structure that contains a X.509 policy qualifier.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "cps_uri": {
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
                      "nesting_mode": "list"
                    }
                  },
                  "custom_extensions": {
                    "computed": true,
                    "description": "Array of X.509 extensions for a certificate.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "critical": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "object_identifier": {
                          "computed": true,
                          "description": "String that contains X.509 ObjectIdentifier information.",
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
                  "extended_key_usage": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "extended_key_usage_object_identifier": {
                          "computed": true,
                          "description": "String that contains X.509 ObjectIdentifier information.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "extended_key_usage_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "key_usage": {
                    "computed": true,
                    "description": "Structure that contains X.509 KeyUsage information.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "crl_sign": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "data_encipherment": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "decipher_only": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "digital_signature": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "encipher_only": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "key_agreement": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "key_cert_sign": {
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
                  },
                  "subject_alternative_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "directory_name": {
                          "computed": true,
                          "description": "Structure that contains X.500 distinguished name information.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "common_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "country": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "custom_attributes": {
                                "computed": true,
                                "description": "Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "object_identifier": {
                                      "computed": true,
                                      "description": "String that contains X.509 ObjectIdentifier information.",
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
                              "distinguished_name_qualifier": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "generation_qualifier": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "given_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "initials": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "locality": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "organization": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "organizational_unit": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "pseudonym": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "serial_number": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "state": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "surname": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "title": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "dns_name": {
                          "computed": true,
                          "description": "String that contains X.509 DnsName information.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "edi_party_name": {
                          "computed": true,
                          "description": "Structure that contains X.509 EdiPartyName information.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name_assigner": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "party_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "ip_address": {
                          "computed": true,
                          "description": "String that contains X.509 IpAddress information.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "other_name": {
                          "computed": true,
                          "description": "Structure that contains X.509 OtherName information.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type_id": {
                                "computed": true,
                                "description": "String that contains X.509 ObjectIdentifier information.",
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
                        "registered_id": {
                          "computed": true,
                          "description": "String that contains X.509 ObjectIdentifier information.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "rfc_822_name": {
                          "computed": true,
                          "description": "String that contains X.509 Rfc822Name information.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "uniform_resource_identifier": {
                          "computed": true,
                          "description": "String that contains X.509 UniformResourceIdentifier information.",
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
            "subject": {
              "computed": true,
              "description": "Structure that contains X.500 distinguished name information.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "common_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "country": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "custom_attributes": {
                    "computed": true,
                    "description": "Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object_identifier": {
                          "computed": true,
                          "description": "String that contains X.509 ObjectIdentifier information.",
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
                  "distinguished_name_qualifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "generation_qualifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "given_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "initials": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "locality": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "organization": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "organizational_unit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "pseudonym": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "serial_number": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "state": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "surname": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "title": {
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
      "arn": {
        "computed": true,
        "description": "The ARN of the issued certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate": {
        "computed": true,
        "description": "The issued certificate in base 64 PEM-encoded format.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_authority_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the private CA to issue the certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_signing_request": {
        "computed": true,
        "description": "The certificate signing request (CSR) for the Certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "signing_algorithm": {
        "computed": true,
        "description": "The name of the algorithm that will be used to sign the Certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_arn": {
        "computed": true,
        "description": "Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, ACM Private CA defaults to the EndEntityCertificate/V1 template.",
        "description_kind": "plain",
        "type": "string"
      },
      "validity": {
        "computed": true,
        "description": "The time before which the Certificate will be valid.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "validity_not_before": {
        "computed": true,
        "description": "The time after which the Certificate will be valid.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::ACMPCA::Certificate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAcmpcaCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAcmpcaCertificate), &result)
	return &result
}
