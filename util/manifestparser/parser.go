package manifestparser

import (
	"errors"
	"io/ioutil"

	"github.com/cloudfoundry/bosh-cli/director/template"
	yaml "gopkg.in/yaml.v2"
)

type Parser struct {
	PathToManifest string

	Applications []Application

	rawManifest []byte
}

func NewParser() *Parser {
	return new(Parser)
}

func (parser *Parser) Parse(manifestPath string) error {
	bytes, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return err
	}
	return parser.parse(bytes)
}

// InterpolateAndParse reads the manifest at the provided paths, interpolates
// variables if a vars file is provided, and sets the current manifest to the
// resulting manifest.
func (parser *Parser) InterpolateAndParse(pathToManifest string, pathsToVarsFiles []string, vars []template.VarKV) error {
	rawManifest, err := ioutil.ReadFile(pathToManifest)
	if err != nil {
		return err
	}

	tpl := template.NewTemplate(rawManifest)
	fileVars := template.StaticVariables{}

	for _, path := range pathsToVarsFiles {
		rawVarsFile, ioerr := ioutil.ReadFile(path)
		if ioerr != nil {
			return ioerr
		}

		var sv template.StaticVariables

		err = yaml.Unmarshal(rawVarsFile, &sv)
		if err != nil {
			return InvalidYAMLError{Err: err}
		}

		for k, v := range sv {
			fileVars[k] = v
		}
	}

	for _, kv := range vars {
		fileVars[kv.Name] = kv.Value
	}

	rawManifest, err = tpl.Evaluate(fileVars, nil, template.EvaluateOpts{ExpectAllKeys: true})
	if err != nil {
		return InterpolationError{Err: err}
	}

	parser.PathToManifest = pathToManifest
	return parser.parse(rawManifest)
}

func (parser Parser) AppNames() []string {
	var names []string
	for _, app := range parser.Applications {
		names = append(names, app.Name)
	}
	return names
}

func (parser Parser) FullRawManifest() []byte {
	return parser.rawManifest
}

func (parser Parser) RawManifest(_ string) ([]byte, error) {
	return parser.rawManifest, nil
}

func (parser *Parser) parse(manifestBytes []byte) error {
	parser.rawManifest = manifestBytes

	var raw struct {
		Applications []Application `yaml:"applications"`
	}

	err := yaml.Unmarshal(manifestBytes, &raw)
	if err != nil {
		return err
	}

	parser.Applications = raw.Applications

	if len(parser.Applications) == 0 {
		return errors.New("must have at least one application")
	}

	for _, application := range parser.Applications {
		if application.Name == "" {
			return errors.New("Found an application with no name specified")
		}
	}

	return nil
}
