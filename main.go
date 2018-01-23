package main

import (
	"errors"
	"net/url"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Flags
	var (
		modeFlag = cli.StringFlag{
			Name:  "mode",
			Usage: "mode",
		}
		placeFlag = cli.StringFlag{
			Name:  "place",
			Usage: "place",
		}
		startDateFlag = cli.StringFlag{
			Name:  "start_date",
			Usage: "start_date",
		}
		endDateFlag = cli.StringFlag{
			Name:  "end_date",
			Usage: "end_date",
		}
		categoryIdFlag = cli.StringFlag{
			Name:  "category_id",
			Usage: "category_id",
		}
		genreIdFlag = cli.StringFlag{
			Name:  "genre_id",
			Usage: "genre_id",
		}
		fromAccountIdFlag = cli.StringFlag{
			Name:  "from_account_id",
			Usage: "from_account_id",
		}
		toAccountIdFlag = cli.StringFlag{
			Name:  "to_account_id",
			Usage: "to_account_id",
		}
		accountIdFlag = cli.StringFlag{
			Name:  "account_id",
			Usage: "account_id",
		}
	)
	app := cli.NewApp()
	app.Name = "zaim"
	app.Usage = "Zaim CLI Client (readonly)"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:   "auth",
			Usage:  "Authenticate",
			Action: Authenticate,
		},
		{
			Name:   "verify",
			Usage:  "Verify",
			Action: Verify,
		},
		{
			Name:   "money",
			Usage:  "Money",
			Action: Money,
			Flags: []cli.Flag{
				modeFlag,
				placeFlag,
				startDateFlag,
				endDateFlag,
				categoryIdFlag,
				genreIdFlag,
				fromAccountIdFlag,
				toAccountIdFlag,
				accountIdFlag,
			},
		},
		{
			Name:   "category",
			Usage:  "Category",
			Action: Category,
		},
		{
			Name:   "genre",
			Usage:  "Genre",
			Action: Genre,
		},
		{
			Name:   "account",
			Usage:  "Account",
			Action: Account,
		},
	}
	app.Run(os.Args)
}

func Authenticate(c *cli.Context) error {
	config := NewConfig()
	client := NewClient1(
		config.ConsumerKey,
		config.ConsumerSecret,
	)
	cred, err := client.GetAccessToken()
	if err != nil {
		return errors.New("client get error")
	}
	config.SetAccessToken(cred.Token, cred.Secret)
	err = config.Save()
	if err != nil {
		return errors.New("Fatal error config file: %s \n")
	}
	return nil
}

func action(path string, params url.Values) error {
	url := "https://api.zaim.net" + path
	config := NewConfig()
	client := NewClient2(
		config.ConsumerKey,
		config.ConsumerSecret,
		config.AccessToken,
		config.AccessTokenSecret,
	)
	err := client.Get(url, params)
	if err != nil {
		return errors.New("client get error")
	}
	return nil
}

func Verify(c *cli.Context) error {
	return action("/v2/home/user/verify", nil)
}

func addParams(c *cli.Context, params url.Values, names ...string) {
	for _, name := range names {
		if c.IsSet(name) {
			params.Set(name, c.String(name))
		}
	}
}

func Money(c *cli.Context) error {
	params := url.Values{}
	addParams(c, params,
		"mode",
		"place",
		"start_date",
		"end_date",
		"category_id",
		"genre_id",
		"from_account_id",
		"to_account_id",
		"account_id",
	)
	return action("/v2/home/money", params)
}

func Category(c *cli.Context) error {
	params := url.Values{}
	return action("/v2/home/category", params)
}

func Genre(c *cli.Context) error {
	params := url.Values{}
	return action("/v2/home/genre", params)
}

func Account(c *cli.Context) error {
	params := url.Values{}
	return action("/v2/home/account", params)
}
