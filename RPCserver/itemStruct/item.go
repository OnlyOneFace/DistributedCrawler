package itemStruct

import "fmt"

/*
* @
* @Author:
* @Date: 2020/3/27 0:31
 */
type ItemSaverService struct {
}

func (*ItemSaverService) Save(item interface{}, result *string) error {
	*result = fmt.Sprintf("Item Saver: got item: %v", item)
	return nil
}
