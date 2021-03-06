package advent

func Example_Day7() {
	HorizontalAlign("testdata/7.given")
	HorizontalAlign("testdata/7.input")
	HorizontalAlignReal("testdata/7.given")
	HorizontalAlignReal("testdata/7.input")
	// output:
	// 37
	// 326132
	// 168
	// 88612508
}

// ----------------------------------------

func Example_Day6() {
	LanternFish("testdata/6.given", 80)
	LanternFish("testdata/6.input", 80)
	LanternFish("testdata/6.given", 256)
	LanternFish("testdata/6.input", 256)
	// output:
	// 5934
	// 354564
	// 26984457539
	// 1609058859115
}

// ----------------------------------------

func Example_Day5() {
	CountHVIntersections("testdata/5.given", 10, 10)
	CountHVIntersections("testdata/5.input", 1000, 1000)
	CountAllIntersections("testdata/5.given", 10, 10)
	CountAllIntersections("testdata/5.input", 1000, 1000)
	// output:
	// 5
	// 6666
	// 12
	// 19081
}

// ----------------------------------------

func Example_Day4() {
	Winner("testdata/4.given")
	Winner("testdata/4.input")
	Looser("testdata/4.input")
	// output:
	// 4512
	// 16716
	// 4880
}

// ----------------------------------------

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
