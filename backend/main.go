package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Input struct {
	IndustryId       string   `json:"industry_id"`
	ImportantAspects []string `json:"important_aspects"`
}

type Response struct {
	Advice        string
	ImprovedPress string `json:"improved_press"`
}

func main() {
	http.HandleFunc("/industry_ids", getIndustryIDs)
	http.HandleFunc("/get_feedback_from_GPT", getFeedbackFromGPT)
	http.HandleFunc("/check", checkApiRequest)
	http.ListenAndServe(":8080", nil)
}

// 指定可能な業種とそのid
func getIndustryIDs(w http.ResponseWriter, r *http.Request) {
	industryMap := map[string]string{
		"0": "その他",
		"1": "水産・農林業",
		"2": "鉱業",
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(industryMap); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getFeedbackFromGPT(w http.ResponseWriter, r *http.Request) {
	if isOK, errStr := isRequestOK(r); !isOK {
		fmt.Fprintln(w, errStr)
		return
	}
	response := Response{"あなたのプレスリリースは以下の点で問題があります...", "#[業界初!!]..."}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

// Apiに対するリクエストについて、とりあえず構造とアクセス方法が正しければOK, ダメなら原因を表示する
func checkApiRequest(w http.ResponseWriter, r *http.Request) {
	if isOK, errStr := isRequestOK(r); !isOK {
		fmt.Fprintln(w, errStr)
		return
	}
	fmt.Fprintln(os.Stdout, "OK")
}

// Apiに対するリクエストについて、とりあえず構造とアクセス方法が正しいかを判定
func isRequestOK(r *http.Request) (bool, string) {
	if r.Method != http.MethodPost {
		return false, "POSTメソッドでアクセスしてください"
	}

	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return false, err.Error() + "\n正しい形でJSONデータを渡してください"
	}

	return true, ""
}
