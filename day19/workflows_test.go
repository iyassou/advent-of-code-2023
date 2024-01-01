package main

import "testing"

func sampleInput() string {
	return "px{a<2006:qkq,m>2090:A,rfg}\r\npv{a>1716:R,A}\r\nlnx{m>1548:A,A}\r\nrfg{s<537:gd,x>2440:R,A}\r\nqs{s>3448:A,lnx}\r\nqkq{x<1416:A,crn}\r\ncrn{x>2662:A,R}\r\nin{s<1351:px,qqz}\r\nqqz{s>2770:qs,m<1801:hdj,R}\r\ngd{a>3333:R,R}\r\nhdj{m>838:A,pv}\r\n\r\n{x=787,m=2655,a=1222,s=2876}\r\n{x=1679,m=44,a=2067,s=496}\r\n{x=2036,m=264,a=79,s=2244}\r\n{x=2461,m=1339,a=466,s=291}\r\n{x=2127,m=1623,a=2188,s=1013}"
}

func TestWorkflowsAccepts(t *testing.T) {
	wfs, parts := parse(sampleInput())
	expectedAccepts := []bool{true, false, true, false, true}
	expectedRatings := []int{7540, -1, 4623, -1, 6951}
	if len(parts) != len(expectedAccepts) {
		t.Fatalf("bruh: %d, %d", len(parts), len(expectedAccepts))
	}
	for i, part := range parts {
		expected := expectedAccepts[i]
		ok, err := wfs.Accepts(part, "in")
		if err != nil {
			t.Fatal(err)
		} else if ok != expected {
			t.Fatalf("[i=%d] expected %t, got %t", i, expected, ok)
		} else if ok {
			if rs := part.RatingsSum(); rs != expectedRatings[i] {
				t.Fatalf("[i=%d] expected %d, got %d", i, expectedRatings[i], rs)
			}
		}
	}
}

func TestWorkflowsDisinctCombinations(t *testing.T) {
	wfs, _ := parse(sampleInput())
	expected := 167409079868000
	if sum, err := wfs.DistinctCombinations("in"); err != nil {
		t.Fatal(err)
	} else if sum != expected {
		t.Fatalf("expected %d, got %d", expected, sum)
	}
}
