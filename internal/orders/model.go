package orders

import (
	"database/sql"
	"time"
)

type Orders struct {
	ID                        int32          `json:"id"`
	PurchaseRequestNumber     sql.NullString `json:"purchaseRequestNumber" db:"purchase_request_number"`
	PurchaseOrderNumber       sql.NullString `json:"purchaseOrderNumber" db:"purchase_order_number"`
	SalesOrderNumber          sql.NullString `json:"salesOrderNumber" db:"sales_order_number"`
	BastNumber                sql.NullString `json:"bastNumber" db:"bast_number"`
	InvoiceNumber             sql.NullString `json:"invoiceNumber" db:"invoice_number"`
	StatusID                  int32          `json:"statusId" db:"status_id"`
	ProjectType               string         `json:"projectType" db:"project_type"`
	InvoiceDate               *time.Time     `json:"invoiceDate" db:"invoice_date"`
	DueAt                     *time.Time     `json:"dueAt" db:"due_at"`
	ReceivedAtLatest          *time.Time     `json:"receivedAtLatest" db:"received_at_latest"`
	IsPreOrder                bool           `json:"isPreOrder" db:"is_pre_order"`
	IsRfq                     bool           `json:"isRfq" db:"is_rfq"`
	IsInvoiceFinancing        bool           `json:"isInvoiceFinancing" db:"is_invoice_financing"`
	IsTermin                  bool           `json:"isTermin" db:"is_termin"`
	ShippingType              string         `json:"shippingType" db:"shipping_type"`
	ShippingAgency            sql.NullString `json:"shippingAgency" db:"shipping_agency"`
	ShippingMethod            sql.NullString `json:"shippingMethod" db:"shipping_method"`
	ShippingAwb               sql.NullString `json:"shippingAwb" db:"shipping_awb"`
	ShippingCost              float64        `json:"shippingCost" db:"shipping_cost"`
	ShippingDiscount          float64        `json:"shippingDiscount" db:"shipping_discount"`
	PaymentType               string         `json:"paymentType" db:"payment_type"`
	PaymentID                 sql.NullString `json:"paymentId" db:"payment_id"`
	PaymentCode               sql.NullString `json:"paymentCode" db:"payment_code"`
	PaymentBankName           sql.NullString `json:"paymentBankName" db:"payment_bank_name"`
	PaymentGateway            sql.NullString `json:"paymentGateway" db:"payment_gateway"`
	PaymentMethodFee          float64        `json:"paymentMethodFee" db:"payment_method_fee"`
	PaymentMethodPpnAmount    float64        `json:"paymentMethodPpnAmount" db:"payment_method_ppn_amount"`
	PaymentMethodPphAmount    float64        `json:"paymentMethodPphAmount" db:"payment_method_pph_amount"`
	PaymentAccountNumber      sql.NullString `json:"paymentAccountNumber" db:"payment_account_number"`
	PaymentURL                sql.NullString `json:"paymentUrl" db:"payment_url"`
	PaymentExpiredAt          *time.Time     `json:"paymentExpiredAt" db:"payment_expired_at"`
	TotalAmount               float64        `json:"totalAmount" db:"total_amount"`
	TotalGoodsAmount          float64        `json:"totalGoodsAmount" db:"total_goods_amount"`
	TotalServicesAmount       float64        `json:"totalServicesAmount" db:"total_services_amount"`
	Discount                  float64        `json:"discount" db:"discount"`
	FinalAmount               float64        `json:"finalAmount" db:"final_amount"`
	SellerIncome              float64        `json:"sellerIncome" db:"seller_income"`
	PpnBy                     sql.NullString `json:"ppnBy" db:"ppn_by"`
	PpnTotal                  float64        `json:"ppnTotal" db:"ppn_total"`
	PpnGoodsTotal             float64        `json:"ppnGoodsTotal" db:"ppn_goods_total"`
	PpnServicesTotal          float64        `json:"ppnServicesTotal" db:"ppn_services_total"`
	PpnShippingTotal          float64        `json:"ppnShippingTotal" db:"ppn_shipping_total"`
	PphTotal                  float64        `json:"pphTotal" db:"pph_total"`
	PphGoodsTotal             float64        `json:"pphGoodsTotal" db:"pph_goods_total"`
	PphServicesTotal          float64        `json:"pphServicesTotal" db:"pph_services_total"`
	PphShippingTotal          float64        `json:"pphShippingTotal" db:"pph_shipping_total"`
	RateTotalAmount           float64        `json:"rateTotalAmount" db:"rate_total_amount"`
	RateTotalAmountPercentage float64        `json:"rateTotalAmountPercentage" db:"rate_total_amount_percentage"`
	RateBaseAmount            float64        `json:"rateBaseAmount" db:"rate_base_amount"`
	RatePpnAmount             float64        `json:"ratePpnAmount" db:"rate_ppn_amount"`
	RatePphAmount             float64        `json:"ratePphAmount" db:"rate_pph_amount"`
	PurchaseRequestDate       time.Time      `json:"purchaseRequestDate" db:"purchase_request_date"`
}

type NewPet struct {
	Name    string `json:"name"`
	Type    string `json:"type,omitempty"`
	OwnerId int    `json:"ownerId"`
}
