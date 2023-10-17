package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Success  bool `json:"success"`
	Response struct {
		Result struct {
			AtomicalID                   string      `json:"atomical_id"`
			TopLevelRealmAtomicalID      string      `json:"top_level_realm_atomical_id"`
			TopLevelRealmName            string      `json:"top_level_realm_name"`
			NearestParentRealmAtomicalID string      `json:"nearest_parent_realm_atomical_id"`
			NearestParentRealmName       string      `json:"nearest_parent_realm_name"`
			RequestFullRealmName         string      `json:"request_full_realm_name"`
			FoundFullRealmName           string      `json:"found_full_realm_name"`
			MissingNameParts             interface{} `json:"missing_name_parts"`
			Candidates                   []struct {
				TxNum                int    `json:"tx_num"`
				AtomicalID           string `json:"atomical_id"`
				TxID                 string `json:"txid"`
				CommitHeight         int    `json:"commit_height"`
				RevealLocationHeight int    `json:"reveal_location_height"`
			} `json:"candidates"`
			NearestParentRealmSubrealmMintRules struct {
				NearestParentRealmAtomicalID string      `json:"nearest_parent_realm_atomical_id"`
				Note                         string      `json:"note"`
				CurrentHeight                int         `json:"current_height"`
				CurrentHeightRules           interface{} `json:"current_height_rules"`
				NextHeight                   int         `json:"next_height"`
				NextHeightRules              interface{} `json:"next_height_rules"`
				Next2Height                  int         `json:"next_2_height"`
				Next2HeightRules             interface{} `json:"next_2_height_rules"`
				Next3Height                  int         `json:"next_3_height"`
				Next3HeightRules             interface{} `json:"next_3_height_rules"`
			} `json:"nearest_parent_realm_subrealm_mint_rules"`
			NearestParentRealmSubrealmMintAllowed bool `json:"nearest_parent_realm_subrealm_mint_allowed"`
		} `json:"result"`
	} `json:"response"`
}

func main() {
	url := "https://ep.atomicals.xyz/proxy/blockchain.atomicals.get_realm_info"
	fileName := "example.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()
	for i := 'a'; i <= 'z'; i++ {
		for j := 'a'; j <= 'z'; j++ {
			for k := 'a'; k <= 'z'; k++ {
				param := fmt.Sprintf("{\"params\": [\"%c%c%c\", 0]}", i, j, k)

				// 发送 POST 请求
				resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(param)))
				if err != nil {
					fmt.Println("HTTP请求失败:", err)
					return
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading response body:", err)
					return
				}

				var response Response
				if err := json.Unmarshal(body, &response); err != nil {
					fmt.Println("JSON解析失败:", err)
					return
				}
				// fmt.Println("AtomicalID:", response.Response.Result)
				fmt.Println("处理到:" + string(i) + string(j) + string(k))
				if response.Response.Result.AtomicalID != "" {
					// fmt.Println("AtomicalID:", response.Response.Result)
				} else {
					_, err = file.WriteString(string(i) + string(j) + string(k) + "\n")
					if err != nil {
						fmt.Println("写入文件失败:", err)
						return
					}
					fmt.Println("AtomicalID为空" + string(i) + string(j) + string(k) + "\n")
				}
			}
		}
	}
}
