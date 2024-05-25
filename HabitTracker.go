package main

import "fmt"

const NMAX int = 1000

type habit struct {
	nama      string
	frekuensi int
	kategori  string
	count     int
}

type storage [NMAX]habit

var habits storage
var numHabits int

func main() {
	Menu()
	for {
		var choice int
		fmt.Print("Enter your choie: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			option1()
			return
		case 2:
			option2()
			return
		case 3:
			option3()
			return
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid, Try Again.")
		}
	}

}

func Menu() {
	habitMenu()
	fmt.Println("============== Main Menu ===============")
	fmt.Println("1. My Habits")
	fmt.Println("2. Track Progress")
	fmt.Println("3. Track Habit")
	fmt.Println("4. Exit")
	fmt.Println("========================================")
}

func habitMenu() {
	fmt.Println("================ HABIT =================")
	fmt.Printf("%-30s %-15s\n", "Nama", "Progress")
	for i := 0; i < numHabits; i++ {
		fmt.Printf("%-32s %d/%d\n", habits[i].nama, habits[i].count, habits[i].frekuensi)
	}
	fmt.Println("========================================")
}

func option1() {
	fmt.Println()
	fmt.Println("============== My Habits ===============")
	fmt.Println("========================================")
	printHabit()
	myHabits()
	for {
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addhabit()
			return
		case 2:
			removeHabit()
			return
		case 3:
			editHabit()
			return
		case 4:
			main()
			return
		default:
			fmt.Println("Invalid, Try Again.")
		}
	}
}

func option2() {
	fmt.Println()
	fmt.Println("=========== Track Progress ============")
	fmt.Printf("%-30s %-15s\n", "Nama", "Progress")
	for i := 0; i < numHabits; i++ {
		fmt.Printf("%-32s %d/%d\n", habits[i].nama, habits[i].count, habits[i].frekuensi)
	}
	fmt.Println("========================================")
	freq()
	for {
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addCount(&habits)
			option2()
			return
		case 2:
			removeCount(&habits)
			option2()
			return
		case 3:
			main()
			return
		default:
			fmt.Println("Invalid, Try Again.")
		}
	}
}

func option3() {
	fmt.Println()
	fmt.Println("============= Track Habit ==============")
	Track()
	for {
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			trackbyname()
			return
		case 2:
			trackbyCategory()
			return
		case 3:
			trackbyCount()
			return
		case 4:
			main()
			return
		default:
			fmt.Println("Invalid, Try Again.")
		}
	}

}

func myHabits() {
	fmt.Println("========================================")
	fmt.Println("1. Add Habit")
	fmt.Println("2. Remove Habit")
	fmt.Println("3. Edit Habit")
	fmt.Println("4. Return")
	fmt.Println("========================================")
}

func freq() {
	fmt.Println("========================================")
	fmt.Println("1. Add Progress")
	fmt.Println("2. Remove or Edit Progress")
	fmt.Println("3. Return")
	fmt.Println("========================================")
}

func Track() {
	fmt.Println("========================================")
	fmt.Println("1. Track by Name")
	fmt.Println("2. Track by Category")
	fmt.Println("3. Track by Count")
	fmt.Println("4. Return")
	fmt.Println("========================================")
}

func addhabit() {
	var lagi int
	tambahHabit(&lagi, &habits, &numHabits)
	for {
		switch lagi {
		case 1:
			tambahHabit(&lagi, &habits, &numHabits)
		case 2:
			option1()
			return
		default:
			fmt.Println("Invalid, Try Again.")
			menutambahlagi()
			lagi = tambahLagi()
		}
	}
}

func tambahHabit(lagi *int, habits *storage, numHabits *int) {
	var kebiasaan habit
	fmt.Println("========================================")
	fmt.Print("Nama Habit: ")
	fmt.Scan(&kebiasaan.nama)
	cekfrekuensi(&kebiasaan)
	category(&kebiasaan)
	kebiasaan.count = 0
	fmt.Println("========================================")
	(*habits)[*numHabits] = kebiasaan
	*numHabits++
	menutambahlagi()
	*lagi = tambahLagi()
}

