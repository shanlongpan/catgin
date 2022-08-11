package dao

import (
	"context"
	"github.com/shanlongpan/catgin/config"
	"github.com/shanlongpan/catgin/xlog"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (p *Product) Migrate(ctx context.Context) {
	err := config.DB.AutoMigrate(p)
	if err != nil {
		xlog.Errorln(ctx, err)
	}
}

func (p *Product) Insert(ctx context.Context, products []*Product) {
	for _, item := range products {
		config.DB.Create(item)
	}
}

func (p *Product) Select(ctx context.Context, params map[string]interface{}) (error, *Product) {
	result := Product{}
	config.DB.First(&result, 1)
	return nil, &result
}
