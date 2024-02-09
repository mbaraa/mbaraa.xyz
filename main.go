package main

import (
	"embed"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

var (
	//go:embed templates
	aaa embed.FS

	//go:embed res
	bbb embed.FS
)

func main() {
	bgs := []string{"/res/bgs/01.jpg", "res/bgs/02.jpg"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFS(aaa, "templates/index.html"))
		_ = tmpl.Execute(w, map[string]any{
			"BGLink": bgs[rand.Intn(3)],
			"Deployments": []struct {
				Link        string
				Title       string
				Description string
			}{
				{Link: "https://mbaraa.com", Title: "mbaraa.com", Description: "Baraa Al-Masri's personal blog and portfolio."},
				{Link: "https://htmx.pics", Title: "HTMX Pictures", Description: "HTMX Profile Picture Maker."},
				{Link: "https://gdscasu.com", Title: "GDSC - ASU's Home Page", Description: "The home page of GDSC-ASU, includes details about the chapter, like events, projects and members."},
				{Link: "https://logogen.gdscasu.com", Title: "GDSC Logo Generator", Description: "The most optimized and working GDSC logo generator."},
				{Link: "https://scheduler.gdscasu.com", Title: "ASU Scheduler", Description: "The only ASU semester schedule generator!"},
				{Link: "https://eloi.mbaraa.com", Title: "Eloi's API", Description: "A package and overlays fetch api for Eloi, more details here https://wiki.gentoo.org/wiki/Eloi"},
				{Link: "https://killedbyjosa.org", Title: "Killed By JOSA", Description: "A killedbygoogle.com fork for dead projects at josa.ngo"},
				{Link: "https://github-graph-drawer.mbaraa.com", Title: "Github Graph Drawer", Description: "A tool that helps typing text on your GitHub contributions graph."},
				{Link: "https://slidemd.com", Title: "SlideMD", Description: "Turn your markdowns into portable web slides in a jiffy."},
				{Link: "https://apis.gdscasu.com", Title: "GDSC-ASU Helper APIs", Description: "Some APIs used for testing purposes."},
				{Link: "https://dotsync.org", Title: "Dotsync's website", Description: "A small, free, open-source, blazingly fast dotfiles synchronizer!"},
				{Link: "https://api.dotsync.org", Title: "Dotsync's syncer API", Description: "A small, free, open-source, blazingly fast dotfiles synchronizer!"},
			},
		})
	})
	http.Handle("/res/", http.FileServer(http.FS(bbb)))
	log.Println("server running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}