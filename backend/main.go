package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// ----- Structs -----
type Input struct {
	Text             string   `json:"text"`
	ReleaseTypeId    string   `json:"release_type_id"`
	ImportantAspects []string `json:"important_aspects"`
}

type Response struct {
	Advice        string `json:"advice"`
	ImprovedPress string `json:"improved_press"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestToGPT struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ResponseFromGPT struct {
	Choices []Choices `json:"choices"`
}

type Choices struct {
	Message Message `json:"message"`
}

type ReleaseType struct {
	ReleaseTypeId   int    `json:"release_type_id"`
	ReleaseTypeName string `json:"release_type_name"`
}

type Press struct {
	Body string `json:"body"`
}

// ----- Constants -----
const modelName = "gpt-4.1-nano"

var aspectMap = map[string]string{
	"0": "時流性",
	"1": "話題性",
}

var db *sql.DB

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: .env ファイル読み込み失敗: %v\n", err)
	}
	initDB()
}

func initDB() {
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	if dbPassword == "" || dbPort == "" {
		fmt.Println("DB_PASSWORD または DB_PORT が設定されていません")
		os.Exit(1)
	}

	connStr := fmt.Sprintf("postgres://hackathon:%s@localhost:%s/hackathon_db?sslmode=require", dbPassword, dbPort)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("DB接続失敗: %v\n", err)
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		fmt.Printf("DB接続確認失敗: %v\n", err)
		os.Exit(1)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		frontend_port := os.Getenv("FRONTEND_PORT")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:"+frontend_port)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ----- Utilities -----
func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getReleaseTypeMap() map[string]string {
	releaseTypes, err := getReleaseTypesFromDB()
	if err != nil {
		fmt.Printf("getReleaseTypesFromDBエラー: %v\n", err)
		return map[string]string{}
	}

	releaseTypeMap := make(map[string]string)
	for _, rel := range releaseTypes {
		releaseTypeMap[fmt.Sprintf("%d", rel.ReleaseTypeId)] = rel.ReleaseTypeName
	}
	return releaseTypeMap
}

func getReleaseTypesFromDB() ([]ReleaseType, error) {
	file, err := os.Open("assets/SQL/get_release_type_table.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sqlQuery string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sqlQuery += strings.TrimSpace(scanner.Text()) + " "
	}

	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var releaseTypes []ReleaseType
	for rows.Next() {
		var rel ReleaseType
		if err := rows.Scan(&rel.ReleaseTypeId, &rel.ReleaseTypeName); err != nil {
			return nil, err
		}
		releaseTypes = append(releaseTypes, rel)
	}
	return releaseTypes, nil
}

func getPopularPressesFromDB(releaseTypeId string) ([]Press, error) {
	vars := map[string]interface{}{
		"release_type_id_int": releaseTypeId,
	}
	sqlQuery, err := generateReplacedText("assets/SQL/get_popular_presses.txt", vars)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var presses []Press
	for rows.Next() {
		var press Press
		if err := rows.Scan(&press.Body); err != nil {
			return nil, err
		}
		presses = append(presses, press)
	}
	return presses, nil
}

// ----- Handlers -----
func returnReleaseTypeMap(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, getReleaseTypeMap())
}

func returnAspectMap(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, aspectMap)
}

func returnPopularPresses(w http.ResponseWriter, r *http.Request) {
	if ok, str := isRequestOK(r); !ok {
		http.Error(w, str, http.StatusBadRequest)
		return
	}
	presses, err := getPopularPressesFromDB("2")
	if err != nil {
		http.Error(w, fmt.Sprintf("DBアクセスに失敗しました: %s", err), http.StatusInternalServerError)
		return
	}
	respondJSON(w, presses)
}

func returnFeedbackFromGPT(w http.ResponseWriter, r *http.Request) {
	if ok, str := isRequestOK(r); !ok {
		http.Error(w, str, http.StatusBadRequest)
		return
	}

	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSONデコードに失敗しました", http.StatusBadRequest)
		return
	}

	responseFromGPT, err := sendRequestToGPT(input)
	re := regexp.MustCompile(`\[(.*?)\]`)

	matches := re.FindAllStringSubmatch(responseFromGPT.Choices[0].Message.Content, -1)

	var response Response

	response.Advice = matches[0][1]
	response.ImprovedPress = matches[1][1]

	if err != nil {
		http.Error(w, fmt.Sprintf("GPTリクエスト失敗: %v", err), http.StatusInternalServerError)
		return
	}

	respondJSON(w, response)
}

func checkApiRequest(w http.ResponseWriter, r *http.Request) {
	if ok, msg := isRequestOK(r); !ok {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "OK")
}

// ----- Validation -----
func isRequestOK(r *http.Request) (bool, string) {
	if r.Method != http.MethodPost {
		return false, "POSTメソッドでアクセスしてください"
	}

	var input Input
	var buf bytes.Buffer
	tee := io.TeeReader(r.Body, &buf)
	if err := json.NewDecoder(tee).Decode(&input); err != nil {
		return false, "正しいJSON構造で送信してください"
	}
	r.Body = io.NopCloser(&buf)

	releaseTypeMap := getReleaseTypeMap()
	if _, ok := releaseTypeMap[input.ReleaseTypeId]; !ok {
		return false, "release_type_id: 不正な値指定です"
	}

	for _, v := range input.ImportantAspects {
		if _, ok := aspectMap[v]; !ok {
			return false, "important_aspects: 不正な値指定です"
		}
	}

	return true, ""
}

func sendRequestToGPT(input Input) ResponseFromGPT {
	file, err := os.Open("assets/GPT/system_prompt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var fullText string
	scanner := bufio.NewScanner(file)

	vars := map[string]interface{}{
		"industry_id":       input.IndustryId,
		"important_aspects": input.ImportantAspects,
	}

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`\$\{(\w+)\}`)
		result := re.ReplaceAllStringFunc(line, func(s string) string {
			matches := re.FindStringSubmatch(s)
			if len(matches) < 2 {
				return s
			}
			key := matches[1]

			val, ok := vars[key]
			if !ok {
				return s
			}

			// interface{} を string に変換
			switch v := val.(type) {
			case string:
				return v
			case []string:
				return strings.Join(v, ",") // 配列ならカンマ区切りで結合
			default:
				return fmt.Sprintf("%v", val)
			}
		})
		fullText += result + "\n"
	}
// ----- GPT API -----
func sendRequestToGPT(input Input) (ResponseFromGPT, error) {
	releaseTypeMap := getReleaseTypeMap()
	popularPresses, err := getPopularPressesFromDB(input.ReleaseTypeId)
	if err != nil {
		return ResponseFromGPT{}, err
	}
	varsSystemPrompt := map[string]interface{}{
		"release_type_id":   releaseTypeMap[input.ReleaseTypeId],
		"important_aspects": input.ImportantAspects,
		"popular_presses":   popularPresses,
	}

	fullText, err := generateReplacedText("assets/GPT/system_prompt.txt", varsSystemPrompt)
	if err != nil {
		return ResponseFromGPT{}, err
	}

	reqBody := RequestToGPT{
		Model:    modelName,
		Messages: []Message{{Role: "system", Content: fullText}, {Role: "user", Content: "私は以下のようなプレスリリースを書きました。このプレスリリースを改善することには私の人生がかかっています。\n" + input.Text}},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return ResponseFromGPT{}, err
	}

	openaiKey := os.Getenv("OPENAPI_KEY")
	if openaiKey == "" {
		return ResponseFromGPT{}, fmt.Errorf("OPENAPI_KEY が設定されていません")
	}

	httpreq, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(jsonData))
	if err != nil {
		return ResponseFromGPT{}, err
	}
	httpreq.Header.Set("Content-Type", "application/json")
	httpreq.Header.Set("Authorization", "Bearer "+openaiKey)

	client := &http.Client{}
	resp, err := client.Do(httpreq)
	if err != nil {
		return ResponseFromGPT{}, err
	}
	defer resp.Body.Close()

	var response ResponseFromGPT
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return ResponseFromGPT{}, err
	}

	return response, nil
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
