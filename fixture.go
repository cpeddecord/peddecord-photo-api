package main

var img1 = Image{
	"img1",
	"is an image",
	"2018/02/03",
	"https://somehwere.com/img1.jpeg",
	"200",
	"200",
}

var img2 = Image{
	"img2",
	"is an image",
	"2018/02/03",
	"https://somehwere.com/img2.jpeg",
	"200",
	"200",
}

var img3 = Image{
	"img3",
	"is an image",
	"2018/02/03",
	"https://somehwere.com/img3.jpeg",
	"200",
	"200",
}

var img4 = Image{
	"img4",
	"is an image",
	"2018/02/03",
	"https://somehwere.com/img4.jpeg",
	"200",
	"200",
}

var thumb1 = Image{
	"thumb1",
	"is an image",
	"2018/02/03",
	"https://somehwere.com/thumb1.jpeg",
	"200",
	"200",
}

var thumb2 = Image{
	"thumb2",
	"is an image",
	"2018/02/03",
	"https://somehwere.com/thumb2.jpeg",
	"200",
	"200",
}

var FixtureGallery1 = Gallery{
	1,
	"category",
	"slug1",
	"Gallery 1 Lede",
	"SubLede",
	"is very long caption",
	thumb1,
	[]Image{img1, img2},
}

var FixtureGallery2 = Gallery{
	2,
	"category",
	"slug2",
	"Gallery 2 Lede",
	"SubLede",
	"is very long caption",
	thumb2,
	[]Image{img3, img4},
}
