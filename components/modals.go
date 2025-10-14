package components

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/rivo/tview"
)

var FileChoice = tview.NewModal()
var warningText = tview.NewTextView().SetText("Warning: Importing a yaml will not populate or look into existing secrets/configmaps. Once imported those fields will be left empty")

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	FilepathInput.SetLabel("File path")
	FilepathInput.SetText(cwd)
	FilepathInput.SetChangedFunc(func(path string) {
		// Check if the path exists, and update the file list
		entries, err := os.ReadDir(path)
		if err != nil {
			return
		}
		UpdateListsFolders(path)
		UpdateListsFiles(path, entries)
	})
	FilepathInput.SetAutocompleteFunc(func(currentPath string) (entries []string) {
		var autoCompleteList []string
		if currentPath == "" {
			return autoCompleteList
		}
		parent := filepath.Dir(currentPath)
		for _, entity := range listDirContents(parent) {
			entityToMatch := filepath.Base(currentPath)
			regPath := fmt.Sprintf("%s.*", entityToMatch)
			matched, _ := regexp.MatchString(regPath, entity.Name())
			endSlashMatched, _ := regexp.MatchString("(/|\\\\){1}$", currentPath)
			if matched || endSlashMatched {
				autoCompleteList = append(autoCompleteList, filepath.Join(parent, entity.Name()))
			}
		}
		return autoCompleteList
	})
	CancelModalButton.SetBorder(true)
	ConfirmButton.SetBorder(true)
}

var ConfirmButton = tview.NewButton("Load")
var CancelModalButton = tview.NewButton("Cancel")

var FilepathInput = tview.NewInputField()

var modal = tview.NewFlex().SetDirection(tview.FlexRow).
	AddItem(warningText, 0, 2, true).
	AddItem(FileBrowserContainer, 0, 5, true).
	AddItem(FilepathInput, 0, 1, true).
	AddItem(ConfirmButton, 0, 1, false).
	AddItem(CancelModalButton, 0, 1, false)

var ModalPage = tview.NewPages().
	AddPage("modal", modal, true, true)
