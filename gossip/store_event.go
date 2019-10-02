package gossip

import (
	"bytes"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
)

// DeleteEvent deletes event.
func (s *Store) DeleteEvent(epoch idx.Epoch, id hash.Event) {
	key := id.Bytes()

	err := s.table.Events.Delete(key)
	if err != nil {
		s.Log.Crit("Failed to delete key", "err", err)
	}
	s.DelEventHeader(epoch, id)
}

// SetEvent stores event.
func (s *Store) SetEvent(e *inter.Event) {
	key := e.Hash().Bytes()

	s.set(s.table.Events, key, e)
	s.SetEventHeader(e.Epoch, e.Hash(), &e.EventHeaderData)
}

// GetEvent returns stored event.
func (s *Store) GetEvent(id hash.Event) *inter.Event {
	key := id.Bytes()

	w, _ := s.get(s.table.Events, key, &inter.Event{}).(*inter.Event)
	return w
}

func getPrefix(epoch idx.Epoch, lamport idx.Lamport, hashPrefix []byte) []byte {
	buf := bytes.NewBuffer(epoch.Bytes())
	buf.Write(lamport.Bytes())
	buf.Write(hashPrefix)
	return buf.Bytes()
}

func (s *Store) FindEventHashes(epoch idx.Epoch, lamport idx.Lamport, hashPrefix []byte) hash.Events {
	prefix := getPrefix(epoch, lamport, hashPrefix)
	res := make(hash.Events, 0, 10)

	it := s.table.Events.NewIteratorWithPrefix(prefix)
	defer it.Release()
	for it.Next() {
		res = append(res, hash.BytesToEvent(it.Key()))
	}

	return res
}

// GetEventRLP returns stored event. Serialized.
func (s *Store) GetEventRLP(id hash.Event) rlp.RawValue {
	key := id.Bytes()

	data, err := s.table.Events.Get(key)
	if err != nil {
		s.Log.Crit("Failed to get key-value", "err", err)
	}
	return data
}

// HasEvent returns true if event exists.
func (s *Store) HasEvent(h hash.Event) bool {
	return s.has(s.table.Events, h.Bytes())
}
