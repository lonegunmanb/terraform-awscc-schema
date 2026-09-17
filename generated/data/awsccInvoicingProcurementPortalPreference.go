package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccInvoicingProcurementPortalPreference = `{
  "block": {
    "attributes": {
      "aws_account_id": {
        "computed": true,
        "description": "The AWS account ID associated with this procurement portal preference.",
        "description_kind": "plain",
        "type": "string"
      },
      "buyer_domain": {
        "computed": true,
        "description": "The domain identifier for the buyer in the procurement portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "buyer_identifier": {
        "computed": true,
        "description": "The unique identifier for the buyer in the procurement portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "contacts": {
        "computed": true,
        "description": "List of contact information for portal administrators and technical contacts.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "email": {
              "computed": true,
              "description": "The email address of the contact person or role.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the contact person or role.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "create_date": {
        "computed": true,
        "description": "The date and time when the procurement portal preference was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "einvoice_delivery_enabled": {
        "computed": true,
        "description": "Indicates whether e-invoice delivery is enabled for this procurement portal preference.",
        "description_kind": "plain",
        "type": "bool"
      },
      "einvoice_delivery_preference": {
        "computed": true,
        "description": "Specifies the preferences for e-invoice delivery.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connection_testing_method": {
              "computed": true,
              "description": "The method to use for testing the connection to the procurement portal.",
              "description_kind": "plain",
              "type": "string"
            },
            "einvoice_delivery_activation_date": {
              "computed": true,
              "description": "The ISO 8601 date-time when e-invoice delivery should be activated.",
              "description_kind": "plain",
              "type": "string"
            },
            "einvoice_delivery_attachment_types": {
              "computed": true,
              "description": "The types of attachments to include with the e-invoice delivery.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "einvoice_delivery_document_types": {
              "computed": true,
              "description": "The types of e-invoice documents to be delivered.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "protocol": {
              "computed": true,
              "description": "The communication protocol to use for e-invoice delivery.",
              "description_kind": "plain",
              "type": "string"
            },
            "purchase_order_data_sources": {
              "computed": true,
              "description": "The sources of purchase order data.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "einvoice_delivery_document_type": {
                    "computed": true,
                    "description": "The type of e-invoice document that requires purchase order data.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "purchase_order_data_source_type": {
                    "computed": true,
                    "description": "The type of source for purchase order data.",
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
      "einvoice_delivery_preference_status": {
        "computed": true,
        "description": "The current status of the e-invoice delivery preference.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_update_date": {
        "computed": true,
        "description": "The date and time when the procurement portal preference was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "procurement_portal_instance_endpoint": {
        "computed": true,
        "description": "The endpoint URL where e-invoices are delivered to the procurement portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "procurement_portal_name": {
        "computed": true,
        "description": "The name of the procurement portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "procurement_portal_preference_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the procurement portal preference.",
        "description_kind": "plain",
        "type": "string"
      },
      "procurement_portal_shared_secret": {
        "computed": true,
        "description": "The shared secret or authentication credential used for secure communication with the procurement portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "purchase_order_retrieval_enabled": {
        "computed": true,
        "description": "Indicates whether purchase order retrieval is enabled for this procurement portal preference.",
        "description_kind": "plain",
        "type": "bool"
      },
      "purchase_order_retrieval_endpoint": {
        "computed": true,
        "description": "The endpoint URL used for retrieving purchase orders from the procurement portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "purchase_order_retrieval_preference_status": {
        "computed": true,
        "description": "The current status of the purchase order retrieval preference.",
        "description_kind": "plain",
        "type": "string"
      },
      "selector": {
        "computed": true,
        "description": "Specifies criteria for selecting which invoices should be processed.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "invoice_unit_arns": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of invoice unit identifiers to which this preference applies.",
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
      "supplier_domain": {
        "computed": true,
        "description": "The domain identifier for the supplier in the procurement portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "supplier_identifier": {
        "computed": true,
        "description": "The unique identifier for the supplier in the procurement portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with this procurement portal preference.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "test_env_preference": {
        "computed": true,
        "description": "Configuration settings for the test environment of the procurement portal.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "buyer_domain": {
              "computed": true,
              "description": "The domain identifier for the buyer in the test environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "buyer_identifier": {
              "computed": true,
              "description": "The unique identifier for the buyer in the test environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "procurement_portal_instance_endpoint": {
              "computed": true,
              "description": "The endpoint URL for e-invoice delivery in the test environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "procurement_portal_shared_secret": {
              "computed": true,
              "description": "The shared secret for secure communication in the test environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "supplier_domain": {
              "computed": true,
              "description": "The domain identifier for the supplier in the test environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "supplier_identifier": {
              "computed": true,
              "description": "The unique identifier for the supplier in the test environment.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "version": {
        "computed": true,
        "description": "The version number of the procurement portal preference configuration.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::Invoicing::ProcurementPortalPreference",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccInvoicingProcurementPortalPreferenceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInvoicingProcurementPortalPreference), &result)
	return &result
}
