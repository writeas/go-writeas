# go-writeas

[![godoc](https://godoc.org/github.com/writeas/go-writeas?status.svg)](https://godoc.org/github.com/writeas/go-writeas)

Official Write.as Go client library.

## Installation

```bash
go get github.com/writeas/go-writeas
```

## Documentation

See all functionality and usages in the [API documentation](https://developers.write.as/docs/api/).

### Example usage

```go
import "github.com/writeas/go-writeas"

func main() {
	// Create the client
	c := writeas.NewClient()

	// Publish a post
	p, err := c.CreatePost(&writeas.PostParams{
		Title:   "Title!",
		Content: "This is a post.",
		Font:    "sans",
	})
	if err != nil {
		// Perhaps show err.Error()
	}

	// Save token for later, since it won't ever be returned again
	token := p.Token

	// Update a published post
	p, err = c.UpdatePost(p.ID, token, &writeas.PostParams{
		Content: "Now it's been updated!",
	})
	if err != nil {
		// handle
	}

	// Get a published post
	p, err = c.GetPost(p.ID)
	if err != nil {
		// handle
	}

	// Delete a post
	err = c.DeletePost(p.ID, token)
}
```

## Contributing

The library covers our usage, but might not be comprehensive of the API. So we always welcome contributions and improvements from the community. Before sending pull requests, make sure you've done the following:

* Run `goimports` on all updated .go files.
* Document all exported structs and funcs.

## License

MIT
