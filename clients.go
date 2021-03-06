package manager

// Client represents a Mainflux client. Each client is owned by one user, and
// it is assigned with the unique identifier and (temporary) access key.
type Client struct {
	Owner       string            `json:"-"`
	ID          string            `json:"id"`
	Type        string            `json:"type"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Key         string            `json:"key"`
	Meta        map[string]string `json:"meta"`
}

func (c *Client) validate() error {
	if c.Name == "" || len(c.Name) > 50 {
		return ErrMalformedClient
	}

	return nil
}

// ClientRepository specifies a client persistence API.
type ClientRepository interface {
	// Save persists the client. Successful operation is indicated by unique
	// identifier accompanied by nil error response. A non-nil error is
	// returned to indicate operation failure.
	Save(Client) (string, error)

	// One retrieves the client having the provided identifier, that is owned
	// by the specified user.
	One(string, string) (Client, error)

	// Remove removes the client having the provided identifier, that is owned
	// by the specified user.
	Remove(string, string) error
}
