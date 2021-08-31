package pagination

// LimitOffset uses limit, offset, and a slice's length to compute start and end indices for said slice.
func LimitOffset(limit, offset, length uint) (start, end uint) {
	if offset > length {
		return length, length
	}

	if limit+offset > length {
		return offset, length
	}

	return offset, offset + limit
}
