package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateLinkWts(number string, txt string) string {
	txtEncoded := url.QueryEscape(txt)

	link := fmt.Sprintf("https://wa.me/55%s?text=%s", number, txtEncoded)

	return link
}


func ExistsKey(m map[string]any, chave string) bool {
	_, existe := m[chave] 
	return existe
}

func ValidateUsername(username string) (string, error) {
	validPattern := `^[a-z0-9]+$`

	matched, err := regexp.MatchString(validPattern, username)
	if err != nil {
		return "", err
	}

	if !matched {
		return "", fmt.Errorf("username inválido: somente letras e números são permitidos e não pode haver espaços ou caracteres especiais")
	}

	return username, nil
}

func GetDomain(c *gin.Context) string {
	host := c.Request.Host
	domain, _ := url.Parse(strings.Split(host, ":")[0])

	return domain.String()
}

func ToLower(value string) string {
	return strings.ToLower(value)
}

func Trim(value string) string {
	return strings.Trim(value, " ")
}

func GetFirstName(full_name string) string {
	return strings.Split(full_name, " ")[0]
}

func CapitalizeSentence(sentence string) string {
	caser := cases.Title(language.Und)
	return caser.String(sentence)
}

func ValidatePlan(c *gin.Context, data_of_user bson.M) error {
	if !data_of_user["active"].(bool) || data_of_user["plan"] == "" {
		c.HTML(http.StatusOK, "activate_account.html", nil)
		return fmt.Errorf("precisa ativar a conta")
	}

	return nil
}