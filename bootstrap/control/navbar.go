package control

import (
	"github.com/spekary/goradd/html"
	"github.com/spekary/goradd/page"
	"goradd/app"
	localPage "goradd/page"
)

const (
	NavTabs      = "nav-tabs"
	NavPills     = "nav-pills"
	NavJustified = "nav-justified"

	NavbarHeader   = "navbar-header"
	NavbarCollapse = "navbar-collapse"
	NavbarBrand    = "navbar-brand"
	NavbarToggle   = "navbar-toggle"
	NavbarNav      = "navbar-nav"
	NavbarLeft     = "navbar-left"
	NavbarRight    = "navbar-right"
	NavbarForm     = "navbar-form"
)

type NavbarExpandClass string

const (
	NavbarExpandExtraLarge NavbarExpandClass = "navbar-expand-xl"
	NavbarExpandLarge                        = "navbar-expand-lg"
	NavbarExpandMedium                       = "navbar-expand-md"
	NavbarExpandSmall                        = "navbar-expand-sm"
	// NavbarExpandNone will always show the navbar as collapsed at any size
	NavbarExpandNone = ""
)

// NavbarCollapsedBrandPlacement controls the location of the brand when the navbar is collapsed
type NavbarCollapsedBrandPlacement int

const (
	// NavbarCollapsedBrandLeft will place the brand on the left and the toggle button on the right when collapsed
	NavbarCollapsedBrandLeft NavbarCollapsedBrandPlacement = iota
	// NavbarCollapsedBrandRight will place the brand on the right and the toggle button on the left when collapsed
	NavbarCollapsedBrandRight
	// NavbarCollapsedBrandHidden means the brand will be hidden when collapsed, and shown when expanded
	NavbarCollapsedBrandHidden
)

// Navbar is a bootstrap navbar object. Use SetText() to set the logo text of the navbar, and
// SetEscapeText() to false to turn off encoding if needed. Add child controls to populate it.
type Navbar struct {
	localPage.Control
	headerAnchor string

	style NavbarStyle
	//container ContainerClass ??
	background    BackgroundColorClass
	expand        NavbarExpandClass
	brandLocation NavbarCollapsedBrandPlacement
}

type NavbarStyle string

const (
	NavbarDark  NavbarStyle = "navbar-dark" // black on white
	NavbarLight             = "navbar-light"
)

// Creates a new standard html button
func NewNavbar(parent page.ControlI) *Navbar {
	b := &Navbar{}
	b.Tag = "nav"
	b.Init(b, parent)
	return b
}

func (b *Navbar) Init(self page.ControlI, parent page.ControlI) {
	b.Control.Init(self, parent)
	b.style = NavbarDark // default
	b.background = BackgroundColorDark
	b.expand = NavbarExpandLarge
	app.LoadBootstrap(b.Form())
}

func (b *Navbar) SetNavbarStyle(style NavbarStyle) *Navbar {
	b.style = style
	return b
}

func (b *Navbar) SetBackgroundClass(c BackgroundColorClass) *Navbar {
	b.background = c
	return b
}

func (b *Navbar) SetHeaderAnchor(a string) *Navbar {
	b.headerAnchor = a
	return b
}

// SetBrandPlacement places the brand left, right, or hidden (meaning inside the collapse area).
// The expand button location will be affected by the placement
func (b *Navbar) SetBrandPlacement(p NavbarCollapsedBrandPlacement) *Navbar {
	b.brandLocation = p
	return b
}

func (b *Navbar) DrawingAttributes() *html.Attributes {
	a := b.Control.DrawingAttributes()
	a.AddClass("navbar")
	a.AddClass(string(b.style))
	a.AddClass(string(b.expand))
	a.AddClass(string(b.background))
	a.SetDataAttribute("grctl", "bs-navbar")
	return a
}