package create_json_ld

import (
	"encoding/json"
	"log"

	"github.com/kazarena/json-gold/ld"
	"github.com/remotejob/apartment_ru_go/domains"
)

func Create(proc *ld.JsonLdProcessor, options *ld.JsonLdOptions, articlefull domains.Articlefull) []byte {

	createdstr := articlefull.Created.Format("2006-01-02")
	updatedstr := articlefull.Updated.Format("2006-01-02")

	pagelink := "http://" + articlefull.Site + "/jobs/" + articlefull.Stitle + ".html"

	publisher := map[string]interface{}{"@type": "Organization", "name": "Remote Job Finland OY", "logo": map[string]interface{}{"@type": "ImageObject", "url": "http://mazurov.eu/img/free_for_job.png", "height": "256px", "width": "256px"}}
	image := map[string]interface{}{"@type": "ImageObject", "url": "http://" + articlefull.Site + "/assets/img/free_for_job.png", "height": "256px", "width": "256px"}
	mainEntityOfPage := map[string]interface{}{"@type": "WebPage", "@id": "http://" + articlefull.Site}

	doc := map[string]interface{}{
		"@context":         "http://schema.org",
		"@type":            "Article",
		"author":           articlefull.Author,
		"headline":         articlefull.Title,
		"publisher":        publisher,
		"image":            image,
		"datepublished":    createdstr,
		"datemodified":     updatedstr,
		"mainEntityOfPage": mainEntityOfPage,
		// "keywords":         articlefull.Tags,
		"url": pagelink,
		//		"description":         "We love to do stuff to help people and stuff",
		"articleSection": "realestate",
		"articleBody":    articlefull.Contents,
	}

	comp, err := proc.Compact(doc, nil, options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)

	}

	b, _ := json.Marshal(comp)
	//	fmt.Println(string(b))
	return b
}
