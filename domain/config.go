package domain

type Config struct {
	Entries []struct {
		Src struct {
			Path string
		}
		Dests []struct {
			Repo string
			Path string
		}
	}
}
