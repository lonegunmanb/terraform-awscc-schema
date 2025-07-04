package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCustomerprofilesSegmentDefinition = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time of this segment definition got created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the segment definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the segment definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The unique name of the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "segment_definition_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the segment definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "segment_definition_name": {
        "computed": true,
        "description": "The unique name of the segment definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "segment_groups": {
        "computed": true,
        "description": "An array that defines the set of segment criteria to evaluate when handling segment groups for the segment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "groups": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dimensions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "calculated_attributes": {
                          "computed": true,
                          "description": "One or more calculated attributes to use as criteria for the segment.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "condition_overrides": {
                                "computed": true,
                                "description": "Overrides the condition block within the original calculated attribute definition.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "range": {
                                      "computed": true,
                                      "description": "Defines the range to be applied to the calculated attribute definition.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "end": {
                                            "computed": true,
                                            "description": "The ending point for this overridden range. Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "start": {
                                            "computed": true,
                                            "description": "The starting point for this overridden range. Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "unit": {
                                            "computed": true,
                                            "description": "The unit to be applied to the range.",
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
                              "dimension_type": {
                                "computed": true,
                                "description": "The type of segment dimension to use.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "values": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "map"
                          }
                        },
                        "profile_attributes": {
                          "computed": true,
                          "description": "Specifies the dimension settings within profile attributes for a segment.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "account_number": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "additional_information": {
                                "computed": true,
                                "description": "Specifies criteria for a segment using extended-length string values.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "address": {
                                "computed": true,
                                "description": "The address based criteria for the segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "city": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "country": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "county": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "postal_code": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "province": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "state": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                              "attributes": {
                                "computed": true,
                                "description": "One or more custom attributes to use as criteria for the segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "nesting_mode": "map"
                                }
                              },
                              "billing_address": {
                                "computed": true,
                                "description": "The address based criteria for the segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "city": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "country": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "county": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "postal_code": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "province": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "state": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                              "birth_date": {
                                "computed": true,
                                "description": "Specifies date based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a date dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "business_email_address": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "business_name": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "business_phone_number": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "email_address": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "first_name": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "gender_string": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "home_phone_number": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "last_name": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "mailing_address": {
                                "computed": true,
                                "description": "The address based criteria for the segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "city": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "country": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "county": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "postal_code": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "province": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "state": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                              "middle_name": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "mobile_phone_number": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "party_type_string": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "personal_email_address": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "phone_number": {
                                "computed": true,
                                "description": "Specifies profile based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a string dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "profile_type": {
                                "computed": true,
                                "description": "Specifies profile type based criteria for a segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimension_type": {
                                      "computed": true,
                                      "description": "The type of segment dimension to use for a profile type dimension.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "values": {
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
                              "shipping_address": {
                                "computed": true,
                                "description": "The address based criteria for the segment.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "city": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "country": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "county": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "postal_code": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "province": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                                    "state": {
                                      "computed": true,
                                      "description": "Specifies profile based criteria for a segment.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimension_type": {
                                            "computed": true,
                                            "description": "The type of segment dimension to use for a string dimension.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "values": {
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
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "source_segments": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "segment_definition_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "source_type": {
                    "computed": true,
                    "description": "Specifies the operator on how to handle multiple groups within the same segment.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Specifies the operator on how to handle multiple groups within the same segment.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "include": {
              "computed": true,
              "description": "Specifies the operator on how to handle multiple groups within the same segment.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "The tags used to organize, track, or control access for this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::CustomerProfiles::SegmentDefinition",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCustomerprofilesSegmentDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCustomerprofilesSegmentDefinition), &result)
	return &result
}
