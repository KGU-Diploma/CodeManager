package container


type ContainerRunner interface {
	RunContainer(projectDir, linterImage string) (string, error)
}
