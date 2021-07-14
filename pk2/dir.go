package pk2

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Directory struct {
	Entries           []PackFileEntry
	EntriesByName     map[string]PackFileEntry
	Directories       []Directory
	DirectoriesByName map[string]Directory
	Name              string
}

// Be careful, this might eat your hard disk space. Very IO intensive
func (d *Directory) PrintFiles() {
	for i := range d.Entries {
		log.Printf("%s/%s", d.Name, d.Entries[i].Name)
	}

	for i := range d.Directories {
		d.Directories[i].PrintFiles()
	}
}

func (d *Directory) Extract(outputDir string, reader Pk2Reader, counter *int, maxFiles int) {
	err := os.Mkdir(outputDir+string(os.PathSeparator)+d.Name, os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range d.Entries {
		if v.Type == TypeFile {
			outputFile := outputDir + string(os.PathSeparator) + d.Name + string(os.PathSeparator) + v.Name
			extractFile(reader, v, outputFile)
			*counter++
		}
	}
	fmt.Printf("\rExtracted [%d / %d]", *counter, maxFiles)

	for _, v := range d.Directories {
		v.Extract(outputDir, reader, counter, maxFiles)
	}
}

func extractFile(reader Pk2Reader, entry PackFileEntry, outputFile string) {
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(reader.ReadEntry(&entry))
	f.Close()
}

func (d *Directory) TotalFiles() int {
	totalFiles := 0
	for i := range d.Entries {
		if d.Entries[i].Type == TypeFile {
			totalFiles++
		}
	}

	for i := range d.Directories {
		totalFiles += d.Directories[i].TotalFiles()
	}

	return totalFiles
}

func (d *Directory) buildEntryMap() {
	if d.EntriesByName == nil {
		d.EntriesByName = make(map[string]PackFileEntry)
	}
	for _, entry := range d.Entries {
		d.EntriesByName[entry.Name] = entry
	}
}

func (d *Directory) buildDirectoryMap() {
	if d.DirectoriesByName == nil {
		d.DirectoriesByName = make(map[string]Directory)
	}

	for _, dir := range d.Directories {
		dir.buildEntryMap()
		dir.buildDirectoryMap()
		d.DirectoriesByName[dir.Name] = dir
	}
}

func (d *Directory) getFile(reader *Pk2Reader, filename string) (data []byte, err error) {
	filename = strings.ReplaceAll(filename, "\\", string(os.PathSeparator))
	filename = strings.ReplaceAll(filename, "/", string(os.PathSeparator))
	if !strings.HasPrefix(filename, d.Name) {
		return nil, errors.New(fmt.Sprintf("file path [%s] does not contain directory [%s]", filename, d.Name))
	}

	pathParts := strings.Split(filename, string(os.PathSeparator))
	fileName := pathParts[len(pathParts)-1]

	if entry, ok := d.EntriesByName[fileName]; ok {
		data = reader.ReadEntry(&entry)
	} else {
		for dirName, dir := range d.DirectoriesByName {
			if strings.HasPrefix(filename, dirName) {
				return dir.getFile(reader, filename)
			}
		}
	}

	if len(data) > 0 {
		return data, nil
	} else {
		return nil, errors.New("no data found for " + filename)
	}
}
