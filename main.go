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
	bgs := []string{"/res/bgs/01.webp", "res/bgs/02.webp", "res/bgs/03.webp"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFS(aaa, "templates/index.html"))
		_ = tmpl.Execute(w, map[string]any{
			"BGLink": bgs[rand.Intn(len(bgs))],
			"Deployments": []struct {
				Link        string
				Title       string
				Description string
			}{
				{Link: "https://mbaraa.com", Title: "mbaraa.com", Description: "Baraa Al-Masri's personal blog and portfolio."},
				{Link: "https://htmx.pics", Title: "HTMX Pictures", Description: "HTMX Profile Picture Maker."},
				{Link: "https://gdsclogo.mbaraa.xyz", Title: "GDSC Logo Generator", Description: "The most optimized and working GDSC logo generator."},
				{Link: "https://api.gdsclogo.mbaraa.xyz", Title: "GDSC Logo Generator", Description: "The most optimized and working GDSC logo generator's API"},
				{Link: "https://eloi.mbaraa.xyz", Title: "Eloi's API", Description: "A package and overlays fetch api for Eloi, more details here https://wiki.gentoo.org/wiki/Eloi"},
				{Link: "https://killedbyjosa.org", Title: "Killed By JOSA", Description: "A killedbygoogle.com fork for dead projects at josa.ngo"},
				{Link: "https://github-graph-drawer.mbaraa.xyz", Title: "Github Graph Drawer", Description: "A tool that helps typing text on your GitHub contributions graph."},
				{Link: "https://slidemd.mbaraa.xyz", Title: "SlideMD", Description: "Turn your markdowns into portable web slides in a jiffy."},
				{Link: "https://apis.gdscasu.com", Title: "GDSC-ASU Helper APIs", Description: "Some APIs used for testing purposes."},
				{Link: "https://dotsync.org", Title: "Dotsync's website", Description: "A small, free, open-source, blazingly fast dotfiles synchronizer!"},
				{Link: "https://api.dotsync.org", Title: "Dotsync's syncer API", Description: "A small, free, open-source, blazingly fast dotfiles synchronizer!"},
				{Link: "https://linuxipsum.org", Title: "Linux Ipsum", Description: "A Linux flavored Lorem ipsum."},
				// {Link: "https://appflowy.mbaraa.xyz", Title: "AppFlowy", Description: "My App Flowy personal hosting or sth."},
				// {Link: "https://traggo.mbaraa.xyz", Title: "Traggo", Description: "My Traggo personal hosting or sth."},
				{Link: "https://waterna.net", Title: "Waterna", Description: "The coolest water delivery service."},
				{Link: "https://dankmuzikk.com", Title: "DankMuzikk", Description: "Create, Share and Play Music Playlists."},
				// {Link: "https://htmxqr.mbaraa.xyz", Title: "HTMX QR Codes", Description: "You May've Heard Of htmx On Floppies, Get Ready For htmx On QR Codes!"},
				{Link: "https://yt.mbaraa.xyz", Title: "ytscrape", Description: "YouTube Scraper Searcher"},
			},
		})
	})
	http.Handle("/res/", http.FileServer(http.FS(bbb)))
	log.Println("server running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
