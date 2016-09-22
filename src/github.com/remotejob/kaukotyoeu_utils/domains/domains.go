package domains

import (
	"encoding/xml"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Gphrase comment
type Gphrase struct {
	Phrase string `bson:"Phrase"`
	Rating int    `bson:"Rating"`
}

type Articlefull struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Title  string
	Stitle string
	// Tags      string
	Contents  string
	Mcontents string
	Site      string
	Author    string
	Created   time.Time
	Updated   time.Time
}

type Articletotempalte struct {
	Title  string
	Stitle string
	// Tags      string
	Contents  string
	Mcontents string
	Site      string
	Created   string
	Updated   string
	Jsonld    []byte
}

type Article struct {
	Title string
	// Tags      string
	Contents  string
	Mcontents string
	Author    string
}

type Sitemap_from_db struct {
	Stitle  string
	Site    string
	Updated time.Time
}

type ServerConfig struct {
	General struct {
		Themes string
		Locale string
	}
	Dbmgo struct {
		Addrs     []string
		Database  string
		Username  string
		Password  string
		Mechanism string
	}

	Dbmgoext struct {
		Addrs     []string
		Database  string
		Username  string
		Password  string
		Mechanism string
	}
	Sites struct {
		Site []string
	}
	Dirs struct {
		Sitemapsdir string
		Webrootdir  string
	}

	Routes struct {
		Mainroute string
	}

	Files struct {
		Commonwords string
	}
}

type JobOffer struct {
	Title       string
	Tags        []string
	Description string
}

type SitemapObj struct {
	Changefreq    string
	Hoursduration float64
	Loc           string
	Lastmod       string
}

//type BlogItem struct {
//	Stopic     string
//	Topic      string
//	Stitle     string
//	Title      string
//	Contents   string
//	Created_at time.Time
//	Updated_at time.Time
//}
//
//type Blog struct {
//	//	Topic string
//	Items map[string][]BlogItem
//}
//
//type KeywordObj struct {
//	Keyword string
//	Cl      int
//	Lang    string
//}
//
//type Contents struct {
//	Title      string
//	Moto       string
//	Contents   string
//	Created_at time.Time
//	Updated_at time.Time
//}

type Pages struct {
	//	Version string   `xml:"version,attr"`
	XMLName xml.Name `xml:"urlset"`
	XmlNS   string   `xml:"xmlns,attr"`
	//	XmlImageNS string   `xml:"xmlns:image,attr"`
	//	XmlNewsNS  string   `xml:"xmlns:news,attr"`
	Pages []*Page `xml:"url"`
}

type Page struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	Lastmod    string   `xml:"lastmod"`
	Changefreq string   `xml:"changefreq"`
	//	Name       string   `xml:"news:news>news:publication>news:name"`
	//	Language   string   `xml:"news:news>news:publication>news:language"`
	//	Title      string   `xml:"news:news>news:title"`
	//	Keywords   string   `xml:"news:news>news:keywords"`
	//	Image      string   `xml:"image:image>image:loc"`
}

// type Config struct {
// 	Maintitle string
// 	Subtitle  string
// 	Cv        []struct {
// 		Name string
// 		Path string
// 		Img  string
// 		Item []struct {
// 			Title    string
// 			Rank     int
// 			Duration int
// 			Link     string
// 			Extra    string
// 			Img      string
// 		}
// 	}
// }

type Job struct {
	Maintitle string
	Subtitle  string
	Jobs      []struct {
		Name string
		Path string
		Img  string
		Item []struct {
			Title    string
			Rank     int
			Duration string
			Position string
			Details  string
			Location string
			Country  string
		}
	}
}
