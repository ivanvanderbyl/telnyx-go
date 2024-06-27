# Telnyx Go SDK

Generated from [team-telnyx/openapi](https://github.com/team-telnyx/openapi/) with some modifications to correct invalid types and add missing fields.

Code generated using [Ogen](https://github.com/ogen-go/ogen)

## Example Client Usage

```go
type Authorizor struct {
  token string
}

func (a *Authorizor) BearerAuth(_ context.Context, operationName string) (telnyx.BearerAuth, error) {
  return telnyx.BearerAuth{
    Token: a.token,
  }, nil
}

func do() error {
  auth := &Authorizor{token: os.Getenv("TELNYX_API_KEY")}

  client, err := telnyx.NewClient(telnyx.TelnyxAPIServer, auth)
  if err != nil {
    return fmt.Errorf("create client: %w", err)
  }

  resp, err := client.ListCallControlApplications(ctx, telnyx.ListCallControlApplicationsParams{})
  if err != nil {
    return err
  }

  switch p := resp.(type) {
  case *telnyx.ListCallControlApplicationsResponse:
    // p.Data
  case *telnyx.ListCallControlApplicationsNotFound:
    return errors.New("not found")
  }
}
```

Also see [cmd](./cmd) for more examples.
