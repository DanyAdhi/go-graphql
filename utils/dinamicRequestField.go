package utils

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func GetRequestedFields(ctx context.Context) []string {
	fields := graphql.CollectFieldsCtx(ctx, nil)

	var selectField []string
	for _, field := range fields {
		dataAppend := mapFields[field.Name]
		selectField = append(selectField, dataAppend)
	}

	return selectField
}

func MapFieldsOrder(fields []string) []string {
	var result []string
	for _, field := range fields {
		dataAppend := mapFields[field]
		result = append(result, dataAppend)
	}
	return result
}

var mapFields = map[string]string{
	`purchaseRequestNumber`: `purchase_request_number`,
	`purchaseOrderNumber`:   `purchase_order_number`,
	`salesOrderNumber`:      `sales_order_number`,
	`bastNumber`:            `bast_number`,
	`invoiceNumber`:         `invoice_number`,
	`statusId`:              `status_id`,
	`projectType`:           `project_type`,
	`invoiceDate`:           `invoice_date`,
}
