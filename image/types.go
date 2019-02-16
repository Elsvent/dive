package image

import (
	"io"

	"github.com/docker/docker/client"
	"github.com/wagoodman/dive/filetree"
)

type Parser interface {
}

type Analyzer interface {
	Fetch() (io.ReadCloser, error)
	Parse(io.ReadCloser) error
	Analyze() (*AnalysisResult, error)
}

type Layer interface {
	Id() string
	ShortId() string
	Index() int
	Command() string
	Size() uint64
	Tree() *filetree.FileTree
	String() string
}

type AnalysisResult struct {
	ID                string
	Layers            []Layer
	RefTrees          []*filetree.FileTree
	Efficiency        float64
	SizeBytes         uint64
	UserSizeByes      uint64  // this is all bytes except for the base image
	WastedUserPercent float64 // = wasted-bytes/user-size-bytes
	WastedBytes       uint64
	Inefficiencies    filetree.EfficiencySlice
}

type dockerImageAnalyzer struct {
	ID        string
	Client    *client.Client
	JsonFiles map[string][]byte
	Trees     []*filetree.FileTree
	LayerMap  map[string]*filetree.FileTree
	Layers    []*dockerLayer
}

type dockerImageHistoryEntry struct {
	ID         string
	Size       uint64
	Created    string `json:"created"`
	Author     string `json:"author"`
	CreatedBy  string `json:"created_by"`
	EmptyLayer bool   `json:"empty_layer"`
}

type dockerImageManifest struct {
	ConfigPath    string   `json:"Config"`
	RepoTags      []string `json:"RepoTags"`
	LayerTarPaths []string `json:"Layers"`
}

type dockerImageConfig struct {
	History []dockerImageHistoryEntry `json:"history"`
	RootFs  dockerRootFs              `json:"rootfs"`
}

type dockerRootFs struct {
	Type    string   `json:"type"`
	DiffIds []string `json:"diff_ids"`
}

// Layer represents a Docker image layer and metadata
type dockerLayer struct {
	TarPath  string
	History  dockerImageHistoryEntry
	RefIndex int
	RefTree  *filetree.FileTree
}
