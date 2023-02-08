package web

import (
	"encoding/json"
	"github.com/oktapascal/app-sima/models/domain"
	"github.com/oktapascal/app-sima/utils"
)

type MenuResponses struct {
	KodeKlpMenu string      `json:"kode_klp_menu"`
	NamaKlpMenu string      `json:"nama_klp_menu"`
	ListMenu    interface{} `json:"list_menu"`
}

func ConvertToMenuResponses(menu domain.Menu) MenuResponses {
	return MenuResponses{
		KodeKlpMenu: menu.KodeKlpMenu,
		NamaKlpMenu: menu.NameKlpMenu,
		ListMenu:    menu.ListMenu,
	}
}

func ConvertToMenusResponses(menus []domain.Menu) []MenuResponses {
	var menuResponse []MenuResponses

	var listMenu []map[string]interface{}
	for _, menu := range menus {
		var objMap map[string][]map[string]interface{}

		err := json.Unmarshal(menu.ListMenuJson, &objMap)
		utils.PanicIfError(err)

		for _, list := range objMap["data_menu"] {
			listMenu = append(listMenu, map[string]interface{}{
				"name_menu": list["name_menu"],
				"icon_menu": list["icon_menu"],
				"list_menu": list["list_menu"],
			})
		}

		menu.ListMenu = listMenu
		menuResponse = append(menuResponse, ConvertToMenuResponses(menu))
	}

	return menuResponse
}
