package gen

import (
	"bufio"
	"bytes"
	"fmt"
	"go/types"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/posener/orm/gen/b0x"
)

//go:generate fileb0x b0x.yml

const suffix = "orm"

const header = `// Autogenerated by github.com/posener/orm
`

type TplProps struct {
	// Table is the table name of the given struct
	Table string
	// The name	of the new created package
	PackageName string
	// Type describes the type of the given struct to generate code for
	Type Type
}

var templates = template.New("").Funcs(template.FuncMap{
	"plus1": func(x int) int { return x + 1 },
})

var commonFiles []string

func init() {
	files, err := b0x.WalkDirs(".", false, "")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		switch {
		case strings.HasSuffix(file, ".go.tpl"):
			data, err := b0x.ReadFile(file)
			if err != nil {
				panic(err)
			}
			templates.New(file[:len(file)-4]).Parse(string(data)) // remove ".tpl" suffix
		case strings.HasSuffix(file, ".go"):
			commonFiles = append(commonFiles, file)
		}
	}
}

// Gen generates all the ORM files for a given struct in a given package.
// pkg is the package of the struct
// st is the type descriptor of the struct
// name is the name of the type of the struct
func Gen(pkg *types.Package, st *types.Struct, name string) error {
	// get the package ormDir on disk
	pkgDir, err := packagePath(pkg.Path())
	if err != nil {
		return err
	}

	// the new created package name is the name of the struct with "orm" suffix
	ormPkgName := strings.ToLower(name + suffix)

	// the files will be generated in a sub package
	ormDir := filepath.Join(pkgDir, ormPkgName)
	log.Printf("Generating code to directory: %s", ormDir)
	if err = os.MkdirAll(ormDir, 0775); err != nil {
		return fmt.Errorf("creating directory %s: %s", ormDir, err)
	}

	props := TplProps{
		Table:       strings.ToLower(name),
		Type:        NewType(name, pkg, st),
		PackageName: ormPkgName,
	}
	log.Printf("Template configuration: %+v", props)

	for _, tpl := range templates.Templates() {
		err := writeTemplate(tpl, props, ormDir)
		if err != nil {
			return err
		}
	}

	if err = copyFixedFiles(ormDir, ormPkgName); err != nil {
		return fmt.Errorf("copy fixed files: %s", err)
	}
	format(ormDir)
	return nil
}

func copyFixedFiles(dir, pkgName string) error {
	// copy fixed go files
	for _, srcName := range commonFiles {
		fileName := filepath.Base(srcName)
		dstName := filepath.Join(dir, fileName)
		err := copyFixedFile(srcName, dstName, pkgName)
		if err != nil {
			return fmt.Errorf("copy file %s to %s: %s", srcName, dstName, err)
		}
	}
	return nil

}

func packagePath(pkg string) (string, error) {
	for _, gopath := range filepath.SplitList(os.Getenv("GOPATH")) {
		pkgPath := filepath.Join(gopath, "src", pkg)
		f, err := os.Stat(pkgPath)
		if err == nil && f.IsDir() {
			return pkgPath, nil
		}
	}
	return "", fmt.Errorf("package path was not found")
}

func writeTemplate(tpl *template.Template, props TplProps, dir string) error {
	// remove the ".tpl" suffix
	fileName := tpl.Name()
	filePath := filepath.Join(dir, fileName)
	log.Printf("Writing file: %s", filePath)
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("creating file %s: %s", filePath, err)
	}
	defer f.Close()
	f.Write([]byte(header))
	return tpl.Execute(f, props)
}

func copyFixedFile(srcName, dstName, pkgName string) error {
	log.Printf("Copy file: %s", dstName)

	// prepare dst writer
	dst, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("creating dst: %s", err)
	}
	defer dst.Close()
	w := bufio.NewWriter(dst)
	defer w.Flush()

	// prepare src reader
	srcData, err := b0x.ReadFile(srcName)
	if err != nil {
		return fmt.Errorf("reading src: %s", err)
	}
	src := bufio.NewReader(bytes.NewBuffer(srcData))
	r := bufio.NewReader(src)

	if err := discardPackageLine(r); err != nil {
		return fmt.Errorf("discarding package line from src: %s", err)
	}
	if err := writePackageLine(w, pkgName); err != nil {
		return fmt.Errorf("writing package line to dst: %s", err)
	}
	if _, err := r.WriteTo(w); err != nil {
		return fmt.Errorf("copy content: %s", err)
	}
	w.Flush()
	return nil
}

func discardPackageLine(r *bufio.Reader) error {
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			return err
		}
		if bytes.HasPrefix(b, []byte("package ")) {
			return nil
		}
	}
}

func writePackageLine(w *bufio.Writer, pkgName string) error {
	_, err := w.WriteString(header + "package " + pkgName)
	return err
}

func format(dir string) {
	_, err := exec.Command("gofmt", "-s", "-w", dir).CombinedOutput()
	if err != nil {
		log.Printf("Failed formatting package: %s", err)
	}
}
