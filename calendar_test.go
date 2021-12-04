package advent

func Example_Day3() {
	PowerCons("testdata/3.given", 5)
	PowerCons("testdata/3.input", 12)
	LifeSupportRating("testdata/3.given", 5)
	LifeSupportRating("testdata/3.input", 12)
	// output:
	// 198
	// 3009600
	// 230
	// 6940518
}

// ----------------------------------------

func Example_Day2() {
	Navigate("testdata/2.given", WithoutAim)
	Navigate("testdata/2.input", WithoutAim)
	Navigate("testdata/2.given", WithAim)
	Navigate("testdata/2.input", WithAim)
	// output:
	// 150
	// 1580000
	// 900
	// 1251263225
}

// ----------------------------------------

func Example_Day1() {
	IncreasingWindow("testdata/1.given", 1)
	IncreasingWindow("testdata/1.input", 1)
	IncreasingWindow("testdata/1.given", 3)
	IncreasingWindow("testdata/1.input", 3)
	// output:
	// 7
	// 1288
	// 5
	// 1311
}
