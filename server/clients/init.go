package clients

import "github.com/sw5005-sus/ceramicraft-order-mservice/server/config"

func InitAllClients(cfg *config.Conf) {
	_ = InitProductClient(cfg.CommodityClient)
	InitPaymentClient(cfg.PaymentClient)
}
