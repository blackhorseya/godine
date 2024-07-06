package wirex

import (
	"github.com/blackhorseya/godine/app/infra/configx"
	biz2 "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
)

// Injector is used to inject restaurant service.
type Injector struct {
	C *configx.Configuration
	A *configx.Application

	RestaurantService biz2.IRestaurantBiz
	MenuService       biz2.IMenuBiz
}
