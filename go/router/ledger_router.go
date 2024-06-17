package router

type LedgerRouter struct {
	ledgerApplication app.LedgerApplication
	prefix            string
}

func NewLegerRouter(ledgerApplication app.LedgerApplication) *LedgerRouter {
	return &LedgerRouter{ledgerApplication: ledgerApplication, prefix: "/api/v2"}
}
