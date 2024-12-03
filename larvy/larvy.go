package larvy

const (
	version = "1.0.0"
)

type Larvy struct {
	AppName string
	Debug   bool
	Version string
}

// New initialises the project
func (l *Larvy) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "temp", "logs", "middlewares"},
	}

	if err := l.Init(pathConfig); err != nil {
		return err
	}
	return nil
}

// Init creates the directories in the provided rootpath
func (l *Larvy) Init(path initPaths) error {
	root := path.rootPath
	for _, path := range path.folderNames {
		// create folders if it doesn't exist
		err := l.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
