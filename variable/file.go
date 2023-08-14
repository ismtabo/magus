package variable

import (
	"encoding/json"
	"strings"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/fs"
	go_errors "github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func FromFile(ctx context.Context, path string) (Variables, error) {
	if !strings.HasSuffix(path, ".json") && !strings.HasSuffix(path, ".yaml") {
		// TODO: wrap error
		return nil, go_errors.Errorf("invalid variables file %s", path)
	}
	file, err := fs.ReadFile(ctx, path)
	if err != nil {
		// TODO: wrap error
		return nil, err
	}
	tmpVars := map[string]any{}
	if err := unmarshall(file, &tmpVars); err != nil {
		// TODO: wrap error
		return nil, err
	}
	vars := Variables{}
	for k, v := range tmpVars {
		vars = append(vars, NewLiteralVariable(k, v))
	}
	return vars, nil
}

func unmarshall(f file.File, v interface{}) error {
	if strings.HasSuffix(f.Path(), ".json") {
		return unmarshalJSON(f.Bytes(), v)
	}
	return unmarshalYAML(f.Bytes(), v)
}

func unmarshalJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func unmarshalYAML(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
