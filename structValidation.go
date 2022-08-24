package main

import (
	_ "connnnntrollers/docs"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//export PATH=$(go env GOPATH)/bin:$PATH

type OwnerInfor struct {
	PubKey      string `json:"pubKey"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
	Authentic   string `json:"authentic"`
}

type MapExtend struct {
	UID     string `json:"uid"`
	Pricacy string `json:"pricacy"`
	PubKey  string `json:"pubkey"`
}

type mapextend1 struct {
	Lang_Data string `json:"lang_data"`
}
type OwnerInfo struct {
	Pubkey      string `json:"pubkey"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
	Authentic   string `json:"authentic"`
}

type ListMediaItem struct {
	Name         string      `json:"name"`
	Mediatype    int64       `json:"mediatype"`
	URL          string      `json:"url"`
	IDMedia      int         `json:"idmedia"`
	IDPost       int         `json:"idpost"`
	TimesTamps   int         `json:"timestamps"`
	Extend       string      `json:"extend"`
	MapExtend    *MapExtend  `json:"mapextend"`
	TotalLikes   int         `json:"totallike"`
	ToTalShare   int         `json:"totalshare"`
	ToTalComment int         `json:"totalcomment"`
	OwnerInfo    *OwnerInfor `json:"ownerinfo"`
	Reaction     int         `json:"reaction"`
}

type Post struct {
	ID            int64          `json:"idpost" binding:"required,gte=1,lte=10000"`
	UID           int64          `json:"uid"`
	Content       string         `json:"content"`
	ListMediaItem *ListMediaItem `json:"listMediaItem"`
	IDFeeLing     string         `json:"idfeeling"`
	Privacy       int64          `json:"privacy"`
	TimesTamps    int64          `json:"timesTamps"`
	Pubkey        string         `json:"pubkey"`
	Extend        string         `json:"extend"`
	TotalReaction int64          `json:"totalReaction"`
	TotalComment  int64          `json:"totalComment"`
	TotalShare    int64          `json:"totalShare"`
	OriginPostID  int64          `json:"originPostId"`
	OwnerInfo     *OwnerInfo     `json:"ownerInfo"`
	MapExtend     *mapextend1    `json:"mapExtend"`
	Reaction      int64          `json:"reaction"`
	TotalLike     int64          `json:"totalLike"`
}

type message struct {
	Message string `json:"message"`
}

var Posts = []Post{
	{
		ID:      1147378,
		UID:     79860,
		Content: "0.01 TRUSTK\n#game\n#playtoearn",
		ListMediaItem: &ListMediaItem{
			Name:       "ReactNative-snapshot-image3236878294560589202.png",
			Mediatype:  1,
			URL:        "https://photocloud.mobilelab.vn/2022-08-08/dc166b1f-bde5-438f-b003-381009672a72.png",
			IDMedia:    1147378,
			IDPost:     1147378,
			TimesTamps: 1659932457,
			Extend:     "",
			MapExtend: &MapExtend{
				UID:     "79860",
				Pricacy: "0",
				PubKey:  "033987f519b267b16a6c5f9a4b15131213bbcda9e726af2b62e6e368c076c6c7b1",
			},
			TotalLikes:   0,
			ToTalShare:   0,
			ToTalComment: 0,
			OwnerInfo: &OwnerInfor{
				PubKey:      "033987f519b267b16a6c5f9a4b15131213bbcda9e726af2b62e6e368c076c6c7b1",
				DisplayName: "saeid",
				Avatar:      "https://photocloud.mobilelab.vn/2022-03-15/560461a8-124c-441e-bfe2-9c8bbc29a569.jpg",
				Authentic:   "",
			},
			Reaction: 0,
		},
		IDFeeLing:     "",
		Privacy:       0,
		TimesTamps:    1659932455,
		Pubkey:        "033987f519b267b16a6c5f9a4b15131213bbcda9e726af2b62e6e368c076c6c7b1",
		Extend:        "",
		TotalReaction: 0,
		TotalComment:  0,
		TotalShare:    0,
		OriginPostID:  0,
		OwnerInfo: &OwnerInfo{
			Pubkey:      "033987f519b267b16a6c5f9a4b15131213bbcda9e726af2b62e6e368c076c6c7b1",
			DisplayName: "saeid",
			Avatar:      "https://photocloud.mobilelab.vn/2022-03-15/560461a8-124c-441e-bfe2-9c8bbc29a569.jpg",
			Authentic:   "",
		},
		MapExtend: &mapextend1{
			Lang_Data: "{\"code\":\"it\",\"name\":\"Italian\"}",
		},
		Reaction:  0,
		TotalLike: 0,
	},
}

// @Summary get all items in the Post list
// @ID get-all-Post
// @Produce json
// @Success 200 {object} Post
// @Router /api/posts [get]
func getAllPost(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Posts)
}

// @Summary find a Post item by ID
// @ID find-Post-by-id
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} Post
// @Failure 404 {object} message
// @Router /api/posts/{idpost} [get]
func FindPost(c *gin.Context) {
	ID := c.Param("id")
	id, _ := strconv.ParseInt(ID, 10, 64)

	for _, find := range Posts {
		if find.ID == id {
			c.JSON(http.StatusOK, find)
			return
		}
	}

	r := message{"not found"}
	c.IndentedJSON(http.StatusNotFound, r)
}

// @Summary add a new item to the Post list
// @ID create-Post
// @Produce json
// @Param data body Post true "Post data"
// @Success 200 {object} Post
// @Failure 400 {object} message
// @Router /api/posts [POST]
func postposts(c *gin.Context) {
	var NewPosts Post

	if err := c.BindJSON(&NewPosts); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id qua dai"})
		return
	}

	Posts = append(Posts, NewPosts)
	c.IndentedJSON(http.StatusCreated, NewPosts)

}

// @Summary delete a Post item by ID
// @ID delete-Post-by-id
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} Post
// @Failure 404 {object} message
// @Router /api/posts/{idposts} [delete]
func deletePost(c *gin.Context) {
	id := c.Param("idpost")
	id2, _ := strconv.ParseInt(id, 10, 64)

	for index, b := range Posts {
		if b.ID == id2 {
			Posts = append(Posts[:index], Posts[index+1:]...)
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
}

// @Summary update a Post item by ID
// @ID update-Post-by-id
// @Produce json
// @Param id path string true "Post ID"
// @description update an Post
// @Success 200 {object} Post
// @Failure 404 {object} message
// @Router /api/posts/{idposts} [PUT]
func updatePost(c *gin.Context) {
	id := c.Param("idpost")
	var NewPost Post
	id2, _ := strconv.ParseInt(id, 10, 64)

	if err := c.BindJSON(&NewPost); err != nil {
		return
	}

	for index, b := range Posts {
		if b.ID == id2 {
			if NewPost.UID != b.UID && NewPost.UID != 0 {
				Posts[index].UID = NewPost.UID
			}
			if NewPost.Content != b.Content && NewPost.Content != "" {
				Posts[index].Content = NewPost.Content
			}

		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "User updated"})
	}
}

func handles() {
	router := gin.Default()

	router.GET("/api/posts", getAllPost)
	router.POST("/api/posts", postposts)
	router.PUT("/api/posts/:idpost", updatePost)
	router.DELETE("/api/posts/:idpost", deletePost)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:8080")
}

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath
// @query.collection.format multi
func main() {
	handles()
}
