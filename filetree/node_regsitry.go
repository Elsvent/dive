package filetree

import (
	"github.com/google/uuid"
)

var NodeRegistry = newFileNodeRegistry()

func newFileNodeRegistry() *FileNodeRegistry {
	return &FileNodeRegistry{
		Nodes: make(map[uuid.UUID]FileNode),
	}
}

func (registry *FileNodeRegistry) Set(node FileNode) {
	registry.Nodes[node.Id] = node
}

func (registry *FileNodeRegistry) Get(id uuid.UUID) *FileNode {
	node, _ := registry.Nodes[id]
	return &node
}
