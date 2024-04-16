package test

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	cm "github.com/campaign-manager/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestUnitCreation(t *testing.T) {
	param := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	unit := cm.NewJob("template.txt", param, "output")
	cm.Generate(*unit)
}

func TestJobCreation(t *testing.T) {
	// given a csv
	csvContent := `
key1,key2,key3,key4,key5
p1.1,p1.2,p1.3,p1.4,p1.5
p2.1,p2.2,p2.3,p2.4,p2.5
p3.1,p3.2,p3.3,p3.4,p3.5
`

	//	// given a template
	//	template := `
	//{{ .key1 }}
	//{{ .key2 }}
	//{{ .key3 }}
	//{{ .key4 }}
	//{{ .key5 }}
	//`

	// given a workspace in which the files are created
	workspaceRel := "./tmp"

	// TODO test what happens if the workspace is an abs path

	workspace, err := filepath.Abs(workspaceRel)
	assert.NoError(t, err)

	err = createWorkspace(workspace)
	assert.NoError(t, err)

	records, err := parseCsv(strings.NewReader(csvContent))
	assert.NoError(t, err)

	parameter, err := csvToMap(records)
	assert.NoError(t, err)

	for p, _ := range parameter {
		err := generateTask(parameter[p], workspace)
		assert.NoError(t, err)
	}

	tasksJsonData := struct {
		Data []map[string]string `json:"tasks"`
	}{
		parameter,
	}

	// write json with parameter
	tasksJson, err := json.MarshalIndent(tasksJsonData, "", "  ")
	assert.NoError(t, err)
	err = os.WriteFile(path.Join(workspace, "tasks.json"), tasksJson, os.ModePerm)
	assert.NoError(t, err)
}

// parseCsv reads a csv and returns
func parseCsv(in io.Reader) ([][]string, error) {
	r := csv.NewReader(in)
	records, err := r.ReadAll()
	if err != nil {
		slog.Error("Error in parsing csv", "error", err)
		return nil, err
	}

	// TODO: check if the csv provides unique parameters

	return records, nil
}

// csvToMap returns a []map[string]string from the csv records
func csvToMap(records [][]string) ([]map[string]string, error) {
	// assume that the first row contains the keys
	keys := records[0]
	var m []map[string]string

	for i, record := range records[1:] {
		// check if the record has the same length as the keys
		if len(record) != len(keys) {
			slog.Error("record has not the same length as the keys", "record length", i, "keys length", len(keys))
			return nil, fmt.Errorf("record %d has not the same length as the keys (%d)", i, len(keys))
		}

		m = append(m, make(map[string]string))
		for j, value := range record {
			m[i][keys[j]] = value
		}
	}
	return m, nil
}

func createWorkspace(workspace string) error {
	// check if workspace already exists. If so, report an error.
	if folderExists(workspace) {
		slog.Error("Workspace already exists", "workspace", workspace)
		return fmt.Errorf("workspace %s already exists", workspace)
		// TODO return full folder path in the error message.
	}

	// create workspace
	err := os.MkdirAll(workspace, os.ModePerm)
	if err != nil {
		slog.Error("Error in creating workspace", "workspace", workspace, "error", err)
	}

	return err
}

func folderExists(folderName string) bool {
	_, err := os.Stat(folderName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func uniqueId() (string, error) {
	const length = 16

	// Generate a random value
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		slog.Error("Error in generating random bytes", "error", err)
		return "", err
	}

	// Get current timestamp
	currentTime := time.Now().UnixNano()

	// Concatenate random value and timestamp
	input := fmt.Sprintf("%s%d", randomBytes, currentTime)

	// Hash the input using SHA-256
	hash := sha256.New()
	hash.Write([]byte(input))
	hashedBytes := hash.Sum(nil)

	// Encode the hashed bytes into a shorter representation
	//shortID := base64.URLEncoding.EncodeToString(hashedBytes)
	shortID := hex.EncodeToString(hashedBytes)

	// Truncate the short ID to a desired length
	// For example, you can take the first 8 characters
	shortID = shortID[:length]

	return shortID, nil
}

func generateTask(parameter map[string]string, workspace string) error {
	id, err := uniqueId()
	if err != nil {
		slog.Error("Error in generating unique id", "error", err)
		return err
	}

	// try to add the id to the parameter
	_, ok := parameter["id"]
	if ok {
		slog.Error("Parameter already contains id", "parameter", parameter)
		return fmt.Errorf("parameter already contains id")
	}
	parameter["id"] = id

	// create task directory
	dir := path.Join(workspace, id)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		slog.Debug("Error in creating folder", "dir", dir, "error", err)
		panic(err)
	}

	_, ok = parameter["dir"]
	if ok {
		slog.Error("Parameter already contains dir", "parameter", parameter)
		return fmt.Errorf("parameter already contains dir")
	}
	parameter["dir"] = dir

	// write json with parameters
	jsonData, err := json.MarshalIndent(parameter, "", "  ")
	err = os.WriteFile(path.Join(dir, "parameter.json"), jsonData, os.ModePerm)
	if err != nil {
		slog.Error("Error in writing parameter.json", "error", err)
		return err
	}

	return nil
}
