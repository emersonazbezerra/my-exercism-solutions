package lasagna

func PreparationTime(layers []string, preparationTime int) int {
    if preparationTime == 0 { preparationTime = 2 }
    return len(layers) * preparationTime
}

func Quantities(layers []string) (int, float64) {
    var noodles int
    var sauce float64
    
	for _, v := range layers {
        if v == "noodles" {
            noodles += 50
        } else if v == "sauce" {
            sauce += 0.2
        }
    }
    return noodles, sauce
}

func AddSecretIngredient(fi, pi []string) {
    pi[len(pi) - 1] = fi[len(fi) - 1]
}

func ScaleRecipe(amountsForTwo []float64, portions int) []float64 {
    var newQuantities []float64
    for i := 0; i < len(amountsForTwo); i++ {
        quantity := (amountsForTwo[i] / 2.0) * float64(portions)
        newQuantities = append(newQuantities, quantity)
    }
    return newQuantities
}