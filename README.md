Client for [Podio](https://podio.com) written in Go.

## Example

```go
func main() {
  authToken, err := podio.AuthWithUserCredentials("my-client-id", "my-client-secret", "my-username", "my-password")

  if err != nil {
    fmt.Println("Auth failed: ", err)
    return
  }

  client := podio.NewClient(authToken)
  orgs, err := client.GetOrganizations()

  if err != nil {
    fmt.Println("Failed to get orgs: ", err)
    return
  }

  for _, org := range orgs {
    fmt.Println("Org: ", org.Name)
  }
}
```

See [example/main.go](example/main.go).

## Status

- The client supports authentication with username and password (see [Username and Password flow](https://developers.podio.com/authentication/username_password)), app authentication (see [App authentication flow](https://developers.podio.com/authentication/app_auth)) and server-side flow (see [Server-side flow](https://developers.podio.com/authentication/server_side)).
- Only supports a fraction of the API methods available, specifically around organizations, spaces, apps, items and files.
- Only a few number of fields have been defined per type.

## Contributors

The following people have contributed to podio-go:

- andreas
- stengaard
