// Copyright (c) 2020 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package portlist

import "testing"

func TestGetList(t *testing.T) {
	pl, err := GetList(nil)
	if err != nil {
		t.Fatal(err)
	}
	for i, p := range pl {
		t.Logf("[%d] %+v", i, p)
	}
	t.Logf("As String: %v", pl.String())
}

func BenchmarkGetList(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := GetList(nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}
