package currency

type Currency struct {
	ID        int32
	Name      string
	ShortName string
	Label     string
}

func GetCurrency(cur *Currency) string {
	cur.ID = 1
	cur.Name = `testName`
	cur.ShortName = `testShortName`
	cur.Label = `testLabel`

	return "Currency: " + cur.Name + " " + cur.ShortName + " " + cur.Label
}
