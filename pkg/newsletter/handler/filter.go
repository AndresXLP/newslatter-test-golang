package handler

import (
	"fmt"
	"strings"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

type Filter struct {
	UserID    string   `json:"userId" form:"userId"`
	BlogID    string   `json:"blogId" form:"blogId"`
	Interests []string `json:"interests" form:"interests"`
}

func (f *Filter) ToSliceInterest() []newsletter.Interest {
	val := f.Interests
	fmt.Println(val)
	value := strings.Split(strings.Join(f.Interests, ","), ",")
	var interest = make([]newsletter.Interest, len(value))
	for i := range interest {
		interest[i] = newsletter.Interest(value[i])
	}
	return interest
}