func tambahLagi() int {
	var lagi int
	fmt.Scan(&lagi)
	return lagi
}

func cekfrekuensi(kebiasaan *habit) {
	for n := 0; n != 1; {
		fmt.Print("Frekuensi Harian: ")
		n, _ = fmt.Scan(&kebiasaan.frekuensi)
		if n != 1 {
			fmt.Println("Invalid, enter an integer please.")
			var discard string
			fmt.Scanln(&discard)
		}
	}
}

func menutambahlagi() {
	fmt.Println("Tambah lagi: ")
	fmt.Println("1. Ya    2. Return")
	fmt.Print("Enter your choice: ")
}

func category(kebiasaan *habit) {
	fmt.Println("Pilih Kategori:")
	fmt.Println("1. Kesehatan")
	fmt.Println("2. Produktivitas")
	fmt.Println("3. Personal Development")
	fmt.Println("4. Other")

	for {
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			kebiasaan.kategori = "Kesehatan"
			return
		case 2:
			kebiasaan.kategori = "Produktivitas"
			return
		case 3:
			kebiasaan.kategori = "Personal Development"
			return
		case 4:
			fmt.Print("Masukan Kategori Baru: ")
			var newCategory string
			fmt.Scan(&newCategory)
			kebiasaan.kategori = newCategory
			return
		default:
			fmt.Println("Invalid, Try Again.")
		}
	}
}

func removeHabit() {
	hapusHabit(&habits, &numHabits)
	for {
		var lagi int
		fmt.Println("Remove lagi: ")
		fmt.Println("1. Ya    2. Return")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&lagi)
		switch lagi {
		case 1:
			hapusHabit(&habits, &numHabits)
		case 2:
			option1()
			return
		default:
			fmt.Println("Invalid, Try Again.")
		}
	}
}

func hapusHabit(habits *storage, numHabits *int) {
	var habitName string
	var i int
	fmt.Print("Habit yang Ingin di Remove: ")
	fmt.Scan(&habitName)

	found := -1
	for i < *numHabits && found == -1 {
		if habits[i].nama == habitName {
			found = i
			for j := i; j < *numHabits-1; j++ {
				habits[j] = habits[j+1]
			}
			*numHabits--
		}
		i++
	}
	if found != -1 {
		fmt.Println("Habit Berhasil di Remove.")
	}
	if found == -1 {
		fmt.Println("Habit tidak ditemukan.")
	}
	fmt.Println("----------------------------------------")
}

func printHabit() {
	fmt.Println()
	fmt.Println("============= Daftar Habit =============")
	fmt.Printf("%-10s %-15s %-15s\n", "Nama", "Frekuensi", "Kategori")
	for i := 0; i < numHabits; i++ {
		fmt.Printf("%-10s %-15d %-15s\n", habits[i].nama, habits[i].frekuensi, habits[i].kategori)
	}
	fmt.Println("========================================")
}

func sortHabitsbyname(habits storage, numHabits int) storage {
	for i := 1; i < numHabits; i++ {
		key := habits[i]
		j := i - 1
		for j >= 0 && habits[j].nama > key.nama {
			habits[j+1] = habits[j]
			j = j - 1
		}
		habits[j+1] = key
	}
	return habits
}

func binarySearch(habits *storage, low, high int, habitName string) int {
	if high >= low {
		mid := low + (high-low)/2
		if habits[mid].nama == habitName {
			return mid
		}
		if habits[mid].nama > habitName {
			return binarySearch(habits, low, mid-1, habitName)
		}
		return binarySearch(habits, mid+1, high, habitName)
	}
	return -1
}

