package menu

import (
	"archv1/internal/entity"
	"archv1/internal/repository/postgres/menu"
	"context"
)

type MenuService struct {
	menuRepo menu.MenuRepository
}

func NewMenuService(menuRepo menu.MenuRepository) MenuServiceI {
	return &MenuService{
		menuRepo: menuRepo,
	}
}

func (m *MenuService) GetSiteMenus(ctx context.Context, filter entity.Filter, lang string) (entity.SiteMenuListResponse, error) {
	response, err := m.menuRepo.GetSiteMenus(ctx, filter, lang)
	if err != nil {
		return entity.SiteMenuListResponse{}, err
	}

	return response, nil
}

func (u *MenuService) List(ctx context.Context, filter entity.Filter, lang string) (entity.ListMenuResponse, error) {
	menuResponse, err := u.menuRepo.List(ctx, filter, lang)
	if err != nil {
		return entity.ListMenuResponse{}, err
	}

	return menuResponse, nil
}

func (u *MenuService) GetByID(ctx context.Context, menuID int, lang string) (entity.GetMenuResponse, error) {
	menuResponse, err := u.menuRepo.GetByID(ctx, menuID, lang)
	if err != nil {
		return entity.GetMenuResponse{}, err
	}

	return menuResponse, nil
}

func (u *MenuService) Create(ctx context.Context, menu entity.CreateMenuRequest) (entity.CreateMenuResponse, error) {
	menuResponse, err := u.menuRepo.Create(ctx, menu)
	if err != nil {
		return entity.CreateMenuResponse{}, err
	}

	return menuResponse, nil
}

func (u *MenuService) Update(ctx context.Context, menu entity.UpdateMenuRequest) (entity.UpdateMenuResponse, error) {
	menuResponse, err := u.menuRepo.Update(ctx, menu)
	if err != nil {
		return entity.UpdateMenuResponse{}, err
	}

	return menuResponse, nil
}

func (u *MenuService) UpdateColumns(ctx context.Context, fields entity.UpdateMenuColumnsRequest) (entity.UpdateMenuResponse, error) {
	menuResponse, err := u.menuRepo.UpdateColumns(ctx, fields)
	if err != nil {
		return entity.UpdateMenuResponse{}, err
	}

	return menuResponse, nil
}

func (u *MenuService) Delete(ctx context.Context, menuID, deletedBy int) (entity.DeleteMenuResponse, error) {
	menuResponse, err := u.menuRepo.Delete(ctx, menuID, deletedBy)
	if err != nil {
		return entity.DeleteMenuResponse{}, err
	}

	return menuResponse, nil
}
