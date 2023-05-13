package alexaquery_test

import (
	"os/user"
	"testing"

	"github.com/dannypeaky/alexaquery"
	"github.com/dannypeaky/alexaquery/data"
	"github.com/stretchr/testify/assert"
)

func TestQueryClient(t *testing.T) {
	usr, _ := user.Current()
	client := alexaquery.NewQueryClient(usr.HomeDir + "/Desktop/cookie.json")
	// client.Login()
	loggedIn, _ := client.CheckStatus()

	if !loggedIn {
		client.Login("Atnr|EwICIOIjd__zSKaR_a3a9_A9-t5_0D8zNcxdnsRGCbHHsV8cTWwm2b81_YVbE_xo_NEMQozIO7jRmiHfEWRSxipIs8XiFY10lH9aJhsBR00Is1hip6mDKB-rDl32D59hCa__BUqGwi1d-UPcc6KWCR6kh94BmIN0o7mAzI02zhwjCSub3I_7EThkabei3HikIWM2sdOigHRjmYuza0FB764zVjEcseT8Lpp8A6Q5VHe-i03DTZi6eau3eXn-6XEjsnaGhVVh4tRPkTPfzTdUWclr9qU2YqxxNpLRpTYmwPHVKqlFB5g1U86tOdmRgL5FFRrzR3TtDZkfWPrv9fTfZ8NdbAq_4tEg_-DJBD2lFydmGSYX3Tt8w76FejDN9n3P7g-6aytOP1DTsMtz0FseRqkWNfEV")
		loggedIn, _ = client.CheckStatus()
		if !loggedIn {
			panic("Not logged in")
		}
	}

	devices, err := client.GetDeviceList()
	assert.Nil(t, err)
	assert.IsType(t, devices, []data.Device{})

	notifications, err := client.GetNotifications("G090LF1180340N01", "A3S5BH2HU6VAYF")
	assert.Nil(t, err)
	assert.IsType(t, notifications, []data.Notification{})

	queue, err := client.GetQueue("G090LF1180340N01", "A3S5BH2HU6VAYF")
	assert.Nil(t, err)
	assert.IsType(t, queue, data.Queue{})
}
