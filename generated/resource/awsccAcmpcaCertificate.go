package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAcmpcaCertificate = `{
  "block": {
    "attributes": {
      "api_passthrough": {
        "computed": true,
        "description": "Specifies X.509 certificate information to be included in the issued certificate. An ` + "`" + `` + "`" + `APIPassthrough` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `APICSRPassthrough` + "`" + `` + "`" + ` template variant must be selected, or else this parameter is ignored.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "extensions": {
              "computed": true,
              "description": "Specifies X.509 extension information for a certificate.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_policies": {
                    "computed": true,
                    "description": "Contains a sequence of one or more policy information terms, each of which consists of an object identifier (OID) and optional qualifiers. For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).\n In an end-entity certificate, these terms indicate the policy under which the certificate was issued and the purposes for which it may be used. In a CA certificate, these terms limit the set of policies for certification paths that include this certificate.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cert_policy_id": {
                          "description": "Specifies the object identifier (OID) of the certificate policy under which the certificate was issued. For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "policy_qualifiers": {
                          "computed": true,
                          "description": "Modifies the given ` + "`" + `` + "`" + `CertPolicyId` + "`" + `` + "`" + ` with a qualifier. AWS Private CA supports the certification practice statement (CPS) qualifier.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "policy_qualifier_id": {
                                "description": "Identifies the qualifier modifying a ` + "`" + `` + "`" + `CertPolicyId` + "`" + `` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "qualifier": {
                                "description": "Defines the qualifier type. AWS Private CA supports the use of a URI for a CPS qualifier in this field.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "cps_uri": {
                                      "description": "Contains a pointer to a certification practice statement (CPS) published by the CA.",
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
                            "nesting_mode": "list"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "custom_extensions": {
                    "computed": true,
                    "description": "Contains a sequence of one or more X.509 extensions, each of which consists of an object identifier (OID), a base64-encoded value, and the critical flag. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29)",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "critical": {
                          "computed": true,
                          "description": "Specifies the critical flag of the X.509 extension.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "object_identifier": {
                          "description": "Specifies the object identifier (OID) of the X.509 extension. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29)",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Specifies the base64-encoded value of the X.509 extension.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "extended_key_usage": {
                    "computed": true,
                    "description": "Specifies additional purposes for which the certified public key may be used other than basic purposes indicated in the ` + "`" + `` + "`" + `KeyUsage` + "`" + `` + "`" + ` extension.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "extended_key_usage_object_identifier": {
                          "computed": true,
                          "description": "Specifies a custom ` + "`" + `` + "`" + `ExtendedKeyUsage` + "`" + `` + "`" + ` with an object identifier (OID).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "extended_key_usage_type": {
                          "computed": true,
                          "description": "Specifies a standard ` + "`" + `` + "`" + `ExtendedKeyUsage` + "`" + `` + "`" + ` as defined as in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "key_usage": {
                    "computed": true,
                    "description": "Defines one or more purposes for which the key contained in the certificate can be used. Default value for each option is false.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "crl_sign": {
                          "computed": true,
                          "description": "Key can be used to sign CRLs.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "data_encipherment": {
                          "computed": true,
                          "description": "Key can be used to decipher data.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "decipher_only": {
                          "computed": true,
                          "description": "Key can be used only to decipher data.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "digital_signature": {
                          "computed": true,
                          "description": "Key can be used for digital signing.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "encipher_only": {
                          "computed": true,
                          "description": "Key can be used only to encipher data.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "key_agreement": {
                          "computed": true,
                          "description": "Key can be used in a key-agreement protocol.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "key_cert_sign": {
                          "computed": true,
                          "description": "Key can be used to sign certificates.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "key_encipherment": {
                          "computed": true,
                          "description": "Key can be used to encipher data.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "non_repudiation": {
                          "computed": true,
                          "description": "Key can be used for non-repudiation.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "subject_alternative_names": {
                    "computed": true,
                    "description": "The subject alternative name extension allows identities to be bound to the subject of the certificate. These identities may be included in addition to or in place of the identity in the subject field of the certificate.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "directory_name": {
                          "computed": true,
                          "description": "Contains information about the certificate subject. The certificate can be one issued by your private certificate authority (CA) or it can be your private CA certificate. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate. The DN must be unique for each entity, but your private CA can issue more than one certificate with the same DN to the same entity.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "common_name": {
                                "computed": true,
                                "description": "For CA and end-entity certificates in a private PKI, the common name (CN) can be any string within the length limit.\n Note: In publicly trusted certificates, the common name must be a fully qualified domain name (FQDN) associated with the certificate subject.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "country": {
                                "computed": true,
                                "description": "Two-digit code that specifies the country in which the certificate subject located.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "custom_attributes": {
                                "computed": true,
                                "description": "Contains a sequence of one or more X.500 relative distinguished names (RDNs), each of which consists of an object identifier (OID) and a value. For more information, see NIST?s definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).\n  Custom attributes cannot be used in combination with standard attributes.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "object_identifier": {
                                      "description": "Specifies the object identifier (OID) of the attribute type of the relative distinguished name (RDN).",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description": "Specifies the attribute value of relative distinguished name (RDN).",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "optional": true
                              },
                              "distinguished_name_qualifier": {
                                "computed": true,
                                "description": "Disambiguating information for the certificate subject.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "generation_qualifier": {
                                "computed": true,
                                "description": "Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "given_name": {
                                "computed": true,
                                "description": "First name.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "initials": {
                                "computed": true,
                                "description": "Concatenation that typically contains the first letter of the *GivenName*, the first letter of the middle name if one exists, and the first letter of the *Surname*.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "locality": {
                                "computed": true,
                                "description": "The locality (such as a city or town) in which the certificate subject is located.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "organization": {
                                "computed": true,
                                "description": "Legal name of the organization with which the certificate subject is affiliated.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "organizational_unit": {
                                "computed": true,
                                "description": "A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "pseudonym": {
                                "computed": true,
                                "description": "Typically a shortened version of a longer *GivenName*. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "serial_number": {
                                "computed": true,
                                "description": "The certificate serial number.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "state": {
                                "computed": true,
                                "description": "State in which the subject of the certificate is located.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "surname": {
                                "computed": true,
                                "description": "Family name. In the US and the UK, for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "title": {
                                "computed": true,
                                "description": "A title such as Mr. or Ms., which is pre-pended to the name to refer formally to the certificate subject.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "dns_name": {
                          "computed": true,
                          "description": "Represents ` + "`" + `` + "`" + `GeneralName` + "`" + `` + "`" + ` as a DNS name.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "edi_party_name": {
                          "computed": true,
                          "description": "Represents ` + "`" + `` + "`" + `GeneralName` + "`" + `` + "`" + ` as an ` + "`" + `` + "`" + `EdiPartyName` + "`" + `` + "`" + ` object.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name_assigner": {
                                "description": "Specifies the name assigner.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "party_name": {
                                "description": "Specifies the party name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "ip_address": {
                          "computed": true,
                          "description": "Represents ` + "`" + `` + "`" + `GeneralName` + "`" + `` + "`" + ` as an IPv4 or IPv6 address.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "other_name": {
                          "computed": true,
                          "description": "Represents ` + "`" + `` + "`" + `GeneralName` + "`" + `` + "`" + ` using an ` + "`" + `` + "`" + `OtherName` + "`" + `` + "`" + ` object.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type_id": {
                                "description": "Specifies an OID.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "Specifies an OID value.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "registered_id": {
                          "computed": true,
                          "description": "Represents ` + "`" + `` + "`" + `GeneralName` + "`" + `` + "`" + ` as an object identifier (OID).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "rfc_822_name": {
                          "computed": true,
                          "description": "Represents ` + "`" + `` + "`" + `GeneralName` + "`" + `` + "`" + ` as an [RFC 822](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc822) email address.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "uniform_resource_identifier": {
                          "computed": true,
                          "description": "Represents ` + "`" + `` + "`" + `GeneralName` + "`" + `` + "`" + ` as a URI.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "subject": {
              "computed": true,
              "description": "Contains information about the certificate subject. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "common_name": {
                    "computed": true,
                    "description": "For CA and end-entity certificates in a private PKI, the common name (CN) can be any string within the length limit.\n Note: In publicly trusted certificates, the common name must be a fully qualified domain name (FQDN) associated with the certificate subject.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "country": {
                    "computed": true,
                    "description": "Two-digit code that specifies the country in which the certificate subject located.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "custom_attributes": {
                    "computed": true,
                    "description": "Contains a sequence of one or more X.500 relative distinguished names (RDNs), each of which consists of an object identifier (OID) and a value. For more information, see NIST?s definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).\n  Custom attributes cannot be used in combination with standard attributes.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object_identifier": {
                          "description": "Specifies the object identifier (OID) of the attribute type of the relative distinguished name (RDN).",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Specifies the attribute value of relative distinguished name (RDN).",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "distinguished_name_qualifier": {
                    "computed": true,
                    "description": "Disambiguating information for the certificate subject.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "generation_qualifier": {
                    "computed": true,
                    "description": "Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "given_name": {
                    "computed": true,
                    "description": "First name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "initials": {
                    "computed": true,
                    "description": "Concatenation that typically contains the first letter of the *GivenName*, the first letter of the middle name if one exists, and the first letter of the *Surname*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "locality": {
                    "computed": true,
                    "description": "The locality (such as a city or town) in which the certificate subject is located.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "organization": {
                    "computed": true,
                    "description": "Legal name of the organization with which the certificate subject is affiliated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "organizational_unit": {
                    "computed": true,
                    "description": "A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pseudonym": {
                    "computed": true,
                    "description": "Typically a shortened version of a longer *GivenName*. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "serial_number": {
                    "computed": true,
                    "description": "The certificate serial number.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "state": {
                    "computed": true,
                    "description": "State in which the subject of the certificate is located.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "surname": {
                    "computed": true,
                    "description": "Family name. In the US and the UK, for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "title": {
                    "computed": true,
                    "description": "A title such as Mr. or Ms., which is pre-pended to the name to refer formally to the certificate subject.",
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
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_authority_arn": {
        "description": "The Amazon Resource Name (ARN) for the private CA issues the certificate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_signing_request": {
        "description": "The certificate signing request (CSR) for the certificate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "signing_algorithm": {
        "description": "The name of the algorithm that will be used to sign the certificate to be issued. \n This parameter should not be confused with the ` + "`" + `` + "`" + `SigningAlgorithm` + "`" + `` + "`" + ` parameter used to sign a CSR in the ` + "`" + `` + "`" + `CreateCertificateAuthority` + "`" + `` + "`" + ` action.\n  The specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_arn": {
        "computed": true,
        "description": "Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, PCAshort defaults to the ` + "`" + `` + "`" + `EndEntityCertificate/V1` + "`" + `` + "`" + ` template. For more information about PCAshort templates, see [Using Templates](https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "validity": {
        "description": "The period of time during which the certificate will be valid.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "description": "Specifies whether the ` + "`" + `` + "`" + `Value` + "`" + `` + "`" + ` parameter represents days, months, or years.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "A long integer interpreted according to the value of ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + `, below.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "validity_not_before": {
        "computed": true,
        "description": "Information describing the start of the validity period of the certificate. This parameter sets the ?Not Before\" date for the certificate.\n By default, when issuing a certificate, PCAshort sets the \"Not Before\" date to the issuance time minus 60 minutes. This compensates for clock inconsistencies across computer systems. The ` + "`" + `` + "`" + `ValidityNotBefore` + "`" + `` + "`" + ` parameter can be used to customize the ?Not Before? value. \n Unlike the ` + "`" + `` + "`" + `Validity` + "`" + `` + "`" + ` parameter, the ` + "`" + `` + "`" + `ValidityNotBefore` + "`" + `` + "`" + ` parameter is optional.\n The ` + "`" + `` + "`" + `ValidityNotBefore` + "`" + `` + "`" + ` value is expressed as an explicit date and time, using the ` + "`" + `` + "`" + `Validity` + "`" + `` + "`" + ` type value ` + "`" + `` + "`" + `ABSOLUTE` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "description": "Specifies whether the ` + "`" + `` + "`" + `Value` + "`" + `` + "`" + ` parameter represents days, months, or years.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "A long integer interpreted according to the value of ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + `, below.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ACMPCA::Certificate` + "`" + `` + "`" + ` resource is used to issue a certificate using your private certificate authority. For more information, see the [IssueCertificate](https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html) action.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAcmpcaCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAcmpcaCertificate), &result)
	return &result
}
