package log_demo

type Formatter interface {
	Format(entry *Entry) error
}
