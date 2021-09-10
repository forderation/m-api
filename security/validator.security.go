package security

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Claims struct {
	IAT int64
	ISS string
	UID uint
	RO  string
}

func MDRequestBarrier() gin.HandlerFunc {
	return func(c *gin.Context) {
		//pbk, err := importKey("D:\\1.ON GOING\\Infotech\\ILab-API-Auth-Server\\keys\\public.pem")
		pbk, err := importKey(os.Getenv("PUBLIC_KEY"))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, "SV Error : 1")
			c.Abort()
			return
		}

		tS, err := readTokenFromCookies(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "ERR:AUTH:READ-TOKEN")
			c.Abort()
			return
		}

		b, err := validate([]byte(tS), pbk)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "ERR:AUTH:VERIFIER")
			c.Abort()
			return
		}

		sess := sessions.Default(c)
		sess.Set("CU-ACCESS", b)

		c.Next()
		return
	}
}

func MDAuthMapper(authKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := GetSessionAccessData(c).RO

		if !strings.ContainsAny(authKey, role) {
			c.JSON(http.StatusUnauthorized, "ERR:AUTH:ROLE-NOT-SATISFIED")
			c.Abort()
			return
		}

		c.Next()
		return
	}
}

func GetSessionAccessData(c *gin.Context) Claims {
	sess := sessions.Default(c)
	return sess.Get("CU-ACCESS").(Claims)
}

func readTokenFromCookies(c *gin.Context) (string, error) {
	ch, err := c.Cookie("ifx-ath")
	if err != nil {
		return "", err
	}

	cb, err := c.Cookie("ifx-at")
	if err != nil {
		return "", err
	}

	cs, err := c.Cookie("ifx-st")
	if err != nil {
		return "", err
	}

	if ch == "" || cb == "" || cs == "" {
		return "", fmt.Errorf("token is blank or invalid")
	}

	str := fmt.Sprintf("%s.%s.%s", ch, cb, cs)
	return str, nil
}

func validate(tokenString []byte, publicKey *rsa.PublicKey) (Claims, error) {
	var b Claims
	token, err := jwt.Parse(tokenString, jwt.WithVerify(jwa.RS256, publicKey))
	if err != nil {
		return Claims{}, err
	}
	buf, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		return Claims{}, err
	}
	err = json.Unmarshal(buf, &b)
	if err != nil {
		return Claims{}, err
	}
	return b, nil
}

func importKey(path string) (*rsa.PublicKey, error) {
	s, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Key Not Available")
	}

	block, _ := pem.Decode(s)
	pbk, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	r := pbk.(*rsa.PublicKey)
	return r, nil
}
