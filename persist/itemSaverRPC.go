package persist

import "log"

/*
* @
* @Author:
* @Date: 2020/3/26 23:40
 */
type ItemSaverService struct {
}

func (*ItemSaverService) Save(item interface{}, result *string) error {
	log.Printf("Item Saver: got item: %v\n", item)
	*result = "ok"
	return nil
}
