package state

import (
	"fmt"
	"hash/fnv"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"

	"github.com/DylanGuedes/laser/files"
	"golang.org/x/exp/slog"
	yaml "gopkg.in/yaml.v3"
)

type SyncedFile struct {
	ID   string
	Path string
	Hash string
}

func (s *SyncedFile) MkCopy(scope string) {
	in, err := os.Open(s.Path)
	defer in.Close()
	if err != nil {
		slog.Error("couldn't open file", "path", s.Path, "scope", scope)
		panic(err)
	}

	o, err := os.Create(path.Join(scope, files.DefautSyncedFolder, s.Hash))
	if err != nil {
		panic(err)
	}
	defer o.Close()

	if _, err := io.Copy(o, in); err != nil {
		panic(err)
	}
}

func computeHash(key string) string {
	h := fnv.New64a()
	_, err := h.Write([]byte(key))
	if err != nil {
		panic(err)
	}

	computedHash := h.Sum64()
	return strconv.Itoa(int(computedHash))
}

func NewSyncedFile(id, p string) SyncedFile {
	return SyncedFile{
		ID:   id,
		Path: p,
		Hash: computeHash(id),
	}
}

type StaticVar struct {
	Name  string
	Value string
}

type DynamicVar struct {
	Name  string
	Value string
}

type L struct {
	Scope       string       `yaml:"scope"`
	SyncedFiles []SyncedFile `yaml:"synced_files"`
	StaticVars  []StaticVar  `yaml:"static_vars"`
	DynamicVars []DynamicVar `yaml:"dynamic_vars"`
}

func (laser *L) AddSyncFile(f SyncedFile) {
	laser.SyncedFiles = append(laser.SyncedFiles, f)
}

func (laser *L) StoreState() {
	slog.Info("saving state")

	filename := path.Join(laser.Scope, files.Config)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := yaml.NewEncoder(f)
	enc.SetIndent(2)
	if err := enc.Encode(laser); err != nil {
		panic(err)
	}
	defer enc.Close()

	slog.Info("state successfully stored")
}

func LoadState(scope string) *L {
	if files.Available(scope) {
		return loadStateFromFile(path.Join(scope, files.Config))
	}
	return initState()
}

func loadStateFromFile(f string) *L {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		panic(fmt.Errorf("couldn't load state from file: %w", err))
	}

	var config L
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	return &config
}

func initState() *L {
	slog.Info("initializing Laser")
	return &L{
		Scope: ".laser",
	}
}
