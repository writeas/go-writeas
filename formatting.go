package writeas

import (
	"fmt"
	"net/http"
)

type BodyResponse struct {
	Body string `json:"body"`
}

// Markdown takes raw Markdown and renders it into usable HTML. See
// https://developers.write.as/docs/api/#render-markdown.
func (c *Client) Markdown(body, collectionURL string) (string, error) {
	p := &BodyResponse{}
	data := struct {
		RawBody       string `json:"raw_body"`
		CollectionURL string `json:"collection_url,omitempty"`
	}{
		RawBody:       body,
		CollectionURL: collectionURL,
	}

	env, err := c.post("/markdown", data, p)
	if err != nil {
		return "", err
	}

	var ok bool
	if p, ok = env.Data.(*BodyResponse); !ok {
		return "", fmt.Errorf("Wrong data returned from API.")
	}
	status := env.Code

	if status != http.StatusOK {
		return "", fmt.Errorf("Problem getting markdown: %d. %s\n", status, env.ErrorMessage)
	}
	return p.Body, nil
}
