package menu

import (
	"archv1/internal/entity"
	"archv1/internal/service/menu"
	"context"
)

type MenuUseCase struct {
	menuService menu.MenuServiceI
}

func NewMenuUseCase(service menu.MenuServiceI) MenuUseCaseI {
	return &MenuUseCase{
		menuService: service,
	}
}

func (m *MenuUseCase) GetSiteMenus(ctx context.Context, filter entity.Filter, lang string) (entity.SiteMenuListResponse, error) {
	response, err := m.menuService.GetSiteMenus(ctx, filter, lang)
	if err != nil {
		return entity.SiteMenuListResponse{}, err
	}

	return response, nil
}

func (u *MenuUseCase) List(ctx context.Context, filter entity.Filter, lang string) (entity.ListMenuResponse, error) {
	menuResponse, err := u.menuService.List(ctx, filter, lang)
	if err != nil {
		return entity.ListMenuResponse{}, err
	}

	return menuResponse, nil
}

func (u *MenuUseCase) GetByID(ctx context.Context, menuID int, lang string) (entity.GetMenuResponse, error) {
	menuResponse, err := u.menuService.GetByID(ctx, menuID, lang)
	if err != nil {
		return entity.GetMenuResponse{}, err
	}

	return menuResponse, nil
}

func (u *MenuUseCase) Create(ctx context.Context, menu entity.CreateMenuRequest) (entity.CreateMenuResponse, error) {
	menuResponse, err := u.menuService.Create(ctx, menu)
	if err != nil {
		return entity.CreateMenuResponse{}, err
	}

	return menuResponse, nil
}

func (u *MenuUseCase) Update(ctx context.Context, menu entity.UpdateMenuRequest) (entity.UpdateMenuResponse, error) {
	menuResponse, err := u.menuService.Update(ctx, menu)
	if err != nil {
		return entity.UpdateMenuResponse{}, err
	}

	return menuResponse, nil
}

func (u *MenuUseCase) UpdateColumns(ctx context.Context, fields entity.UpdateMenuColumnsRequest) (entity.UpdateMenuResponse, error) {
	menuResponse, err := u.menuService.UpdateColumns(ctx, fields)
	if err != nil {
		return entity.UpdateMenuResponse{}, err
	}

	return menuResponse, nil
}

func (u *MenuUseCase) Delete(ctx context.Context, menuID int) (entity.DeleteMenuResponse, error) {
	menuResponse, err := u.menuService.Delete(ctx, menuID)
	if err != nil {
		return entity.DeleteMenuResponse{}, err
	}

	return menuResponse, nil
}
