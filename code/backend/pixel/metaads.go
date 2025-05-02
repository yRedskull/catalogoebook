package pixel

import (
	"bytes"
	"code/hash_app"
	"code/structs_utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Estruturas e funções para envio de evento ao Facebook Pixel

type PixelEvent struct {
	EventName       string `json:"event_name"`
	EventTime       int64  `json:"event_time"`
	UserData        any    `json:"user_data,omitempty"`         // Dados do usuário (hash de email, telefone, nome, etc.)
	CustomData      any    `json:"custom_data,omitempty"`       // Dados customizados
	ActionSource    string `json:"action_source"`               // Ex: "website"
	ClientIPAddress string `json:"client_ip_address,omitempty"` // IP do usuário
	ClientUserAgent string `json:"client_user_agent,omitempty"`
}

// PixelRequest representa a estrutura de payload para envio de eventos
type PixelRequest struct {
	Data          []PixelEvent `json:"data"`
	TestEventCode string       `json:"test_event_code,omitempty"`
}

// PixelManager gerencia a comunicação com o Facebook Pixel
type PixelManager struct {
	PixelID     string
	AccessToken string
	HTTPClient  *http.Client
}

// NewPixelManager cria e retorna uma nova instância de PixelManager
func NewPixelManager(pixelID, accessToken string) *PixelManager {
	return &PixelManager{
		PixelID:     pixelID,
		AccessToken: accessToken,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// SendEvent envia os eventos para o endpoint do Facebook Pixel
func (pm *PixelManager) SendEvent(events []PixelEvent, testEventCode string) error {
	url := fmt.Sprintf("https://graph.facebook.com/v12.0/%s/events?access_token=%s", pm.PixelID, pm.AccessToken)
	requestData := PixelRequest{
		Data:          events,
		TestEventCode: testEventCode,
	}

	payload, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("erro ao serializar os dados: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("erro ao criar a requisição: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := pm.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao enviar a requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Lê o corpo da resposta
		body, _ := io.ReadAll(resp.Body)

		// Retorna o código de status e a mensagem de erro do corpo
		return fmt.Errorf("erro: status code %d, mensagem: %s", resp.StatusCode, string(body))
	}

	return nil
}

func ExtractFbc(c *gin.Context) string {
	var fbc string

	if cookie_fbc, err_fbc := c.Cookie("_fbc"); err_fbc == nil {
		fbc = cookie_fbc
	} else if fbclid := c.Query("fbclid"); fbclid != "" {
		timestamp := time.Now().UnixMilli()
		fbc = fmt.Sprintf("fb.1.%d.%s", timestamp, fbclid)
	}

	return fbc
}

func GenerateUserData(first_name, last_name, contact,  fbc string, request_info structs_utils.RequestInfo) map[string]string {
	UserData := map[string]string{}

	if first_name != "" {
		UserData["fn"] = hash_app.HashString(first_name)
	}

	if last_name != "" {
		UserData["ln"] = hash_app.HashString(last_name)
	}
	if contact != "" {
		UserData["ph"] = hash_app.HashString(contact)
	}

	if request_info.IP != "" {
		UserData["client_ip_address"] = request_info.IP
	}
	if request_info.UserAgent != "" {
		UserData["client_user_agent"] = request_info.UserAgent
	}

	if fbc != "" {
		UserData["fbc"] = fbc
	}

	return UserData
}

// Estrutura que representa os dados do lead recebidos via JSON

// Handler do endpoint que recebe os dados do lead e envia o evento
func LeadRegister(pixel_id, access_token, first_name, last_name, contact, fbc, test_event_code string, request_info structs_utils.RequestInfo) error {

	pm := NewPixelManager(pixel_id, access_token)

	UserData := GenerateUserData(first_name, last_name, contact, fbc, request_info)

	leadEvent := PixelEvent{
		EventName:    "Lead",
		EventTime:    time.Now().Unix(),
		ActionSource: "website",
		UserData:     UserData,

		CustomData: map[string]any{
			"lead_type": "formulário de contato",
		},
	}

	// Envia o evento para o Facebook Pixel
	if err := pm.SendEvent([]PixelEvent{leadEvent}, test_event_code); err != nil {
		return fmt.Errorf("Erro ao enviar o evento: %v", err)
	}

	return nil
}

func PageView(pixel_id, access_token, fbc, test_event_code string, request_info structs_utils.RequestInfo) error {

	pm := NewPixelManager(pixel_id, access_token)

	UserData := GenerateUserData("", "", "", fbc, request_info)

	pageEvent := PixelEvent{
		EventName:    "PageView",
		EventTime:    time.Now().Unix(),
		ActionSource: "website",
		UserData:     UserData,
	}

	if err := pm.SendEvent([]PixelEvent{pageEvent}, test_event_code); err != nil {
		return fmt.Errorf("Erro ao enviar o evento: %v", err)
	}

	return nil
}
