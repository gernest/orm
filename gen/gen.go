package gen

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/posener/orm/dialect"
	"github.com/posener/orm/gen/b0x"
	"github.com/posener/orm/load"
)

//go:generate fileb0x b0x.yml
//go:generate sed -i 1d b0x/ab0x.go

var header = template.Must(template.New("header").Parse(`
// Code generated by github.com/posener/orm; DO NOT EDIT
//
// ORM functions for type {{$.Name}}

package {{$.Package}}
`))

// TemplateData arguments for the templates
type TemplateData struct {
	// The name	of the new created package
	Type     *load.Type
	Dialects []dialect.Generator
	Public   string
	Private  string
}

var templates = template.New("").Funcs(template.FuncMap{
	"plus1": func(x int) int { return x + 1 },
})

func init() {
	files, err := b0x.WalkDirs(".", false, "")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if !strings.HasSuffix(file, ".tmpl") {
			continue
		}
		data, err := b0x.ReadFile(file)
		if err != nil {
			panic(err)
		}
		templates.New(file).Parse(string(data))
	}
}

// Gen generates all the ORM files for a given struct in a given package.
// st is the type descriptor of the struct
func Gen(tp *load.Type) error {
	// get the package ormDir on disk
	structPkgDir, err := packagePath(tp.ImportPath)
	if err != nil {
		return err
	}

	dialects := dialect.NewGen()

	data := TemplateData{
		Type:     tp,
		Dialects: dialects,
		Public:   tp.Name,
		Private:  strings.ToLower(tp.Name),
	}

	ormFileName := strings.ToLower(tp.Name + "_orm.go")
	ormFilePath := filepath.Join(structPkgDir, ormFileName)

	log.Printf("Generating code for %s into %s", tp, ormFilePath)

	ormFile, err := os.Create(ormFilePath)
	if err != nil {
		return fmt.Errorf("creating file %s: %s", ormFilePath, err)
	}

	// write header
	err = header.Execute(ormFile, tp)
	if err != nil {
		return err
	}

	// write templates
	for _, name := range templateNames() {
		_, err = ormFile.WriteString(fmt.Sprintf("\n\n// ========= Template %s ===================\n\n", name))
		if err != nil {
			return err
		}
		err = templates.ExecuteTemplate(ormFile, name, data)
		if err != nil {
			return err
		}
	}
	format(ormFilePath)
	return nil
}

// templateNames return sorted list of template names
func templateNames() []string {
	tpls := templates.Templates()
	names := make([]string, 0, len(tpls))
	for _, tmpl := range tpls {
		names = append(names, tmpl.Name())
	}
	sort.Strings(names)
	return names
}

func packagePath(pkg string) (string, error) {
	for _, gopath := range filepath.SplitList(os.Getenv("GOPATH")) {
		pkgPath := filepath.Join(gopath, "src", pkg)
		f, err := os.Stat(pkgPath)
		if err == nil && f.IsDir() {
			return pkgPath, nil
		}
	}
	return "", fmt.Errorf("package path was not found: %s", pkg)
}

func format(path string) {
	_, err := exec.Command("goimports", "-w", path).CombinedOutput()
	if err != nil {
		log.Printf("Failed formatting package: %s", err)
	}
}
