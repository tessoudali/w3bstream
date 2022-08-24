package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	g "github.com/iotexproject/Bumblebee/gen/codegen"
	"github.com/iotexproject/Bumblebee/kit/modelgen"
	"github.com/iotexproject/Bumblebee/x/pkgx"
	"github.com/saitofun/qlib/util/qnaming"
	"gopkg.in/yaml.v2"
)

type Applet struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

func (c *Applet) Schema() string {
	if c.Version == "" {
		return c.Name
	}
	return c.Name + "_" + c.Version
}

type Table struct {
	Name      string  `yaml:"name"`
	TableName string  `yaml:"table_name"`
	Comment   string  `yaml:"comment"`
	Defs      []Def   `yaml:"defs"`
	Fields    []Field `yaml:"fields"`
}

type Def struct {
	Type   string   `yaml:"type"`
	Name   string   `yaml:"name"`
	Fields []string `yaml:"fields"`
}

type Field struct {
	Name        string   `yaml:"name"`
	FieldName   string   `yaml:"field_name"`
	Type        Datatype `yaml:"type"`
	Constraints string   `yaml:"constraints"`
	Comment     string   `yaml:"comments"`
}

func (f *Field) Tag() string {
	if f.Constraints != "" {
		return fmt.Sprintf(`db:"%s,%s"`, f.FieldName, f.Constraints)
	}
	return fmt.Sprintf(`db:"%s"`, f.FieldName)
}

type Config struct {
	Applet Applet  `yaml:"applet"`
	Tables []Table `yaml:"tables"`
	files  map[string]*g.File
}

func LoadConfigFrom(path string) (*Config, error) {
	root := filepath.Dir(path)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	c := &Config{files: make(map[string]*g.File)}
	err = yaml.NewDecoder(f).Decode(c)
	if err != nil {
		return nil, err
	}

	if err := c.WriteSchema(); err != nil {
		return nil, err
	}

	if err := c.WriteTables(root); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) Schema() string { return c.Applet.Schema() }

func (c *Config) Package() string { return c.Applet.Name }

func (c *Config) File(name string) *g.File {
	if f, ok := c.files[name]; !ok {
		panic("not define table " + name)
	} else {
		return f
	}
}

func (c *Config) SnippetDefs() map[string]g.Snippet {
	snippets := make(map[string]g.Snippet)
	for _, t := range c.Tables {
		types := make([]string, 0)
		maxLenType := 0
		names := make([]string, 0)
		maxLenName := 0
		fields := make([][]string, 0)

		for _, d := range t.Defs {
			types = append(types, d.Type)
			if len(d.Type) > maxLenType {
				maxLenType = len(d.Type)
			}
			names = append(names, d.Name)
			if len(d.Name) > maxLenName {
				maxLenName = len(d.Name)
			}
			fields = append(fields, d.Fields)
		}

		for i := 0; i < len(types); i++ {
			types[i] += strings.Repeat(" ", maxLenType-len(types[i]))
			names[i] += strings.Repeat(" ", maxLenName-len(names[i]))
		}

		defs := make([]string, 0)
		for i := 0; i < len(types); i++ {
			defs = append(defs,
				fmt.Sprintf("@def %s %s %s",
					types[i], names[i], strings.Join(fields[i], " "),
				),
			)
		}
		snippets[t.Name] = g.Comments(defs...)

	}
	return snippets
}

func (c *Config) SnippetStruct() map[string]g.Snippet {
	snippets := make(map[string]g.Snippet)
	for _, t := range c.Tables {
		fields := make([]*g.SnippetField, 0)
		for _, f := range t.Fields {
			sf := (&g.SnippetField{
				Type:  f.Type.CodeGenType(c.File(t.Name)),
				Names: []*g.SnippetIdent{g.Ident(f.Name)},
			}).WithTag(f.Tag())
			if f.Comment != "" {
				sf = sf.WithOneLineComment(f.Comment)
			}
			fields = append(fields, sf)
		}

		decl := g.DeclType(g.Var(g.Struct(fields...), t.Name))
		if t.Comment != "" {
			decl = decl.WithComments(t.Name + " " + t.Comment)
		}
		snippets[t.Name] = decl
	}
	return snippets
}

func (c *Config) WriteSchema() error {
	// var DB = sqlx.NewDatabase("demo").WithSchema("applet_management")
	return nil
}

func (c *Config) WriteTables(root string) error {
	for i := range c.Tables {
		c.Tables[i].Name = qnaming.UpperCamelCase(c.Tables[i].Name)
		if c.Tables[i].TableName == "" {
			c.Tables[i].TableName = "t_" +
				qnaming.LowerSnakeCase(c.Tables[i].Name)
		}
		for j := range c.Tables[i].Fields {
			c.Tables[i].Fields[j].Name = qnaming.UpperCamelCase(
				c.Tables[i].Fields[j].Name,
			)
			if c.Tables[i].Fields[j].FieldName == "" {
				c.Tables[i].Fields[j].FieldName = "f_" +
					qnaming.LowerSnakeCase(
						c.Tables[i].Fields[j].FieldName,
					)
			}
		}
		c.files[c.Tables[i].Name] = g.NewFile(
			c.Package(),
			filepath.Join(
				root,
				qnaming.LowerSnakeCase(c.Tables[i].Name)+".go",
			),
		)
	}

	defs := c.SnippetDefs()
	typs := c.SnippetStruct()

	for name, file := range c.files {
		file.WriteSnippet(defs[name], typs[name])
		if _, err := file.Write(); err != nil {
			panic(err)
		}
	}

	cwd, _ := os.Getwd()

	for _, t := range c.Tables {
		pkg, err := pkgx.LoadFrom(root)
		if err != nil {
			return err
		}
		mg := modelgen.New(pkg)
		mg.WithComments = true
		mg.WithTableName = true
		mg.WithTableInterfaces = true
		mg.StructName = t.Name
		mg.Database = "DB"
		mg.Scan()
		mg.Output(cwd)
	}
	return nil
}
