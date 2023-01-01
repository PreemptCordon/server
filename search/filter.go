package search

import (
	"context"

	"github.com/preemptcordon/server/db"
	"github.com/preemptcordon/server/obj"
)

func SearchFilter(thisuser obj.UserObj, terms obj.SearchTerms) []obj.UnrankedResult {
	var firstpass []obj.UnrankedResult
	var prefilter []string
	if terms.Consider.Author {
		var check string
		err := db.DBPool.QueryRow(context.Background(), "select 'something'").Scan(&check)
		if err != nil {
			panic(err)
		}
		prefilter = append(prefilter, check)
	}
	if terms.Consider.Summary {
		var check string
		err := db.DBPool.QueryRow(context.Background(), "select 'something'").Scan(&check)
		if err != nil {
			panic(err)
		}
		prefilter = append(prefilter, check)
	}
	if terms.Consider.Title {
		var check string
		err := db.DBPool.QueryRow(context.Background(), "select 'something'").Scan(&check)
		if err != nil {
			panic(err)
		}
		prefilter = append(prefilter, check)
	}

	for _, match := range prefilter {
		article = db.Article.Load(match)
		if !CheckTagDistance(article, terms.Tags) {
			continue
		}
		if CheckBlock(article.Author, thisuser) {
			continue
		}
		if !CheckACL(article.ACL, thisuser) {
			continue
		}
		if CheckMute(article, thisuser) {
			continue
		}
		if CheckModBlock(article) {
			match.PotentialVisitDenied(terms)
			continue
		}
		firstpass = append(firstpass, match)
	}
	return firstpass
}
