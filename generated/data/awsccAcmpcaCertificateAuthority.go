package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAcmpcaCertificateAuthority = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the certificate authority.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_signing_request": {
        "computed": true,
        "description": "The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "csr_extensions": {
        "computed": true,
        "description": "Structure that contains CSR pass through extension information used by the CreateCertificateAuthority action.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
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
            "subject_information_access": {
              "computed": true,
              "description": "Array of X.509 AccessDescription.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_location": {
                    "computed": true,
                    "description": "Structure that contains X.509 GeneralName information. Assign one and ONLY one field.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "directory_name": {
                          "computed": true,
                          "description": "Structure that contains X.500 distinguished name information for your CA.",
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
                      "nesting_mode": "single"
                    }
                  },
                  "access_method": {
                    "computed": true,
                    "description": "Structure that contains X.509 AccessMethod information. Assign one and ONLY one field.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_method_type": {
                          "computed": true,
                          "description": "Pre-defined enum string for X.509 AccessMethod ObjectIdentifiers.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "custom_object_identifier": {
                          "computed": true,
                          "description": "String that contains X.509 ObjectIdentifier information.",
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
      "key_algorithm": {
        "computed": true,
        "description": "Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "key_storage_security_standard": {
        "computed": true,
        "description": "KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.",
        "description_kind": "plain",
        "type": "string"
      },
      "revocation_configuration": {
        "computed": true,
        "description": "Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "crl_configuration": {
              "computed": true,
              "description": "Your certificate authority can create and maintain a certificate revocation list (CRL). A CRL contains information about certificates that have been revoked.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "crl_distribution_point_extension_configuration": {
                    "computed": true,
                    "description": "Configures the default behavior of the CRL Distribution Point extension for certificates issued by your certificate authority",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "omit_extension": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "crl_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "custom_cname": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "custom_path": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "expiration_in_days": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "s3_bucket_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_object_acl": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ocsp_configuration": {
              "computed": true,
              "description": "Helps to configure online certificate status protocol (OCSP) responder for your certificate authority",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "ocsp_custom_cname": {
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
      "signing_algorithm": {
        "computed": true,
        "description": "Algorithm your CA uses to sign certificate requests.",
        "description_kind": "plain",
        "type": "string"
      },
      "subject": {
        "computed": true,
        "description": "Structure that contains X.500 distinguished name information for your CA.",
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
      "type": {
        "computed": true,
        "description": "The type of the certificate authority.",
        "description_kind": "plain",
        "type": "string"
      },
      "usage_mode": {
        "computed": true,
        "description": "Usage mode of the ceritificate authority.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ACMPCA::CertificateAuthority",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAcmpcaCertificateAuthoritySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAcmpcaCertificateAuthority), &result)
	return &result
}
