package builtin

import (
	"fmt"
	"os"

	"github.com/zillolo/bim"
)

// Init creates the filesystem structure for a repository.
func Init(name string) {
	workingDir, err := os.Getwd()
	if err != nil {
		// What happened here? lol (NOTE: My editor tries to correct lol to LoadBlob, lol)
		fmt.Printf("Error during...wtf?: %v\n", err)
		return
	}

	// Create project folder
	err = os.Mkdir(name, 0755)
	if err != nil {
		fmt.Printf("Error during creation of project folder: %v\n", err)
		return
	}

	err = os.Chdir(name)
	if err != nil {
		fmt.Printf("Error during changing of directory: %v\n", err)
		return
	}
	// Create the repository folder.
	err = os.Mkdir(bim.RepoDir, 0755)
	if err != nil {
		fmt.Printf("Error during creation of repository folder: %v\n", err)
		return
	}

	err = os.Chdir(bim.RepoDir)
	if err != nil {
		fmt.Printf("Error during changing of directory: %v\n", err)
		return
	}

	err = os.Mkdir(bim.BlobDir, 0755)
	if err != nil {
		fmt.Printf("Error during creation of blob directory: %v\n", err)
		return
	}

	err = os.Mkdir(bim.CommitDir, 0755)
	if err != nil {
		fmt.Printf("Error during creation of commit directory: %v\n", err)
		return
	}

	_, err = os.Create(bim.StagingArea)
	if err != nil {
		fmt.Printf("Error during creation of staging area: %v\n", err)
		return
	}

	err = os.Chmod(bim.StagingArea, 0644)
	if err != nil {
		fmt.Printf("Error during changing of staging area file permissions: %v\n", err)
		return
	}

	err = os.Chdir(workingDir)
	if err != nil {
		fmt.Printf("Failed to restore working directory: %v\n", err)
		return
	}
}
