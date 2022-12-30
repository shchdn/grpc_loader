package api

type Metadata struct {
	//Map uuid -> filename -> FileParams
	Map map[string]map[string]FileParams
}

type FileParams struct {
	ModifiedAt   string
	ChunksCount  int
	CurrentChunk int
	Size         int
}

func NewMetadata() Metadata {
	metadataMap := make(map[string]map[string]FileParams)
	return Metadata{metadataMap}
}

func (m *Metadata) CompareModified(clientId, filename, modified string) bool {
	return m.Map[clientId][filename].ModifiedAt == modified
}

func (m *Metadata) AddFile(clientId, filename, modified string, size, chunksCount int) {
	if _, ok := m.Map[clientId]; !ok {
		m.Map[clientId] = make(map[string]FileParams)
	}
	m.Map[clientId][filename] = FileParams{modified, chunksCount, 0, size}
}

func (m *Metadata) ClientExists(clientId string) bool {
	if _, ok := m.Map[clientId]; ok {
		return true
	}
	return false
}

func (m *Metadata) FileExists(clientId, filename string) bool {
	if m.ClientExists(clientId) {
		if _, ok := m.Map[clientId][filename]; ok {
			return true
		}
	}
	return false
}

func (m *Metadata) GetCurrentChunk(clientId, filename string) int {
	return m.Map[clientId][filename].CurrentChunk
}

func (m *Metadata) GetChunksCount(clientId, filename string) int {
	return m.Map[clientId][filename].ChunksCount
}

func (m *Metadata) GetModifiedAt(clientId, filename string) string {
	return m.Map[clientId][filename].ModifiedAt
}

func (m *Metadata) IncreaseCurrentChunk(clientId, filename string, currentChunk int) {
	if params, ok := m.Map[clientId][filename]; ok {
		params.CurrentChunk = currentChunk
		m.Map[clientId][filename] = params
	}
}
