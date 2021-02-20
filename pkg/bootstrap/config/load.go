package config

import (
	"github.com/goradd/goradd/pkg/config"
	"github.com/goradd/goradd/pkg/html"
	"github.com/goradd/goradd/pkg/page"
	"path/filepath"
)

// Loader is the injected loader. Set it during your application's initialization
// if you want to load bootstrap differently than below.
var Loader func(page.FormI)

// Configuration options for Bootstrap

// LoadBootstrap loads the various bootstrap files required by bootstrap. It is called automatically
// by the bootstrap components, but this gives you an opportunity to customize where the client
// gets the files.
func LoadBootstrap(form page.FormI) {
	if form.Page().HasMetaTag("viewport") {
		// already loaded
		return
	}
	if Loader != nil {
		Loader(form)
	} else {
		form.Page().AddHtmlHeaderTag(html.VoidTag{
			Tag:  "meta",
			Attr: html.NewAttributes().
				AddAttributeValue("name", "viewport").
				AddAttributeValue("content","width=device-width, initial-scale=1, shrink-to-fit=no"),
			})
		form.AddJQuery()
		if config.Release {
			form.AddJavaScriptFile("https://cdn.jsdelivr.net/npm/popper.js@1.16.1/dist/umd/popper.min.js", false,
				html.NewAttributes().Set("integrity", "sha384-9/reFTGAW83EW2RDu2S0VKaIzap3H66lZH81PoYlFhbGU+6BZp6G7niu735Sk7lN").Set("crossorigin", "anonymous"))
			form.AddJavaScriptFile("https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/js/bootstrap.min.js", false,
				html.NewAttributes().Set("integrity", "sha384-+YQ4JLhjyBLPDQt//I+STsc9iw4uQqACwlvpslubQzn4u2UU2UFM80nGisd026JF").Set("crossorigin", "anonymous"))
			form.AddStyleSheetFile("https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/css/bootstrap.min.css",
				html.NewAttributes().Set("integrity", "sha384-B0vP5xmATw1+K9KRQjQERJvTumQW0nPEzvF6L/Z6nronJ3oUOFUFpCjEUQouq2+l").Set("crossorigin", "anonymous"))
			form.AddJavaScriptFile(filepath.Join(BootstrapAssets(), "js", "gr.bs.shim.js"), false, nil)
		} else {
			form.AddJavaScriptFile(filepath.Join(BootstrapAssets(), "js", "bootstrap.bundle.js"), false, nil)
			form.AddStyleSheetFile(filepath.Join(BootstrapAssets(), "css", "bootstrap.min.css"), nil)
			form.AddJavaScriptFile(filepath.Join(BootstrapAssets(), "js", "gr.bs.shim.js"), false, nil)
		}
	}
}

func init() {
	page.RegisterAssetDirectory(BootstrapAssets(), config.AssetPrefix+"bootstrap")
}
