package domain

type ClientTypeType int
type ClientIdType string
type ClientSecretType string

const (
  clientType ClientTypeType = 1 + iota
  public bool
)

type Client struct {
  clientType ClientTypeType
  clientId ClientIdType
  clientSecret ClientSecretType
  name string
  
}

var ClientTypes = [...]string {
  "confidential",
  "public",
}

func (clientType ClientTypeType) String() string {
 return ClientTypes[clientType - 1]
}
