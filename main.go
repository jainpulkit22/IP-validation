package main
import (
	"strings"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckForIPv4 checks whether the parsed string is
// a valid IPv4 address or not.
func CheckForIPv4(ip string) bool {
	var countDot int = strings.Count(ip, ".")
	if countDot!=3 {
		return false
	}
	ip += "."
	var part string = ""
	for i := 0; i < len(ip); i++ {
		if ip[i]=='.' {
			if len(part)==0 {
				return false
			}
			num,err := strconv.ParseInt(part, 10, 32)
			if err!=nil {
				return false
			}
			if num<0 || num>255 {
				return false
			}
			part = ""
		} else {
			part += string(ip[i])
		}
	}
	return true
}

// CheckForIPv6 checks whether the parsed string is
// a valid IPv6 address or not.
func CheckForIPv6(ip string) bool {
	var countColon int = strings.Count(ip, ":")
	if countColon!=7 {
		return false
	}
	ip += ":"
	var part string = ""
	for i := 0; i < len(ip); i++ {
		if ip[i]==':' {
			if len(part)!=4 {
				return false;
			}
			part = ""
		} else {
			if (ip[i]<'A' ||  ip[i]>'F') && (ip[i]<'a' && ip[i]>'f') && (ip[i]<'0' || ip[i]>'9') {
				return false
			}
			part += string(ip[i])
		}
	}
	return true
}

func validateIP(c *gin.Context) {
	var ip string
	err := c.BindJSON(&ip)
	if err!=nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, ip)
	var val1 bool
	val1 = CheckForIPv4(ip)
	if val1 {
		c.String(200, "\n %s is a valid IPv4 address", ip)
	} else {
		val1 = CheckForIPv6(ip)
		if val1 {
			c.String(200, "\n %s is a valid IPv6 address", ip)
		}
	}
	if !val1 {
		c.String(200, "\n %s is not a valid IP address", ip)
	}
}

func main() {
	router := gin.Default()
	router.POST("/ip", validateIP)

	router.Run(":8080")

}