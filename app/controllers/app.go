package controllers

import "os"
import "github.com/revel/revel"
import "fmt"
import "io/ioutil"
import "regexp"
import "strings"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	searchPath := os.Getenv("DOCKER_REPOSITORY_PATH")
	serverUrl := os.Getenv("DOCKER_REGISTRY_URL")

	fis, err := ioutil.ReadDir(searchPath)
	if err != nil {
		fmt.Println(searchPath)
		fmt.Println(serverUrl)
		panic(err)
	}

	users := []*User{}
	for _, fi := range fis {
		user := fi.Name()

		repoList, repoErr := ioutil.ReadDir(searchPath + "/" + user)
		// Repository Not found
		if repoErr != nil {
			panic(repoErr)
		}
		repos := []*Repository{}

		for _, repo := range repoList {

			tagList, tagErr := ioutil.ReadDir(searchPath + "/" + user + "/" + repo.Name())
			tags := []string{}
			//TODO: errを使う
			r, experr := regexp.Compile(`^tag_`)
			// Tag Not found
			if tagErr != nil {
				panic(tagErr)
			}
			if experr != nil {
				panic(experr)
			}

			for _, tag := range tagList {
				if r.MatchString(tag.Name()) {
					tags = append(tags, strings.TrimLeft(tag.Name(), "tag_"))
				}
			}

			fmt.Println(tags)

			fmt.Println(user + "/" + repo.Name())
			repos = append(repos, NewRepository(repo.Name(), tags))
		}

		users = append(users, NewUser(user, repos))
	}

	return c.Render(users, serverUrl)
}

// ----------------------------

type User struct {
	Name  string
	Repos []*Repository
}

func NewUser(name string, repos []*Repository) *User {
	user := &User{Name: name, Repos: repos}
	return user
}

// --------------------------

type Repository struct {
	Name string
	Tags []string
}

func NewRepository(name string, tags []string) *Repository {
	repo := &Repository{Name: name, Tags: tags}
	return repo
}