func suntingHabit() {
	habits = sortHabitsbyname(habits, numHabits)
	var habitName string
	fmt.Print("Habit yang Ingin di Edit: ")
	fmt.Scan(&habitName)

	index := binarySearch(&habits, 0, numHabits-1, habitName)
	if index != -1 {
		var kebiasaan habit
		fmt.Println("----------------------------------------")
		fmt.Print("Nama Baru: ")
		fmt.Scan(&kebiasaan.nama)
		cekfrekuensi(&kebiasaan)
		category(&kebiasaan)
		habits[index] = kebiasaan
		fmt.Println("----------------------------------------")
		fmt.Println("Habit Berhasil di Edit.")
		fmt.Println("----------------------------------------")
	} else {
		fmt.Println("Habit tidak ditemukan.")
		fmt.Println("----------------------------------------")
	}

}

func editHabit() {
	suntingHabit()
	for {
		var lagi int
		fmt.Println("Edit lagi: ")
		fmt.Println("1. Ya    2. Return")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&lagi)
		switch lagi {
		case 1:
			suntingHabit()
		case 2:
			option1()
			return
		default:
			fmt.Println("Invalid, Try Again.")
		}
	}
}

func addCount(habits *storage) {
	var habitName string
	var countToAdd int
	fmt.Print("Nama Habit to Add Progress: ")
	fmt.Scan(&habitName)
	cekcount(&countToAdd)
	index := findHabitIndex(habitName)
	if index != -1 {
		habits[index].count += countToAdd
		fmt.Println("Progress added.")
		if habits[index].count >= habits[index].frekuensi {
			fmt.Println("WOW, U ARE DOING A GOOD JOB")
		}
	} else {
		fmt.Println("Habit not found.")
	}
}

func removeCount(habits *storage) {
	var habitName string
	var countToRemove int
	fmt.Print("Nama Habit to Remove Progress: ")
	fmt.Scan(&habitName)
	cekcount(&countToRemove)
	index := findHabitIndex(habitName)
	if index != -1 {
		if habits[index].count > countToRemove {
			habits[index].count -= countToRemove
			fmt.Println("Progress edited.")
		} else if habits[index].count == countToRemove {
			habits[index].count -= countToRemove
			fmt.Println("Progress removed.")
		} else {
			fmt.Println("Jumlah melebihi progress sekarang.")
		}
	} else {
		fmt.Println("Habit not found.")
	}
}

func findHabitIndex(habitName string) int {
	for i := 0; i < numHabits; i++ {
		if habits[i].nama == habitName {
			return i
		}
	}
	return -1
}

func cekcount(count *int) {
	for n := 0; n != 1; {
		fmt.Print("Seberapa Banyak Progress: ")
		n, _ = fmt.Scan(count)
		if n != 1 {
			fmt.Println("Invalid, enter an integer please.")
			var discard string
			fmt.Scanln(&discard)
		}
	}
}

func trackbyname() {
	fmt.Println()
	fmt.Println("============= Track by Name ============")
	asscenordecend()
	for {
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			nameascend(habits, numHabits)
			fmt.Println("Press Enter to return...")
			fmt.Scanln()
			fmt.Scanln()
			trackbyname()
			return
		case 2:
			namedescend(habits, numHabits)
			fmt.Println("Press Enter to return...")
			fmt.Scanln()
			fmt.Scanln()
			trackbyname()
			return
		case 3:
			option3()
			return
		default:
			fmt.Println("Invalid, Try Again.")
		}
	}
}

func asscenordecend() {
	fmt.Println("========================================")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Println("3. Return")
	fmt.Println("========================================")
}

func nameascend(habits storage, numHabits int) {
	var sorted storage
	for i := 0; i < numHabits; i++ {
		sorted[i] = habits[i]
	}
	for i := 0; i < numHabits-1; i++ {
		minIndex := i
		for j := i + 1; j < numHabits; j++ {
			if sorted[j].nama < sorted[minIndex].nama {
				minIndex = j
			}
		}
		if minIndex != i {
			sorted[i], sorted[minIndex] = sorted[minIndex], sorted[i]
		}
	}
	printfortrack(sorted, numHabits)
}

