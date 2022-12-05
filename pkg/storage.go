package pkg

import (
	user "GolandProjects/30.1"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

var CreateIndex int

type Friend struct {
	SourceId int `json:"source_id"`
	TargetID int `json:"target_id"`
}

type Service struct {
	Store map[int]*user.User
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var u user.User

	if err = json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	CreateIndex++
	u.Id = CreateIndex
	s.Store[u.Id] = &u
	id := strconv.Itoa(u.Id)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User war created id: " + id + " name: " + u.Name))
	w.WriteHeader(http.StatusBadRequest)
}

func (s *Service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var u user.User

	if err = json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	s.Store[u.Friends[0]].Friends = append(s.Store[u.Friends[0]].Friends, u.Friends[1])
	s.Store[u.Friends[1]].Friends = append(s.Store[u.Friends[1]].Friends, u.Friends[0])
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(s.Store[u.Friends[0]].Name + " and " + s.Store[u.Friends[1]].Name + " friend"))
	w.WriteHeader(http.StatusBadRequest)
}

func (s *Service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var u user.User

	if err = json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	for _, el := range s.Store {
		for i, value := range el.Friends {
			if value == u.Id {
				copy(el.Friends[i:], el.Friends[i+1:])
				el.Friends[len(el.Friends)-1] = 0
				el.Friends = el.Friends[:len(el.Friends)-1]
			}
		}
	}
	w.Write([]byte("user " + s.Store[u.Id].Name + " delete"))
	delete(s.Store, u.Id)
	w.WriteHeader(http.StatusCreated)
	w.WriteHeader(http.StatusBadRequest)
}

func (s *Service) GetFriend(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var u user.User
	if err = json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	for _, elem := range s.Store[u.Id].Friends {
		w.Write([]byte(s.Store[elem].Name + ", "))
	}
	w.Write([]byte("is friend " + s.Store[u.Id].Name))
}

func (s *Service) UpdateAge(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var u user.User
	if err = json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	s.Store[u.Id].Age = u.Age
	w.Write([]byte("Возраст пользователя " + s.Store[u.Id].Name + " изменён!"))
	w.WriteHeader(http.StatusCreated)
	w.WriteHeader(http.StatusBadRequest)
}
