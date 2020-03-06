package car_pooling

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func Init() {
	http.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		_, err := fmt.Fprint(writer, "OK")
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/cars", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPut {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := json.NewDecoder(request.Body).Decode(&cars)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		sort.SliceStable(cars, func(i, j int) bool {
			return cars[i].Seats < cars[j].Seats
		})
	})

	http.HandleFunc("/journey", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var group Group
		err := json.NewDecoder(request.Body).Decode(&group)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		car := findAvailableCar(&group)
		if car != nil {
			addGroupToCar(&group, car)
		}
		groups = append(groups, &group)
	})

	http.HandleFunc("/dropoff", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := request.ParseForm()
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(request.Form.Get("ID"))
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		var group *Group
		var index *int
		for i, g := range groups {

			if g.Id == id {
				group = g
				index = &i
				break
			}
		}

		if group == nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		dropoff(index, group)
		checkForNewSpaces()
	})

	http.HandleFunc("/locate", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := request.ParseForm()
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(request.Form.Get("ID"))
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		var group *Group
		for _, g := range groups {

			if g.Id == id {
				group = g
				break
			}
		}

		if group == nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		if group.Car == nil {
			writer.WriteHeader(http.StatusNoContent)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(group.Car)
		if err != nil {
			log.Println(err)
		}
	})
}

func checkForNewSpaces() {
	for _, group := range groups {
		if group.Car == nil {
			car := findAvailableCar(group)
			if car != nil {
				addGroupToCar(group, car)
			}
		}
	}
}

func dropoff(index *int, group *Group) {
	groups = append(groups[:*index], groups[*index+1:]...)
	freeSeats := *group.Car.FreeSeats + group.People
	group.Car.FreeSeats = &freeSeats

	carGroups := group.Car.Groups
	for i, g := range carGroups {
		if g.Id == group.Id {
			group.Car.Groups = append(carGroups[:i], carGroups[i+1:]...)
			break
		}
	}
}

func addGroupToCar(group *Group, car *Car) {
	car.Groups = append(car.Groups, group)
	group.Car = car

	if car.FreeSeats == nil {
		freeSeats := car.Seats - group.People
		car.FreeSeats = &freeSeats
	} else {
		freeSeats := *car.FreeSeats - group.People
		car.FreeSeats = &freeSeats
	}
}

func findAvailableCar(group *Group) *Car {
	for _, car := range cars {
		people := group.People

		hasEnoughSeat := car.Seats >= people
		hasEnoughFreeSeat := car.FreeSeats == nil || *car.FreeSeats >= people

		if hasEnoughSeat && hasEnoughFreeSeat {
			return car
		}
	}
	return nil
}
