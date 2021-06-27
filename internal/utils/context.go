package utils

type ContextValues struct {
    M map[string]string
}

func (v ContextValues) Get(key string) string {
    return v.M[key]
}