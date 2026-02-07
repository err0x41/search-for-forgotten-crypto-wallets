package scanner

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type RPCResponse struct {
	Result string `json:"result"`
	Error  struct {
		Message string `json:"message"`
	} `json:"error"`
}

type SmartClient struct {
	Client *http.Client
}

func NewSmartClient() *SmartClient {
	return &SmartClient{
		Client: &http.Client{
			Timeout: 4 * time.Second,
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 10,
			},
		},
	}
}

// CheckEVMBalance реально запрашивает баланс через JSON-RPC
func (s *SmartClient) CheckEVMBalance(rpcUrl, address string) (float64, error) {
	payload := RPCRequest{
		Jsonrpc: "2.0",
		Method:  "eth_getBalance",
		Params:  []interface{}{address, "latest"},
		ID:      1,
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", rpcUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var rpcResp RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return 0, err
	}

	if rpcResp.Result == "" || rpcResp.Result == "0x0" {
		return 0, nil
	}

	// Конвертируем Hex в Wei, затем в ETH (упрощенно для проверки наличия)
	balanceWei, _ := strconv.ParseUint(rpcResp.Result[2:], 16, 64)
	return float64(balanceWei), nil
}
