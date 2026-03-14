package model

import (
	"time"

	ceramicraftsecure "github.com/sw5005-sus/ceramicraft-secure"
)

type Order struct {
	ID                int       `gorm:"primaryKey;autoIncrement"`
	OrderNo           string    `gorm:"type:varchar(64);unique;not null"` // 订单编号
	UserID            int       `gorm:"not null"`                         // 下单用户
	Status            int       `gorm:"not null"`                         // 订单状态 (0-无效状态，不应该有此状态； 1-创建； 2-已付款； 3-已发货； 4-已收获； 5-取消)
	TotalAmount       int       `gorm:"type:int;not null"`                // 总金额
	PayAmount         int       `gorm:"type:int;not null"`                // 实际支付金额
	PayTime           time.Time `gorm:"default:null"`                     // 支付时间
	CreateTime        time.Time `gorm:"autoCreateTime"`                   // 创建时间
	UpdateTime        time.Time `gorm:"autoUpdateTime"`                   // 更新时间
	ReceiverFirstName string    `gorm:"type:varchar(64)"`                 // 收货人姓名
	ReceiverLastName  string    `gorm:"type:varchar(64)"`                 // 收货人姓名
	ReceiverPhone     string    `gorm:"type:varchar(128)"`                // 收货人电话
	ReceiverAddress   string    `gorm:"type:varchar(512)"`                // 收货地址
	ReceiverCountry   string    `gorm:"type:varchar(64)"`                 // 收货人国家
	ReceiverZipCode   string    `gorm:"type:varchar(128)"`                // 收货人邮政编码
	ShippingFee       int       `gorm:"type:int;not null"`                // 运费
	Tax               int       `gorm:"type:int;not null"`                // 税
	Remark            string    `gorm:"type:varchar(256)"`                // 备注
	LogisticsNo       string    `gorm:"type:varchar(64)"`                 // 物流单号
	DeliveryTime      time.Time `gorm:"default:null"`                     // 发货时间
	ConfirmTime       time.Time `gorm:"default:null"`                     // 收货确认时间
}

// TableName sets the insert table name for this struct type
func (Order) TableName() string {
	return "orders"
}

func (o *Order) Encrypt() error {
	if o.ReceiverPhone != "" {
		encrypReceiverPhone, err := ceramicraftsecure.AesEncrypt(o.ReceiverPhone)
		if err != nil {
			return err
		}
		o.ReceiverPhone = encrypReceiverPhone
	}
	if o.ReceiverZipCode != "" {
		encrypReceiverZipCode, err := ceramicraftsecure.AesEncrypt(o.ReceiverZipCode)
		if err != nil {
			return err
		}
		o.ReceiverZipCode = encrypReceiverZipCode
	}
	if o.ReceiverAddress != "" {
		encrypReceiverAddress, err := ceramicraftsecure.AesEncrypt(o.ReceiverAddress)
		if err != nil {
			return err
		}
		o.ReceiverAddress = encrypReceiverAddress
	}
	return nil
}

func (o *Order) Decrypt() error {
	if o.ReceiverPhone != "" {
		decryptReceiverPhone, err := ceramicraftsecure.AesDecrypt(o.ReceiverPhone)
		if err != nil {
			return err
		}
		o.ReceiverPhone = decryptReceiverPhone
	}
	if o.ReceiverZipCode != "" {
		decryptReceiverZipCode, err := ceramicraftsecure.AesDecrypt(o.ReceiverZipCode)
		if err != nil {
			return err
		}
		o.ReceiverZipCode = decryptReceiverZipCode
	}
	if o.ReceiverAddress != "" {
		decryptReceiverAddress, err := ceramicraftsecure.AesDecrypt(o.ReceiverAddress)
		if err != nil {
			return err
		}
		o.ReceiverAddress = decryptReceiverAddress
	}
	return nil
}
