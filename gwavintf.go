package gwav

type GWav interface {
	Read() (*[]byte, error)
}
