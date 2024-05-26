package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/ivanvanderbyl/telnyx-go"
)

type Authorizor struct {
	token string
}

func (a *Authorizor) BearerAuth(ctx context.Context, operationName string) (telnyx.BearerAuth, error) {
	return telnyx.BearerAuth{
		Token: a.token,
	}, nil
}

func run(ctx context.Context) error {
	var arg struct {
		BaseURL string
		ID      int64
	}
	flag.StringVar(&arg.BaseURL, "url", "https://api.telnyx.com/v2", "target server url")
	flag.Parse()

	auth := &Authorizor{token: os.Getenv("TELNYX_API_KEY")}

	client, err := telnyx.NewClient(arg.BaseURL, auth)
	if err != nil {
		return fmt.Errorf("create client: %w", err)
	}

	resp, err := client.ListCallControlApplications(ctx, telnyx.ListCallControlApplicationsParams{})
	if err != nil {
		return err
	}

	switch p := resp.(type) {
	case *telnyx.ListCallControlApplicationsResponse:
		data, err := p.MarshalJSON()
		if err != nil {
			return err
		}

		var out bytes.Buffer
		if err := json.Indent(&out, data, "", "  "); err != nil {
			return err
		}

		color.New(color.FgGreen).Println(out.String())
	case *telnyx.ListCallControlApplicationsNotFound:
		return errors.New("not found")
	}

	// fmt.Println(resp.(*telnyx.ListCallControlApplicationsResponse).Data[0].GetApplicationName().Get())

	// res, err := client.GetPetById(ctx, telnyx.GetPetByIdParams{
	// 	PetId: arg.ID,
	// })
	// if err != nil {
	// 	return fmt.Errorf("get pet %d: %w", arg.ID, err)
	// }

	// switch p := res.(type) {
	// case *petstore.Pet:
	// 	data, err := p.MarshalJSON()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	var out bytes.Buffer
	// 	if err := json.Indent(&out, data, "", "  "); err != nil {
	// 		return err
	// 	}
	// 	color.New(color.FgGreen).Println(out.String())
	// case *petstore.GetPetByIdNotFound:
	// 	return errors.New("not found")
	// }

	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		color.New(color.FgRed).Printf("%+v\n", err)
		os.Exit(2)
	}
}
