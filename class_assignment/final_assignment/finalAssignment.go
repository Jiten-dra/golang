package final_assignment

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Post struct {
	title     string
	content   string
	mediaLink string
}

type SocialMediaAPI interface {
	CreatePost(post Post) string
	FetchCommentsAndLikes(postID string) (comments int, likes int)
}

// dummy api
type API struct {
	platform string
}

func (a API) FetchCommentsAndLikes(postID string) (comments int, likes int) {
	return rand.Intn(20), rand.Intn(50)
}

func (a API) CreatePost(post Post) string {
	return fmt.Sprintf("%s post %s created", a.platform, post.title)
}

func createDummyAPI() []SocialMediaAPI {
	socialAccounts := []string{"instagram", "twitter", "linkedin", "facebook"}
	var apiList []SocialMediaAPI
	for _, platform := range socialAccounts {
		apiList = append(apiList, API{
			platform: platform,
		})
	}
	return apiList
}

func sharePost(post Post, apis []SocialMediaAPI, results chan<- string) {
	for _, api := range apis {
		go func(api SocialMediaAPI) {
			msg := api.CreatePost(post)
			results <- msg
		}(api)
	}
}

func SharePostToAll(post Post, apis []SocialMediaAPI, wg *sync.WaitGroup) {
	wg.Add(1)
	results := make(chan string)
	sharePost(post, apis, results)
	go func() {
		defer wg.Done()
		for i := 0; i < len(apis); i++ {
			fmt.Println(<-results)
		}
		close(results)
	}()
	wg.Wait()
	fmt.Println("Post share to all account is Completed using FanOut pattern")
}

func aggregateEngagement(postID string, apis []SocialMediaAPI) (int, int) {
	var wg sync.WaitGroup
	commentsChan := make(chan int, len(apis))
	likesChan := make(chan int, len(apis))

	for _, api := range apis {
		wg.Add(1)
		go func(api SocialMediaAPI) {
			defer wg.Done()
			comments, likes := api.FetchCommentsAndLikes(postID)
			//FanIn pattern
			commentsChan <- comments
			likesChan <- likes
		}(api)
	}

	go func() {
		wg.Wait()
		close(commentsChan)
		close(likesChan)
	}()

	//FanIn pattern
	totalComments, totalLikes := 0, 0
	for comments := range commentsChan {
		totalComments += comments
	}
	for likes := range likesChan {
		totalLikes += likes
	}

	return totalComments, totalLikes
}

func Run() {
	var wg sync.WaitGroup

	apis := createDummyAPI()
	fmt.Println("all account->", apis)
	post := Post{
		title:     "title",
		content:   "content",
		mediaLink: "link",
	}
	// for _, api := range apis {
	// 	fmt.Println(api.CreatePost(post))
	// }
	SharePostToAll(post, apis, &wg)

	fmt.Println("waiting for peoples to like and comment XD\n\n\n")
	time.Sleep(5 * time.Second)
	postId := "1964"
	comments, likes := aggregateEngagement(postId, apis)
	fmt.Printf("post id %v\n   total comments: %v\n   total likes: %v\n", postId, comments, likes)

	fmt.Println("Collected all likes and comments using FanIn pattern")
}
