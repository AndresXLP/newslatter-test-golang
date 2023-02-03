package newsletter

import (
	"github.com/google/uuid"
)

type Interest string

var (
	InterestTech     Interest = "tech"
	InterestPolitics Interest = "politics"
	InterestSports   Interest = "sports"
)

type Subscription struct {
	UserID    uuid.UUID
	BlogID    uuid.UUID
	Interests []Interest
}

var Subs = []*Subscription{
	{
		UserID:    uuid.MustParse("e020e7f8-79e6-4d16-80ce-7cbf86cefe1f"),
		BlogID:    uuid.MustParse("b4f2e2d9-e399-46d0-a458-ef75527896c1"),
		Interests: []Interest{InterestSports, InterestTech},
	}, {
		UserID:    uuid.MustParse("e020e7f8-79e6-4d16-80ce-7cbf86cefe1f"),
		BlogID:    uuid.MustParse("b4f2e2d9-e399-46d0-a458-ef75527896c1"),
		Interests: []Interest{InterestSports, InterestTech},
	}, {
		UserID:    uuid.MustParse("e020e7f8-79e6-4d16-80ce-7cbf86cefe1f"),
		BlogID:    uuid.MustParse("b4f2e2d9-e399-46d0-a458-ef75527896c1"),
		Interests: []Interest{InterestSports, InterestTech},
	}, {
		UserID:    uuid.MustParse("e020e7f8-79e6-4d16-80ce-7cbf86cefe1f"),
		BlogID:    uuid.MustParse("b4f2e2d9-e399-46d0-a458-ef75527896c1"),
		Interests: []Interest{InterestSports, InterestTech},
	}, {
		UserID:    uuid.MustParse("e020e7f8-79e6-4d16-80ce-7cbf86cefe1f"),
		BlogID:    uuid.MustParse("b4f2e2d9-e399-46d0-a458-ef75527896c1"),
		Interests: []Interest{InterestSports, InterestTech},
	},
}
