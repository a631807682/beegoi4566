package u

import (
	"github.com/beego/beego/v2/server/web"
)

type RMRouters struct {
	Method   string
	Patterns []string
	Count    int64
}

func GetAllRegisteredRouters() (registRouters []RMRouters) {
	registRouters = make([]RMRouters, 0)
	content := web.PrintTree()
	methodsData := content["Data"]
	methods := content["Methods"].([]string)

	for _, method := range methods {
		if method != "POST" && method != "GET" {
			continue
		}
		router := RMRouters{
			Method:   method,
			Patterns: make([]string, 0),
			Count:    0,
		}

		existRouter := make(map[string]interface{}, 0)
		methodData := methodsData.(web.M)[method].(*[][]string)
		for _, v := range *methodData {
			routerPattern := v[0]
			_, ok := existRouter[routerPattern]
			if !ok {
				router.Patterns = append(router.Patterns, routerPattern)
				router.Count++
				existRouter[routerPattern] = nil
			}
		}
		registRouters = append(registRouters, router)
	}
	return
}
