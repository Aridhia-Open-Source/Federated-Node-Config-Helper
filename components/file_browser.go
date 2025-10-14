package components

import (
	"os"
	"path/filepath"

	"github.com/rivo/tview"
)

// local
var fileTree = tview.NewList()
var fileList = tview.NewList()

// Export
var FileBrowserContainer = tview.NewFlex().
	AddItem(fileTree, 0, 2, false).
	AddItem(fileList, 0, 8, true)

func init() {
	fileTree.SetBorder(true)
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fileList.SetBorder(true)
	FileBrowserContainer.SetBorder(true).SetTitle("File Browser")
	UpdateListsFolders(cwd)
	UpdateListsFiles(cwd, listDirContents(cwd))
}

func UpdateListsFiles(currentPath string, entries []os.DirEntry) {
	currentPath = filepath.Clean(currentPath)
	// Clear previous file list
	for idx := range fileList.GetItemCount() {
		fileList.RemoveItem(idx)
	}
	// Update the folder list with the new path
	UpdateListsFolders(currentPath)
	// Update file list with the new path
	for _, file := range entries {
		fpath := filepath.Join(currentPath, file.Name())
		if !file.IsDir() {
			fileList.AddItem(file.Name(), "", '0', func() {
				FilepathInput.SetText(fpath)
			})
		}
	}
}

func UpdateListsFolders(currentPath string) {
	currentPath = filepath.Clean(currentPath)
	entries := listDirContents(currentPath)
	// Clear previous folder list
	for idx := range fileTree.GetItemCount() {
		fileTree.RemoveItem(idx)
	}
	fileTree.AddItem("..", "", 127, func() {
		parent := filepath.Dir(currentPath)
		UpdateListsFiles(parent, listDirContents(parent))
	})
	// Populate with the dirs in the new folder
	for _, dir := range entries {
		if dir.IsDir() {
			fpath := filepath.Join(currentPath, dir.Name())
			fileTree.AddItem(dir.Name(), "", '0', func() {
				UpdateListsFiles(fpath, listDirContents(fpath))
			})
		}
	}
}

func listDirContents(path string) []os.DirEntry {
	// Get list files
	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	return entries
}
