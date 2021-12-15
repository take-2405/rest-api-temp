package dto

import (
	"database/sql"
	"log"
)

//Request struct
//記事一覧、検索のリクエスト
type ArticleListRequest struct {
	Genre string `json:"genre"`
	Month int `json:"month"`
	Year int `json:"year"`
}

//タグ検索リクエスト
type TagRequest struct {
	Tag string `json:"tag"`
}

//記事の詳細のリクエスト
type ArticleDetailRequest struct {
	ArticleID string `json:"articleID"`
}

//-------------------------------------//
//SQL struct
//-------------------------------------//

//  記事の一覧表示データ
type Article struct{
	ArticleID string `json:"id"`
	Title string `json:"title"`
	ImagePath string `json:"imagePath"`
	Tags []string `json:"tags"`
}

type ArticleDetail struct {
Title string `json:"title"`
ImagePath string `json:"imagePath" `
Context string `json:"context" `
Nice int `json:"nice"`
ArticleUserInfo ArticleUserInfo `json:"userStatus"'`
Tags []string `json:"tags"`
Comments []Comment `json:"comments"`
}

type ArticleUserInfo struct {
Nice bool
List bool
}

// コメント
type Comment struct {
	UserName string `json:"userName"`
	UserImage string `json:"userImage"`
	Contents string `json:"contents"`
}

//取得したrowsデータを記事データに変換します
func ConvertToArticle(rows *sql.Rows) (*Article, error){
	var article Article
	if err := rows.Scan(&article.ArticleID,&article.Title,&article.ImagePath); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &article, nil
}

func ConvertToArticleDetails(row *sql.Row) (*ArticleDetail, error) {
	var article ArticleDetail
	if err := row.Scan(&article.Title,&article.ImagePath,&article.Context,&article.Nice); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &article, nil
}

func ConvertToArticleID(rows *sql.Rows)(*string,error){
	var articleID *string
	if err := rows.Scan(&articleID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return articleID, nil
}