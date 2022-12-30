package api

import sync "sync"

// As mvp, meta is stored in-memory.
// It is better to store it externally, any sql-like database would be fine.
// External storage could help us to survive server restart without losing all stat.

type metaKey struct {
	clientId ClientId
	filename string
}

func buildKey(clientId ClientId, filename string) metaKey {
	// We update file if client sends file with same name,
	// so client_id + filename pair is unique
	return metaKey{
		clientId: clientId,
		filename: filename,
	}
}

type Metadata struct {
	container map[metaKey]*FileParams
	mutex     sync.Mutex
}

type FileParams struct {
	ModifiedAt   string
	ChunksCount  int
	CurrentChunk int
	Size         int
}

func NewMetadata() Metadata {
	metadataMap := make(map[metaKey]*FileParams)
	return Metadata{metadataMap, sync.Mutex{}}
}

func (m *Metadata) IsMetaExists(clientId ClientId, filename, modified string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.container[buildKey(clientId, filename)].ModifiedAt == modified
}

func (m *Metadata) AddFile(clientId ClientId, filename, modified string, size, chunksCount int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.container[buildKey(clientId, filename)] = &FileParams{modified, chunksCount, 0, size}
}

func (m *Metadata) GetParams(clientId ClientId, filename string) *FileParams {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.container[buildKey(clientId, filename)]
}

func (m *Metadata) SetCurrentChunk(clientId ClientId, filename string, currentChunk int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if params, ok := m.container[buildKey(clientId, filename)]; ok {
		params.CurrentChunk = currentChunk
	}
}

func (m *Metadata) RemoveMeta(clientId ClientId, filename string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.container, buildKey(clientId, filename))
}
