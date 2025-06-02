package cueparser

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var fileLineRegex = regexp.MustCompile(`^FILE\s+"(.+)"\s+(\S+)$`)

type stateFn func(*parser, string) stateFn

type parser struct {
	sheet        *CueSheet
	currentFile  *File
	currentTrack *Track
	state        stateFn
}

func Parse(r io.Reader) (*CueSheet, error) {
	p := &parser{
		sheet: &CueSheet{},
		state: stateStart,
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		p.state = p.state(p, line)
	}
	if p.currentFile != nil {
		if p.currentTrack != nil {
			p.currentFile.Tracks = append(p.currentFile.Tracks, *p.currentTrack)
		}
		p.sheet.Files = append(p.sheet.Files, *p.currentFile)
	}
	return p.sheet, scanner.Err()
}

func stateStart(p *parser, line string) stateFn {
	switch {
	case strings.HasPrefix(line, "TITLE"):
		p.sheet.Title = extractValue(line)
	case strings.HasPrefix(line, "PERFORMER"):
		p.sheet.Performer = extractValue(line)
	case strings.HasPrefix(line, "FILE"):
		matches := fileLineRegex.FindStringSubmatch(line)
		if matches == nil {
			panic("invalid FILE line")
		}
		p.currentFile = &File{
			Name: strings.Trim(matches[1], "\""),
			Type: matches[2],
		}
		return stateFile
	}
	return stateStart
}

func stateFile(p *parser, line string) stateFn {
	switch {
	case strings.HasPrefix(line, "TRACK"):
		parts := strings.SplitN(line, " ", 3)
		num, _ := strconv.Atoi(parts[1])
		p.currentTrack = &Track{
			Number: num,
		}
		return stateTrack
	}
	return stateFile
}

func stateTrack(p *parser, line string) stateFn {
	switch {
	case strings.HasPrefix(line, "TITLE"):
		p.currentTrack.Title = extractValue(line)
	case strings.HasPrefix(line, "PERFORMER"):
		p.currentTrack.Performer = extractValue(line)
	case strings.HasPrefix(line, "INDEX"):
		parts := strings.SplitN(line, " ", 3)
		num, _ := strconv.Atoi(parts[1])
		time := parts[2]
		p.currentTrack.Indexes = append(p.currentTrack.Indexes, Index{
			Number: num,
			Time:   time,
		})
	case strings.HasPrefix(line, "TRACK"):
		// finalize previous track
		p.currentFile.Tracks = append(p.currentFile.Tracks, *p.currentTrack)
		num, _ := strconv.Atoi(strings.Fields(line)[1])
		p.currentTrack = &Track{Number: num}
		return stateTrack
	case strings.HasPrefix(line, "FILE"):
		// finalize current track and file
		p.currentFile.Tracks = append(p.currentFile.Tracks, *p.currentTrack)
		p.sheet.Files = append(p.sheet.Files, *p.currentFile)

		parts := strings.SplitN(line, " ", 3)
		p.currentFile = &File{
			Name: strings.Trim(parts[1], "\""),
			Type: parts[2],
		}
		p.currentTrack = nil
		return stateFile
	}
	return stateTrack
}

func extractValue(line string) string {
	idx := strings.Index(line, " ")
	if idx < 0 {
		return ""
	}
	return strings.Trim(line[idx+1:], "\" ")
}
