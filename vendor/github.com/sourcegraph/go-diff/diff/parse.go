	"path/filepath"
	// FileDiff is added/deleted file
	// No further collection of hunks needed
	if fd.NewName == "" {
		return fd, nil
	}

		fd.OrigTime = origTime
		fd.NewTime = newTime
// timestamps). Or which starts with "Only in " with dir path and filename.
// "Only in" message is supported in POSIX locale: https://pubs.opengroup.org/onlinepubs/9699919799/utilities/diff.html#tag_20_34_10
	if r.fileHeaderLine != nil {
		if isOnlyMessage, source, filename := parseOnlyInMessage(r.fileHeaderLine); isOnlyMessage {
			return filepath.Join(string(source), string(filename)),
				"", nil, nil, nil
		}
	}

	return fmt.Sprintf("overflowed into next file: %s", string(e))
		// Reached message that file is added/deleted
		if isOnlyInMessage, _, _ := parseOnlyInMessage(line); isOnlyInMessage {
			r.fileHeaderLine = line // pass to readOneFileHeader (see fileHeaderLine field doc)
			return xheaders, nil
		}

	case lineCount == 6 && strings.HasPrefix(fd.Extended[5], "Binary files ") && strings.HasPrefix(fd.Extended[2], "rename from ") && strings.HasPrefix(fd.Extended[3], "rename to "):
		names := strings.SplitN(fd.Extended[0][len("diff --git "):], " ", 2)
		fd.OrigName = names[0]
		fd.NewName = names[1]
		return true

	// ErrBadOnlyInMessage is when a file have a malformed `only in` message
	// Should be in format `Only in {source}: {filename}`
	ErrBadOnlyInMessage = errors.New("bad 'only in' message")
			if len(line) >= 1 && (!linePrefix(line[0]) || bytes.HasPrefix(line, []byte("--- "))) {
// parseOnlyInMessage checks if line is a "Only in {source}: {filename}" and returns source and filename
func parseOnlyInMessage(line []byte) (bool, []byte, []byte) {
	if !bytes.HasPrefix(line, onlyInMessagePrefix) {
		return false, nil, nil
	}
	line = line[len(onlyInMessagePrefix):]
	idx := bytes.Index(line, []byte(": "))
	if idx < 0 {
		return false, nil, nil
	}
	return true, line[:idx], line[idx+2:]
}
