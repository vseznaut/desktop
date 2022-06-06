package main

const dotCharacter = 46

// Hidden is a struct that contains the hidden message and the key.
func isHidden(path string) (bool, error) {
	// dotfiles also count as hidden (if you want)
	if path[0] == dotCharacter {
		return true, nil
	}

	return false, nil
}
