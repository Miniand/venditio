package asset

import (
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/inject"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	DEP_ASSET                = "asset"
	DEP_THEMER               = "assetThemer"
	THEMES_PATH              = "themes"
	BUILT_PACKAGE_ASSET_PATH = ".package_assets"
)

type Resolver interface {
	AddPackagePath(p string)
	PackagePaths() []string
	ResolvePaths() []string
	ResolveTheme() string
	Resolve(asset string) []string
}

type Themer func() string

type Manager struct {
	AssPath       string
	PkgPaths      []string
	ThemeResolver func() string
}

func Register(v *core.Venditio) {
	v.BindFactory(DEP_ASSET, func(i inject.Injector) interface{} {
		a := New().(*Manager)
		a.ThemeResolver = func() string {
			themer, ok := i.Get(DEP_THEMER)
			if !ok {
				return ""
			}
			return themer.(Themer)()
		}
		return a
	})
}

func New() Resolver {
	return &Manager{
		PkgPaths: []string{},
	}
}

func (m *Manager) AddPackagePath(p string) {
	m.PkgPaths = append(m.PkgPaths, p)
}

func (m *Manager) PackagePaths() []string {
	return m.PkgPaths
}

func (m *Manager) Resolve(asset string) []string {
	result := []string{}
	for _, p := range m.ResolvePaths() {
		f := path.Join(p, asset)
		if _, err := os.Stat(f); !os.IsNotExist(err) {
			result = append(result, f)
		}
	}
	return result
}

func (m *Manager) ResolvePaths() []string {
	var resolvePaths []string
	wd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err.Error())
	}
	searchPaths := SearchPaths()
	packagePaths := m.PackagePaths()
	spLen := len(searchPaths)
	ppLen := len(packagePaths)
	// Set the theme directory as the initial one if we've got one
	theme := m.ResolveTheme()
	startOffset := 0
	if theme != "" {
		resolvePaths = make([]string, spLen*ppLen+1)
		resolvePaths[0] = path.Join(wd, THEMES_PATH, theme)
		startOffset = 1
	} else {
		resolvePaths = make([]string, spLen*ppLen)
	}
	for spIndex, sp := range searchPaths {
		for ppIndex, pp := range packagePaths {
			resolvePaths[ppLen*spIndex+ppIndex+startOffset] = path.Join(sp, pp)
		}
	}
	return resolvePaths
}

func (m *Manager) ResolveTheme() string {
	if m.ThemeResolver == nil {
		return ""
	}
	return m.ThemeResolver()
}

func SearchPaths() []string {
	wd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err.Error())
	}
	goPath := os.Getenv("GOPATH")
	goPaths := []string{}
	if goPath != "" {
		goPaths = strings.Split(goPath, string(os.PathListSeparator))
	}
	searchPaths := make([]string, len(goPaths)+1)
	searchPaths[0] = path.Join(wd, BUILT_PACKAGE_ASSET_PATH)
	for i, gp := range goPaths {
		searchPaths[i+1] = path.Join(gp, "src")
	}
	return searchPaths
}
