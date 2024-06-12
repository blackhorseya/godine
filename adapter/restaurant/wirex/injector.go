package wirex

import (
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/entity/restaurant/biz"
)

// Injector is used to inject restaurant service.
type Injector struct {
	A *configx.Application

	RestaurantService biz.IRestaurantBiz
	MenuService       biz.IMenuBiz
}
