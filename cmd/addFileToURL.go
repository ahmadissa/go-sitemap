package main

import (
	"os"

	Sitemap "github.com/ahmadissa/go-sitemap"
)

func main() {
	if len(os.Args) != 4 {
		println("Usage: addFileToURL localfile.xml http://domain/sitemap.xml sitemap.xml")
	}
	Sitemap.AddFileToURL(os.Args[1], os.Args[2], os.Args[3])
}