func namedescend(habits storage, numHabits int) {
	var sorted storage
	for i := 0; i < numHabits; i++ {
		sorted[i] = habits[i]
	}
	for i := 1; i < numHabits; i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j].nama < key.nama {
			sorted[j+1] = sorted[j]
			j = j - 1
		}
		sorted[j+1] = key
	}
	printfortrack(sorted, numHabits)
}

func trackbyCategory() {
	fmt.Println()
	fmt.Println("=========== Track by Category ==========")
	asscenordecend()
	for {
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			categoryascend(habits, numHabits)
			fmt.Println("Press Enter to return...")
			fmt.Scanln()
			fmt.Scanln()
			trackbyCategory()
			return
		case 2:
			categorydescend(habits, numHabits)
			fmt.Println("Press Enter to return...")
			fmt.Scanln()
			fmt.Scanln()
			trackbyCategory()
			return
		case 3:
			option3()
			return
		default:
			fmt.Println("Invalid, Try Again.")
		}
	}
}

func categoryascend(habits storage, numHabits int) {
	var sorted storage
	for i := 0; i < numHabits; i++ {
		sorted[i] = habits[i]
	}
	for i := 1; i < numHabits; i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j].kategori > key.kategori {
			sorted[j+1] = sorted[j]
			j = j - 1
		}
		sorted[j+1] = key
	}
	printfortrack(sorted, numHabits)
}

func categorydescend(habits storage, numHabits int) {
	var sorted storage
	for i := 0; i < numHabits; i++ {
		sorted[i] = habits[i]
	}
	for i := 1; i < numHabits; i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j].kategori < key.kategori {
			sorted[j+1] = sorted[j]
			j = j - 1
		}
		sorted[j+1] = key
	}
	printfortrack(sorted, numHabits)
}

func trackbyCount() {
	fmt.Println()
	fmt.Println("============ Track by Count ============")
	asscenordecend()
	for {
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			countascend(habits, numHabits)
			fmt.Println("Press Enter to return...")
			fmt.Scanln()
			fmt.Scanln()
			trackbyCount()
			return
		case 2:
			countdescend(habits, numHabits)
			fmt.Println("Press Enter to return...")
			fmt.Scanln()
			fmt.Scanln()
			trackbyCount()
			return
		case 3:
			option3()
			return
		default:
			fmt.Println("Invalid, Try Again.")
		}
	}
}

func countascend(habits storage, numHabits int) {
	var sorted storage
	for i := 0; i < numHabits; i++ {
		sorted[i] = habits[i]
	}
	for i := 0; i < numHabits-1; i++ {
		minIndex := i
		for j := i + 1; j < numHabits; j++ {
			if sorted[j].count < sorted[minIndex].count {
				minIndex = j
			}
		}
		if minIndex != i {
			sorted[i], sorted[minIndex] = sorted[minIndex], sorted[i]
		}
	}
	printfortrack(sorted, numHabits)
}

func countdescend(habits storage, numHabits int) {
	var sorted storage
	for i := 0; i < numHabits; i++ {
		sorted[i] = habits[i]
	}
	for i := 1; i < numHabits; i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j].count < key.count {
			sorted[j+1] = sorted[j]
			j = j - 1
		}
		sorted[j+1] = key
	}
	printfortrack(sorted, numHabits)
}

func printfortrack(sorted storage, numHabits int) {
	fmt.Println()
	fmt.Println("============= Daftar Habit =============")
	fmt.Printf("%-10s %-15s %-15s\n", "Nama", "Progress", "Kategori")
	for i := 0; i < numHabits; i++ {
		fmt.Printf("%-10s %d/%-15d %-15s\n", sorted[i].nama, sorted[i].count, sorted[i].frekuensi, sorted[i].kategori)
	}
	fmt.Println("========================================")
}
