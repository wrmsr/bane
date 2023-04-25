//
/*
https://en.wikipedia.org/wiki/List_of_birds_by_common_name
https://en.wikipedia.org/w/api.php?action=parse&page=List_of_birds_by_common_name&format=json&prop=links
*/
package birdle

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	eu "github.com/wrmsr/bane/core/errors"
)

//

type wikiApiResult struct {
	Parse wikiParse `json:"parse"`
}

type wikiParse struct {
	Title  string     `json:"title"`
	PageId int        `json:"pageid"`
	Links  []wikiLink `json:"links"`
}

type wikiLink struct {
	Title string `json:"*"`
}

func fetchWikiParse(ctx context.Context, page string) (wp *wikiParse, err error) {
	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=parse&page=%s&format=json&prop=links", page)

	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return
	}

	req = req.WithContext(ctx)
	client := http.DefaultClient

	var res *http.Response
	res, err = client.Do(req)
	defer eu.AppendInvoke(&err, eu.Close(res.Body))
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		return
	}

	var war wikiApiResult
	if err = json.NewDecoder(res.Body).Decode(&war); err != nil {
		return
	}

	wp = &war.Parse
	return
}

//

const DefaultWikiBirdsPage = "List_of_birds_by_common_name"

func FetchWikiLinks(ctx context.Context, page string) ([]string, error) {
	wp, err := fetchWikiParse(ctx, page)
	if err != nil {
		return nil, err
	}

	ws := make([]string, len(wp.Links))
	for i, l := range wp.Links {
		ws[i] = l.Title
	}
	return ws, nil
}
