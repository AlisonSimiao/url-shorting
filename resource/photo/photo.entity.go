package photo

type Photo struct {
	Id        int    `json:"id"`
	Type      int    `json:"type"` // 1: profile, 2: project cover, 3: default
	Url       string `json:"url"` // url da imagem
	Public_id string `json:"public_id"` // id unica no storage
	Path      string `json:"path"` // caminho no storage
}
