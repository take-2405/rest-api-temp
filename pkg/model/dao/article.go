package dao

import (
	"backend-record/pkg/model/dto"
	"database/sql"
	"log"
)

const (
	getArticlesWithoutFilters = "SELECT article_id,title,image_path FROM articles_contents WHERE genre=? LIMIT 10;"
	getYearDesignatedArticles = "SELECT article_id,title,image_path FROM articles_contents WHERE genre=? AND era_year=? LIMIT 10;"
	getMonthlyArticles        = "SELECT article_id,title,image_path FROM articles_contents WHERE genre=? AND era_month=? LIMIT 10;"
	getArticlesWithFilters    = "SELECT article_id,title,image_path FROM articles_contents WHERE genre=? AND era_year=? AND era_month=? LIMIT 10;"
	searchTargetArticleTag	  = "SELECT article_id FROM articles_tag WHERE article_tag=? LIMIT 10;"
	searchArticleWithTag      = "SELECT article_id,title,image_path FROM articles_contents WHERE article_id=?;"
	getArticleTags            = "SELECT article_tag FROM articles_tag WHERE article_id=?;"
	getArticleDetail          = "SELECT title,image_path,context,nice FROM articles_contents WHERE article_id=?;"
	checkNiceStatus = "SELECT nice_id FROM articles_nice_status WHERE user_id=? AND article_id=?;"
	checkListStatus           = "SELECT list_id FROM users_list WHERE user_id=? AND article_id=?;"
	getCommentOnArticle       = "SELECT user_name,user_image,comments_contents FROM articles_comments WHERE article_id=? LIMIT 20;"
)

// /read/articles
type readArticles struct{
	Genre string `json:"genre"`
	Month int `json:"month"`
	Year int `json:"year"`
}

func MakeReadArticlesClient (request dto.ArticleListRequest)readArticles{
	return readArticles{Genre: request.Genre,Month: request.Month,Year: request.Year}
}

func (info *readArticles)Request() ([]dto.Article, error) {
	var (
		articleList []dto.Article
		rows        *sql.Rows
		err         error
		article     *dto.Article
	)
	//フィルタリングの有無によって実行するSQL文を変更
	if info.Month == 0 && info.Year == 0 {
		rows, err = Conn.Query(getArticlesWithoutFilters, info.Genre)
	} else if info.Month == 0 {
		rows, err = Conn.Query(getYearDesignatedArticles, info.Genre, info.Year)
	} else if info.Year == 0 {
		rows, err = Conn.Query(getMonthlyArticles, info.Genre, info.Month)
	} else {
		rows, err = Conn.Query(getArticlesWithFilters, info.Genre, info.Year, info.Month)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//取得してきた複数(単数)のレコード1つずつ処理
	for rows.Next() {
		//レコードを構造体Articleに整形
		article, err = dto.ConvertToArticle(rows)
		//記事に対応する複数のタグを取得する
		article.Tags, err = GetTags(article.ArticleID)
		if err != nil {
			return nil, err
		}
		articleList = append(articleList, *article)
	}
	return articleList, nil
}

func GetTags(articleID string) ([]string, error) {
	var tags []string
	rows, err := Conn.Query(getArticleTags, articleID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			log.Println(err)
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// /read/tag/articles
type readTagArticles struct{
	Tag string `json:"tag"`
}
func MakeReadTagArticlesClient(request dto.TagRequest)readTagArticles{
	return readTagArticles{Tag: request.Tag}
}

func (info *readTagArticles)Request()([]dto.Article, error){
	var (
		articles []dto.Article
		row        *sql.Row
		err         error
	)
	log.Println(info.Tag)
	articleIDs, err := searchTagTargetArticleID(info.Tag);
	if err!=nil{
		return nil,err
	} else if (len(articleIDs)==0){
		log.Println("There are no articles assigned the specified tag.")
		return nil,err
	}
	for i := 0; i < len(articleIDs); i++ {
		var article    dto.Article
		row = Conn.QueryRow(searchArticleWithTag, articleIDs[i])
		if err = row.Scan(&article.ArticleID,&article.Title,&article.ImagePath); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			log.Println(err)
			return nil, err
		}
		//記事に対応する複数のタグを取得する
		article.Tags, err = GetTags(article.ArticleID)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func searchTagTargetArticleID(tag string)([]string,error){
	var articleIDs []string
	rows, err := Conn.Query(searchTargetArticleTag, tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//取得してきた複数(単数)のレコード1つずつ処理
	for rows.Next() {
		var articleID *string
		//レコードを構造体Articleに整形
		articleID, err = dto.ConvertToArticleID(rows)
		if err != nil {
			return nil, err
		}
		articleIDs = append(articleIDs, *articleID)
	}
	return articleIDs, nil
}

// /read/article
type readArticle struct{
	ArticleID string `json:"articleID"`
}

func MakeReadArticleClient(request dto.ArticleDetailRequest)readArticle{
	return readArticle{ArticleID: request.ArticleID}
}

func (info *readArticle)Request() (*dto.ArticleDetail, error) {
	row := Conn.QueryRow(getArticleDetail, info.ArticleID)
	articleDeatil,err :=dto.ConvertToArticleDetails(row)
	if err != nil{
		return nil, err
	}
	articleDeatil.Tags,err=GetTags(info.ArticleID)
	return articleDeatil,err
}

func (info *readArticle)CheckNiceStatus(userID,ArticleID string) (bool, error) {
	var niceID string
	status := false
	row := Conn.QueryRow(checkNiceStatus, userID, ArticleID)
	if err := row.Scan(&niceID); err != nil {
		if err == sql.ErrNoRows {
			return status, nil
		}
		log.Println(err)
		return status, err
	}
	if niceID != "" {
		status = true
	}
	return status, nil
}

func (info *readArticle)CheckListStatus(userID string) (bool, error) {
	var listID string
	status := false
	row := Conn.QueryRow(checkListStatus, userID, info.ArticleID)
	if err := row.Scan(&listID); err != nil {
		if err == sql.ErrNoRows {
			return status, nil
		}
		log.Println(err)
		return false, err
	}
	if listID != "" {
		status = true
	}
	return status, nil
}

func (info *readArticle)GetComments() ([]dto.Comment, error) {
	var comments []dto.Comment
	var comment dto.Comment
	rows, err := Conn.Query(getCommentOnArticle, info.ArticleID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&comment.UserName,&comment.UserImage,&comment.Contents); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			log.Println(err)
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
