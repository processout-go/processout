package processout

import (
	"net/http"
)

var (
	// RequestAPIVersion is the version of the API used in requests made
	// with this package
	RequestAPIVersion = "1.4.0.0"
	// Host is the URL where API requests are made
	Host = "https://api.processout.com"
)

// ProcessOut wraps all the components of the package in a
// single structure
type ProcessOut struct {
	// APIVersion is the version of the API to use
	APIVersion string
	// ProcessOut project ID
	projectID string
	// ProcessOut project secret key
	projectSecret string
}

// Options represents the options available when doing a request to the
// ProcessOut API
type Options struct {
	IdempotencyKey string   `json:"-"`
	Expand         []string `json:"expand"`
	Filter         string   `json:"filter"`
	Limit          uint64   `json:"limit"`
	EndBefore      string   `json:"end_before"`
	StartAfter     string   `json:"start_after"`
	DisableLogging bool     `json:"-"`
}

// New creates a new struct *ProcessOut with the given API credentials. It
// initializes all the resources available so they can be used immediately.
func New(projectID, projectSecret string) *ProcessOut {
	p := &ProcessOut{
		APIVersion:    RequestAPIVersion,
		projectID:     projectID,
		projectSecret: projectSecret,
	}

	return p
}

func setupRequest(client *ProcessOut, opt *Options, req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", client.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(client.projectID, client.projectSecret)
}

// NewActivity creates a new Activity object
func (c *ProcessOut) NewActivity(prefill ...*Activity) *Activity {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Activity, or none.")
	}
	if len(prefill) == 0 {
		return &Activity{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewAddon creates a new Addon object
func (c *ProcessOut) NewAddon(prefill ...*Addon) *Addon {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Addon, or none.")
	}
	if len(prefill) == 0 {
		return &Addon{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewAPIRequest creates a new APIRequest object
func (c *ProcessOut) NewAPIRequest(prefill ...*APIRequest) *APIRequest {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the APIRequest, or none.")
	}
	if len(prefill) == 0 {
		return &APIRequest{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewAPIVersion creates a new APIVersion object
func (c *ProcessOut) NewAPIVersion(prefill ...*APIVersion) *APIVersion {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the APIVersion, or none.")
	}
	if len(prefill) == 0 {
		return &APIVersion{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewAuthorizationRequest creates a new AuthorizationRequest object
func (c *ProcessOut) NewAuthorizationRequest(prefill ...*AuthorizationRequest) *AuthorizationRequest {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the AuthorizationRequest, or none.")
	}
	if len(prefill) == 0 {
		return &AuthorizationRequest{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewCard creates a new Card object
func (c *ProcessOut) NewCard(prefill ...*Card) *Card {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Card, or none.")
	}
	if len(prefill) == 0 {
		return &Card{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewCardInformation creates a new CardInformation object
func (c *ProcessOut) NewCardInformation(prefill ...*CardInformation) *CardInformation {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the CardInformation, or none.")
	}
	if len(prefill) == 0 {
		return &CardInformation{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewCoupon creates a new Coupon object
func (c *ProcessOut) NewCoupon(prefill ...*Coupon) *Coupon {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Coupon, or none.")
	}
	if len(prefill) == 0 {
		return &Coupon{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewCustomer creates a new Customer object
func (c *ProcessOut) NewCustomer(prefill ...*Customer) *Customer {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Customer, or none.")
	}
	if len(prefill) == 0 {
		return &Customer{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewToken creates a new Token object
func (c *ProcessOut) NewToken(prefill ...*Token) *Token {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Token, or none.")
	}
	if len(prefill) == 0 {
		return &Token{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewDiscount creates a new Discount object
func (c *ProcessOut) NewDiscount(prefill ...*Discount) *Discount {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Discount, or none.")
	}
	if len(prefill) == 0 {
		return &Discount{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewEvent creates a new Event object
func (c *ProcessOut) NewEvent(prefill ...*Event) *Event {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Event, or none.")
	}
	if len(prefill) == 0 {
		return &Event{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewGateway creates a new Gateway object
func (c *ProcessOut) NewGateway(prefill ...*Gateway) *Gateway {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Gateway, or none.")
	}
	if len(prefill) == 0 {
		return &Gateway{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewGatewayConfiguration creates a new GatewayConfiguration object
func (c *ProcessOut) NewGatewayConfiguration(prefill ...*GatewayConfiguration) *GatewayConfiguration {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the GatewayConfiguration, or none.")
	}
	if len(prefill) == 0 {
		return &GatewayConfiguration{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewInvoice creates a new Invoice object
func (c *ProcessOut) NewInvoice(prefill ...*Invoice) *Invoice {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Invoice, or none.")
	}
	if len(prefill) == 0 {
		return &Invoice{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewInvoiceDetail creates a new InvoiceDetail object
func (c *ProcessOut) NewInvoiceDetail(prefill ...*InvoiceDetail) *InvoiceDetail {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the InvoiceDetail, or none.")
	}
	if len(prefill) == 0 {
		return &InvoiceDetail{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewCustomerAction creates a new CustomerAction object
func (c *ProcessOut) NewCustomerAction(prefill ...*CustomerAction) *CustomerAction {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the CustomerAction, or none.")
	}
	if len(prefill) == 0 {
		return &CustomerAction{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewDunningAction creates a new DunningAction object
func (c *ProcessOut) NewDunningAction(prefill ...*DunningAction) *DunningAction {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the DunningAction, or none.")
	}
	if len(prefill) == 0 {
		return &DunningAction{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewPlan creates a new Plan object
func (c *ProcessOut) NewPlan(prefill ...*Plan) *Plan {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Plan, or none.")
	}
	if len(prefill) == 0 {
		return &Plan{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewProduct creates a new Product object
func (c *ProcessOut) NewProduct(prefill ...*Product) *Product {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Product, or none.")
	}
	if len(prefill) == 0 {
		return &Product{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewProject creates a new Project object
func (c *ProcessOut) NewProject(prefill ...*Project) *Project {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Project, or none.")
	}
	if len(prefill) == 0 {
		return &Project{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewRefund creates a new Refund object
func (c *ProcessOut) NewRefund(prefill ...*Refund) *Refund {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Refund, or none.")
	}
	if len(prefill) == 0 {
		return &Refund{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewSubscription creates a new Subscription object
func (c *ProcessOut) NewSubscription(prefill ...*Subscription) *Subscription {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Subscription, or none.")
	}
	if len(prefill) == 0 {
		return &Subscription{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewTransaction creates a new Transaction object
func (c *ProcessOut) NewTransaction(prefill ...*Transaction) *Transaction {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Transaction, or none.")
	}
	if len(prefill) == 0 {
		return &Transaction{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewTransactionOperation creates a new TransactionOperation object
func (c *ProcessOut) NewTransactionOperation(prefill ...*TransactionOperation) *TransactionOperation {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the TransactionOperation, or none.")
	}
	if len(prefill) == 0 {
		return &TransactionOperation{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewWebhook creates a new Webhook object
func (c *ProcessOut) NewWebhook(prefill ...*Webhook) *Webhook {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Webhook, or none.")
	}
	if len(prefill) == 0 {
		return &Webhook{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewWebhookEndpoint creates a new WebhookEndpoint object
func (c *ProcessOut) NewWebhookEndpoint(prefill ...*WebhookEndpoint) *WebhookEndpoint {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the WebhookEndpoint, or none.")
	}
	if len(prefill) == 0 {
		return &WebhookEndpoint{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}
