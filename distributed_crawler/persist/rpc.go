package persist

import (
	"github.com/elastic/go-elasticsearch/v8"
	"goprojects/crawler/engine"
	"goprojects/crawler/persist"
)

type ItemSaverService struct {
	Es    *elasticsearch.Client
	Index string
}

func (i *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(i.Es, i.Index, item)
	if err != nil {
		return err
	}

	*result = "ok"
	return nil
}
