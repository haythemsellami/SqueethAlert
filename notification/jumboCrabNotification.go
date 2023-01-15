package notification

import (
	"fmt"

	mcommon "github.com/haythemsellami/SqueethAlert/common"
)

func JumboCrabNotify(encodedEvents []interface{}) {
	fmt.Println("notification")

	for _, event := range encodedEvents {
		switch eventType := event.(type) {
		case mcommon.USDCQueuedEvent:
			fmt.Printf("%+v\n", eventType)
		case mcommon.CrabQueuedEvent:
			fmt.Printf("%+v\n", eventType)
		default:
			fmt.Println("unknown")
		}
	}
}
