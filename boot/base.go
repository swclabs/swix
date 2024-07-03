package boot

import "swclabs/swipecore/internal/types"

type IBase interface {
	Connect(adapter types.IAdapter) error
}
