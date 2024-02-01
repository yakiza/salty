package api

import "net/http"

type InvoiceHandler struct {
}

func (h InvoiceHandler) NewInvoice(w http.ResponseWriter, r *http.Request) {

}

func NewInvoiceHandler() InvoiceHandler {
	return InvoiceHandler{}
}
