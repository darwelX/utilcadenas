// util cadenas este paquete contiene funciones para trabajar
// con cadenas de caracteres
package utilcadenas

// Reverso invierte su argumento s dejandolo legible de izquierda a derecha
func Reverso(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i<len(r)/2; i,j = i+1, j-1{
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
