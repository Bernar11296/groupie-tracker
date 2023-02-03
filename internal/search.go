package internal

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func search(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		Errors(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		Errors(w, http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}
	var result []Artists

	val := r.FormValue("search")

	fmt.Println(val)
	if val != "" {
		result = FindCard(val)
	}
	fmt.Println(result)

	err = tmp.Execute(w, result)
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}
}

func FindCard(s string) []Artists {
	res := TakeCards()
	for i := 1; i < 52; i++ {
		res[i].Relations = TakeConcert(i)
	}
	result := []Artists{}
	for p, v := range res {
		var flag bool
		if isContain(v.Name, s) {
			result = append(result, v)
			continue
		} else if isContain(strconv.Itoa(int(v.CreationDate)), s) {
			result = append(result, v)
			continue
		} else if isContain(v.FirstAlbum, s) {
			result = append(result, v)
			continue
		} else {
			for _, v1 := range v.Members {
				if isContain(v1, s) {
					result = append(result, v)
					flag = true
					break
				}
			}
		}
		if flag {
			continue
		}

		for j, v1 := range v.Relations.DatesLocations {
			fmt.Println(v1)
			if isContain(s, j) {
				result = append(result, res[p-1])
			}
			// for _, v2 := range v1 {
			// 	if isContain(v2, s) {
			// 		fmt.Println(v2)
			// 		result = append(result, v)
			// 		break
			// 	}
			// 	continue
			// }
			continue
		}
	}
	return result
}

func isContain(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
