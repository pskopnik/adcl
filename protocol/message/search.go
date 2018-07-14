package message

type SearchTermAction int

const (
	SearchTermUndefined SearchTermAction = iota
	SearchTermInclude
	SearchTermExclude
	SearchTermExtension
)

type SearchTerm struct {
	TermAction SearchTermAction
	Term       string
}

// SearchTermIterator is utility to read only terms of a []SearchTerm slice
// with a specific SearchTermAction (Filter).
//
// Meant to be used like this:
//
//     it := SearchTermIterator{
//         SearchTerms: searchTerms,
//         Filter:      SearchTermInclude,
//     }
//     for it.Next() {
//         searchTerm := it.SearchTerm()
//         // some processing using searchTerm
//     }
type SearchTermIterator struct {
	// SearchTerms is the []SearchTerm slice the iterator iterates over.
	SearchTerms []SearchTerm
	// Filter defines an action, only search terms with this action will be
	// returned by the iterator. SearchTermUndefined (the default) returns all
	// search terms.
	Filter      SearchTermAction
	initialised bool
	pos         int
}

// Next advances the iterator to the next term matching its Filter.
// If such a term exists, Next() returns true and the term may be retrieved
// using SearchTerm().
func (s *SearchTermIterator) Next() bool {
	if s.initialised {
		s.pos++
	} else {
		s.initialised = true
	}

	if s.Filter != SearchTermUndefined {
		for s.pos < len(s.SearchTerms) {
			if s.SearchTerms[s.pos].TermAction == s.Filter {
				break
			}
			s.pos++
		}
	}

	if s.pos >= len(s.SearchTerms) {
		return false
	}

	return true
}

// SearchTerm returns the SearchTerm currently pointed to by the iterator.
// If Next() has not been called yet or all terms have been read, SearchTerm()
// returns nil.
func (s *SearchTermIterator) SearchTerm() *SearchTerm {
	if !s.initialised || s.pos >= len(s.SearchTerms) {
		return nil
	}

	return &s.SearchTerms[s.pos]
}
