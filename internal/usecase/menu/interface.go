package menu

import (
	"archv1/internal/entity"
	"context"
)

type MenuUseCaseI interface {
	GetSiteMenus(ctx context.Context, filter entity.Filter, lang string) (entity.SiteMenuListResponse, error)
	List(ctx context.Context, filter entity.Filter, lang string) (entity.ListMenuResponse, error)
	GetByID(ctx context.Context, menuID int, lang string) (entity.GetMenuResponse, error)
	Create(ctx context.Context, menu entity.CreateMenuRequest) (entity.CreateMenuResponse, error)
	Update(ctx context.Context, menu entity.UpdateMenuRequest) (entity.UpdateMenuResponse, error)
	UpdateColumns(ctx context.Context, fields entity.UpdateMenuColumnsRequest) (entity.UpdateMenuResponse, error)
	Delete(ctx context.Context, menuID, deletedBy int) (entity.DeleteMenuResponse, error)
}
