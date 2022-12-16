package main

import "testing"

func Test_ParseDate(t *testing.T) {
	suite := map[string]struct {
		test string
		want string
	}{
		"rss1": {
			test: "Wed, 14 Dec 2022 14:24:07 +0000",
			want: "2022-12-14T14:24:07",
		},
		"rss2": {
			test: "Fri, 04 Nov 2022 21:37:38 +0000",
			want: "2022-11-04T21:37:38",
		},
		"atom1": {
			test: "2022-11-09T19:31:00-08:00",
			want: "2022-11-09T19:31:00",
		},
		"atom2": {
			test: "2022-11-10T03:35:02-08:00",
			want: "2022-11-10T03:35:02",
		},
	}
	for name, test := range suite {
		t.Run(name, func(t *testing.T) {
			d := ParseDate(test.test)
			if d.String() != test.want {
				t.Errorf("wanted %q, got %q for %q", test.want, d.String(), test.test)
			}
		})
	}
}
